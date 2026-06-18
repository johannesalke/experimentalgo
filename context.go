package main

import (
	"context"
	"fmt"

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
