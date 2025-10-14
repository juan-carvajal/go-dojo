package protocols

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// Example_http2_without_TLS demonstrates how HTTP 2 will not be enabled with default if not using TLS.
func Example_http2_without_TLS() {
	testServer := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		getRequestProtocol(w, req)
	}))
	testServer.EnableHTTP2 = true
	testServer.Start()
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

// Example_http2_TLS demonstrates how to enable HTTP 2 using httptest server package.
func Example_http2_TLS() {
	testServer := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		getRequestProtocol(w, req)
	}))
	testServer.EnableHTTP2 = true
	defer testServer.Close()
	testServer.StartTLS()

	client := testServer.Client()
	resp, err := client.Get(testServer.URL)
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
	// Output: Request Protocol: HTTP/2.0
}
