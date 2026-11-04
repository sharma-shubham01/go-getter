package main

import (
	"fmt"
	"net/http"
)

// Serve Mux
// 1. ServeMux is the router provided by Go standard lib net/http package
// 2. ServeMux is a route multiplexer allows you to match different request handlers to different URLs served by your server
// 3. Registers patterns with their respective handlers

// Patterns
// 1. They are the URL patterns that serve mux will accept
// 2. The general pattern is of form [Method] [Host] / [Path]

// Path Conflict
// 1. If two or more patterns match a request, then the most specific pattern takes precedence
// 2. As an example of general rule, '/hello/there/' is more specific than '/hello', so both can be registered
// 3. The former matches paths beginning with '/hello/there/' and the latter will match any other path in the '/hello' subtree

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello"))
	})

	router.HandleFunc("POST /hello/user_id", func(writer http.ResponseWriter, request *http.Request) {
		userId := request.PathValue("user_id")

		writer.Write([]byte(fmt.Sprintf("Hello %s", userId)))
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
	}

}
