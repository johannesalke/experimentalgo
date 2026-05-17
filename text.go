package main

import (
	//"strings"
	"fmt"
)

func experiment_format_preexisting_string() {
	template := "I've tried this before, but %s"
	fmt.Printf(template, "maybe this time it actually works...\n")
	result := fmt.Sprintf(template, "maybe this time it actually works...")

	fmt.Println(result)
}
