# Communication Protocols - HTTP
Probably one of the most used procotols for web communication. Based on TCP (most of the time, more on this later).

## HTTP 1.1
HTTP can use a single connection for multiple requests, but it has a big drawback, requests need to Go out in order and come back in order, meaning that any response that is delayed, will delay the entire queue. This is called [Head-of-line (HoL) blocking](https://en.wikipedia.org/wiki/Head-of-line_blocking). To solve this, most clients maintain multiple connections, but this means extra overhead both for the client and the server.

## HTTP 2.0
HTTP 2 came with a lot of improvements, one in particular that makes sense to `HoL` blocking is the implementation of HTTP streams. Each stream can work in parallel at the application layer, solving part of the `HoL` blocking problem.
However, since HTTP 2 is still based on TCP, at the transport layer, the packets still need to be delivered in order, so, even if all the other streams’ data is sitting in the buffer ready to go, the server still has to wait for the delayed stream’s data to arrive before it can process the rest. 

## HTTP 3 (QUIC)
HTTP 3 is powered by the QUIC protocol, which is itself powered by UDP. UDP does not have to deliver packets in order, so the `HoL` problem is solved with this protocol.

# HTTP 2.0 in Go
HTTP 2.0 is actually enabled by default **IF AND ONLY IF** using TLS.

Example:
```go
// server
func getRequestProtocol(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request Protocol: %s\n", r.Proto)
}

func main() {
	http.HandleFunc("/", getRequestProtocol) // Root endpoint
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

// client
func main() {
	resp, _ := (&http.Client{}).Get("http://localhost:8080")
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("Response:", string(body))
}

// Response: Request Protocol: HTTP/1.1
```

This means that default HTTP server and client are HTTP 2.0 **READY**, but only over HTTPS by default. It is possible to run HTTP 2.0 connections over regular HTTPS tho:
```go
var protocols http.Protocols
protocols.SetUnencryptedHTTP2(true)

// server
server := &http.Server{
    Addr:      ":8080",
    Handler:   http.HandlerFunc(rootHandler),
    Protocols: &protocols,
}

// client
client := &http.Client{
    Transport: &http.Transport{
        ForceAttemptHTTP2: true,
        Protocols:         &protocols,
    },
}

// Response: Request Protocol: HTTP/2.0
```

