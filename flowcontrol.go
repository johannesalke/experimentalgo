package main

import "fmt"

// I was surprised to see what's essentially an assembly instruction in Go.
func experiment_goto_label() {

	var n int = 0
thispoint:

	n++

	if n < 5 {
		fmt.Println(n)
		goto thispoint
	}
	fmt.Println("Escaped the loop!")

}

// Turns out you can escape more than one loop at once. Definitely I could've used a few times before.
func experiment_break_outer_loop() {
outerloop:
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			if j == 2 && i == 1 {
				break outerloop
			}
			fmt.Println(i, j)
		}

	}

	fmt.Println("Escaped the loop!")

}

// Deferred functions will still execute, even if 'panic' was triggered. This makes them great for necessary cleanup, such as lifting a lock.
func experiment_panic_and_defer() {
	defer fmt.Println("This print statement is deferred until further notice.")
	panicking_subfunction()
}

// By calling recover() inside a deferred function, panic can be 'caught', a bit like a try/catch block.
// It must be inside a deferred function since everything else after the panic statement is skipped.
func experiment_panic_and_recovery() {
	fmt.Printf("calling recover during execution: %v\n", recover())
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered after panicking. Received this message:", r)
		}
	}()
	panicking_subfunction()

}

func panicking_subfunction() {

	panic("Help! I'm panicking!")

}
