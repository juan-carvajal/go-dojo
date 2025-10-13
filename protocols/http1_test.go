package protocols

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func getRequestProtocol(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request Protocol: %s\n", r.Proto)
}

// Example_http1_1 demonstrates handling an HTTP/1.1 request and printing the protocol version.
func Example_http1_1() {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		getRequestProtocol(w, req)
	}))
	defer testServer.Close()

	resp, err := http.Get(testServer.URL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}
	fmt.Print(string(b))
	// Output: Request Protocol: HTTP/1.1
}
