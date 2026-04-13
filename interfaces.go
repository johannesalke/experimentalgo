package main

import (
	"fmt"
	"io"
)

type math interface {
	Area() int
}

type square struct {
	length int
}

type rectangle struct {
	length int
	width  int
}

func (s square) Area() int {
	return s.length * s.length
}

func (r rectangle) Area() int {
	return r.length * r.width
}

func experiment_interfaces() {
	var math math

	math = square{5}
	fmt.Printf("Area: %d\n", math.Area())
	math = &rectangle{2, 5}
	fmt.Printf("Area: %d\n", math.Area())

}

////////| Making my own io.Reader type object |//////////////////

type Content struct {
	contents []byte
	offset   int
}

func (c *Content) Read(buf []byte) (n int, err error) {
	n = copy(buf, c.contents)
	return n, nil
}

func (c *Content) Write(buf []byte) (n int, err error) {
	n = copy(c.contents, buf)
	return n, nil
}

func experiment_reader() {
	cont := Content{contents: []byte("12345")}
	fmt.Printf("Struct: %s\n", cont.contents)
	var reader io.ReadWriter
	reader = &cont
	buf := make([]byte, 16)
	n, _ := reader.Read(buf)
	fmt.Print(n, "\n")
	fmt.Printf("Read: %s\n", buf)

}
