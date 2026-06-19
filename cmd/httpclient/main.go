package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	context_provider_client()

}

func context_provider_client() { //This is the client for "experiment_context_from_http_request()" in the context.go file.

	resp, err := http.Get("http://localhost:21512")
	if err != nil {
		print("Error!\n")
		fmt.Println(err)
	} else {
		print("Success!\n")
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}
