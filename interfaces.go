package main

import (
	"fmt"
	"io"
)

/////| Testing the behavior of interface-type variables |/////

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
	contents    []byte
	readOffset  int
	writeOffset int
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
	cont := Content{[]byte("test-this-thing"), 0, 0}

	fmt.Printf("Struct: %s\n", cont.contents)
	var reader io.ReadWriter
	reader = &cont
	buf := make([]byte, 16)
	n, _ := reader.Read(buf)
	fmt.Print(n, "\n")
	fmt.Printf("Read: %s\n", buf)

}

////////|Building a ReadWriter that consumes data as it is read |///////////

type Queue struct {
	contents []byte
}

func (c *Queue) Read(buf []byte) (n int, err error) {
	n = copy(buf, c.contents)
	c.contents = c.contents[n:]
	return n, nil
}

func (c *Queue) Write(buf []byte) (n int, err error) {
	c.contents = append(c.contents, buf...)
	n = len(buf)
	return n, nil
}

func experiment_readwriter() {
	queue := Queue{[]byte{}}
	queue.Write([]byte("test1test2test3"))
	buf := make([]byte, 4)
	n, _ := queue.Read(buf)
	fmt.Printf("Quantity: %d\nConent: %s\n", n, string(buf))

	queue.Write([]byte(" other-test"))
	for n, _ = queue.Read(buf); n > 0; n, _ = queue.Read(buf) {
		fmt.Printf("Quantity: %d\nConent: %s\n", n, string(buf[0:n]))
	}
}

////////| Building a ReadWriter with offsets |////////////////////////////////

type Shared struct {
	contents    []byte
	readOffset  int
	writeOffset int
}

func (c *Shared) Read(buf []byte) (n int, err error) {
	n = copy(buf, c.contents)
	return n, nil
}

func (c *Shared) Write(buf []byte) (n int, err error) {
	n = copy(c.contents, buf)
	return n, nil
}

func experiment_shared() {
	cont := Content{[]byte{}, 0, 0}

	fmt.Printf("Struct: %s\n", cont.contents)
	var reader io.ReadWriter
	reader = &cont
	buf := make([]byte, 16)
	n, _ := reader.Read(buf)
	fmt.Print(n, "\n")
	fmt.Printf("Read: %s\n", buf)

}

/*
	fmt.Printf("Struct: %s\n", queue.contents)
	var reader io.ReadWriter
	reader = &queue

	n, _ := reader.Read(buf)
	fmt.Print(n, "\n")
	fmt.Printf("Read: %s\n", buf)
*/
