package main

import (
	"context"
	"fmt"
	"net/http"

	"time"
)

func experiment_context_withcancel() {
	fmt.Println("Running tests on context with cancel:")
	ctx, cancel := context.WithCancel(context.Background())
	print("Starting execution...\n")

	go timed_printer(ctx, 5)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	ctx, cancel = context.WithCancel(context.Background())
	print("Starting execution...\n")

	go timed_printer(ctx, 2)
	time.Sleep(5 * time.Second)
	cancel()

}

func timed_printer(ctx context.Context, delaySeconds int) {
	select {
	case <-time.After(time.Duration(delaySeconds) * time.Second):
		fmt.Printf("Execution completed after %d seconds!\n", delaySeconds)
	case <-ctx.Done():
		fmt.Printf("Execution interrupted!\n")

	}

}

func experiment_context_withdeadline() {
	fmt.Println("Running tests on context with timeout:")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //This should run fine due to time running out after printing.
	print("Starting execution...\n")
	go timed_printer(ctx, 2)
	time.Sleep(6 * time.Second)
	cancel()
	print("Starting execution...\n")
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second) //This should be interrupted
	go timed_printer(ctx, 5)
	time.Sleep(6 * time.Second)
	cancel()

}

func experiment_context_from_http_request() {
	http.HandleFunc("/", contextMiddleware(context_handler))
	http.ListenAndServe(":21512", nil)
}

func context_handler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(299)
	w.Write([]byte("Response"))
	ctx := r.Context()
	fmt.Println("Time received: ", ctx.Value("time_received"))

}

func contextMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "time_received", time.Now())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
