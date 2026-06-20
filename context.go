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

// What happens if the http request is canceled by the client before the response is sent by the server?
func experiment_context_request_cancelled() {
	http.HandleFunc("/", contextMiddleware(request_cancelled_handler))
	http.ListenAndServe(":21512", nil)

}

func request_cancelled_handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-time.After(10 * time.Second):
		w.Write([]byte("Response"))

		fmt.Println("Time received: ", ctx.Value("time_received"))
	case <-ctx.Done():
		fmt.Println("Request was cancelled before it could be fulfilled!")

	}
} // As expected, if the request is cancelled or the connection is interrupted, the context is also cancelled.

// ////////////////| How does context inheritance work in tree patterns? |//////////////////////////
func experiment_context_tree_branch() {
	ctx1 := context.Background()
	ctx2, cancel2 := context.WithCancel(ctx1)
	ctx3 := context.WithValue(ctx2, "this", "that")
	ctx4 := context.WithValue(ctx2, "1", "2")

	go subcontext_resolver(ctx1, "ctx1")
	go subcontext_resolver(ctx2, "ctx2")
	go subcontext_resolver(ctx3, "ctx3")
	go subcontext_resolver(ctx4, "ctx4")

	cancel2()

}

func subcontext_resolver(ctx context.Context, ctxName string) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Context resolved:", ctxName)
	case <-ctx.Done():
		fmt.Println("Context cancelled:", ctxName)
		fmt.Println(ctx.Value("this"))
	}
	time.Sleep(5 * time.Second)
}

/*Output:
Context cancelled: ctx4
<nil>
Context cancelled: ctx2
<nil>
Context cancelled: ctx3
that
*/ //=> Value storage is only transmitted to descending contexts. ctx1 never gets to finish due to the

func experiment_context_tree_leaf() {
	ctx1 := context.Background()
	ctx2, _ := context.WithCancel(ctx1)
	ctx3, cancel3 := context.WithCancel(context.WithValue(ctx2, "this", "that"))
	ctx4 := context.WithValue(ctx2, "1", "2")

	go subcontext_resolver(ctx1, "ctx1")
	go subcontext_resolver(ctx2, "ctx2")
	go subcontext_resolver(ctx3, "ctx3")
	go subcontext_resolver(ctx4, "ctx4")
	cancel3()
	//cancel2()

	time.Sleep(5 * time.Second)

}
