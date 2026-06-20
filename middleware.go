package main

import "net/http"

func testMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your middleware logic goes here...
		next.ServeHTTP(w, r)
	})
}

//The http.HandlerFunc(...) represents a conversion/casting of the anonymous function into the specific shape of an http.Handler
//This is possible due to the anonymous function fulfilling the function signature.

func experiment_http_middleware(next http.Handler) http.Handler { //This time, written from memory instead of copied from somewhere.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}
