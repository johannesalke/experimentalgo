package main

import (
	"fmt"
	"reflect"
)

//Normally don't make experimental function that require input, but in this case it seems near-required.

func experiment_generic_function[T any](args []T) T {
	if len(args) > 0 {
		t := reflect.TypeOf(args[0])
		fmt.Printf("Type of input elements: %s\n", t)
		fmt.Println(args)
		var res T = args[0]
		return res
	} else {
		fmt.Println("Nothing provided.")
		var res T
		return res
	}

}

// Limit what the generic can be to types satisfying specific interfaces.
func experiment_constrained_generic[T interface512](arg T) {

	fmt.Printf("Signal: %b", arg.Signal())

}

type interface512 interface {
	Signal() int
}
type struct256 int8

func (s struct256) Signal() int {
	return int(s)
}

// Limit what the generic can be to a subset of types
func experiment_union_constrained_generic[T number](n T) {
	fmt.Println(n * 2)

}

// These are not true unions as Javascript might have them. They only work for generic constraints.
type number interface {
	~float64 | ~float32 | ~int | ~int64
}

// The ~ makes it so that any type based on the base types also satisfies the interface
type intProxy int

func experiment_generic_interface() {
	db := doubler(2)
	db.Multiply(5)
	m := Multiplier[float64]{val: 7.123}
	m.Multiply(12.1341) //Automatically adjusts to require a float64

	var gen genericInterface[int] = db
	gen.Multiply(4)
	var gen2 genericInterface[float64] = m
	gen2.Multiply(12.123)

} //Ok, I'M still not 100% sure how exactly this works, but it's fine enough for now.

type genericInterface[P number] interface {
	Multiply(P)
}

type doubler int

func (d doubler) Multiply(factor int) {
	fmt.Println(factor * int(d))
}

type Multiplier[P number] struct {
	val P
}

func (m Multiplier[P]) Multiply(factor P) {
	fmt.Println(m.val * factor)
}
