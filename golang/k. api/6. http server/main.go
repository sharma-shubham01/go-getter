package main

import (
	"fmt"
	"net/http"
)

// http server

// listenAndServe
// 1. listenAndServe function provided by net/http package to start a http server
// 2. func listenAndServe(addr string, handler Handler) error
// 3. listenAndServe listens on the provided address and uses handler to handle requests
// 4. It will use the default serve mux if the provided handler is nil

// http.Handler

// 1. type Handler interface {
// 	   ServeHTTP(ResponseWriter, *Request)
// }

// 2. A Handler responds to a HTTP request or handled the incoming requests for the server

// Handling Requests
// 1. The handler should write reply headers and data to the ResponseWriter and then return
// 2. It is not valid to write the ResponseWriter or Read from the request body after the request has been completed
// 3. Read from the request before writing response
// 4. Request handlers should not modify the incoming requests
// 5. Panics inside handlers are logged to server logs
// 6. Panic with ErrAbortHandler (Sentinel error) if you want the client to know

// http.ResponseWriter

// 1. type ResponseWriter interface {
// Header() Header
// Write([]byte) (int, error)
// WriteHeader(statusCode int)
// }

// 2. The method Header() returns a header map that is set on the Response
// 3. The Write() method writes the data to a connection as part of an HTTP reply
// 4. The WriteHeader() method sends an http response header with the provided status code

// http.Request
// 1. The http.Request struct represents an http request received by a server or to be sent by a client
// 2. Field method of type string specifies the Http method (Get, Post, Put etc...)
// 3. For client request, an empty string means Get
// 4. URL *url.URL specifies either URI being requested when used on the server side or the URL to access when used on the client side
// 5. Header of type map[string][]string contains the request header fields either received by the server or to be sent by the client

// Request.Body
// 1. For client requests, a nil body means the request has no body, such as a Get request
// 2. For server requests, a server body is always non-nil but will return EOF immediately when no body is present
// 3. The Server will close the request body. The ServeHTTP handler does not need to
// 4. Form url.Values contains parsed form data, including both the URL field's query parameters and the PATCH, POST or PUT form data
// 5. This field is only available after ParseForm is called

func main() {
	http.Handle("/hello", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello!"))
	}))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

// Run the server in one terminal --> go run main.go
// Hit the curl in other terminal --> curl -X GET localhost:8080/hello
