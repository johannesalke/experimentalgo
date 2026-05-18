package main

import (
	"fmt"
)

func experiment_format_preexisting_string() {
	template := "I've tried this before, but %s"
	fmt.Printf(template, "maybe this time it actually works...\n")
	result := fmt.Sprintf(template, "maybe this time it actually works...")

	fmt.Println(result) //It does in fact work. I wonder why it didn't before. Must've made some mistake.
}

func experiment_printing_runes() {
	r := 'a'
	fmt.Printf("%s\n", r)         //This doesn't work and produces a sort-of error
	fmt.Printf("%s\n", string(r)) //Converting to string first works
	fmt.Printf("%x\n", r)         //Printing as hexadecimal works, as expected since a rune is effectively a number.
	fmt.Printf("%d\n", r)         //Yup, printing to decimal works too.
	fmt.Printf("%c\n", r)         //I assume c stands for character. Looks like this is the 'intended' way of printing a rune?
	fmt.Printf("%q\n", r)         //Quoting the literal contents also works, sort of. More of a debug thing through.

	fmt.Printf("%c\n", 121)        //Good to know this also works.
	fmt.Printf("%s\n", 0b01001011) //Since this was just a byte, I though it might work for a string, but it seems to just count as a number.
	// Actually not that surprising once I consider that having exactly 8 bits is arbitrary on my side and the compiler can't know I mean it to be a byte instead of just a binary number.
	fmt.Printf("%c\n", 0b01001010) //And in turn, this one does work.
	var b byte = 0b01001010
	fmt.Printf("%s\n", string(b))          // I guess this works, but I'm not happy about it :(
	fmt.Printf("%s\n", string(0b01001010)) //...ok

	_, err := fmt.Printf("%s\n", r)
	fmt.Printf("String format error: %s,", err) //Wrong, but not an actual error it seems.

}
