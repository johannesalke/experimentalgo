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
