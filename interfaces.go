package main

import (
	"fmt"
	"io"
	"os"
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
	contents   []byte
	readOffset int
}

func (c *Content) Read(buf []byte) (n int, err error) {
	n = copy(buf, c.contents[n:])
	c.readOffset += n
	return n, nil
}

func experiment_reader() {
	cont := Content{[]byte("test-this-thing"), 0}

	fmt.Printf("Struct: %s\n", cont.contents)
	var reader io.Reader
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
	q := Queue{[]byte{}}
	var queue io.ReadWriter
	queue = &q
	queue.Write([]byte("test1test2test3"))
	buf := make([]byte, 4)
	n, _ := queue.Read(buf)
	fmt.Printf("Quantity: %d\nConent: %s\n", n, string(buf))

	queue.Write([]byte(" other-test"))
	for n, _ = queue.Read(buf); n > 0; n, _ = queue.Read(buf) {
		fmt.Printf("Quantity: %d\nConent: %s\n", n, string(buf[0:n]))
	}

}

////////| Building a ReadWriter with writer offsets |////////////////////////////////

type Offsets struct {
	contents    []byte
	readOffset  int
	writeOffset int
}

func (c *Offsets) Read(buf []byte) (n int, err error) {
	n = copy(buf, c.contents[c.readOffset:])
	c.readOffset += n
	return n, nil
}

func (c *Offsets) Write(buf []byte) (n int, err error) {
	required := c.writeOffset + len(buf)

	if required > len(c.contents) {
		newBuf := make([]byte, required)
		copy(newBuf, c.contents)
		c.contents = newBuf

	}

	n = copy(c.contents[c.writeOffset:], buf)
	c.writeOffset += n
	return n, nil
}

func experiment_offsets() {
	offset := Offsets{[]byte{}, 0, 0}
	offset.Write([]byte("Offset to where?"))
	fmt.Printf("Struct: %s\n", offset.contents)
	//fmt.Print(offset)
	buf := make([]byte, 32)
	n, _ := offset.Read(buf)
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

//////////////| os.Stdin & Stdout are reader/writer interfaces in GO??? |/////////////////////

func experiment_stdio() {
	var stdin io.Reader

	stdin = os.Stdin
	fmt.Print(stdin)
}

///////////////| Does Go have function types? |/////////////////

type exp_func func(string, string) string

func experiment_functiontype() { //Turns out, it does. I came across the concept while learning Javascript and initially thought of Interfaces, but then considered that Go seems like it would have the just basic functiontypes as well.

	var fun exp_func

	fun = func(s1 string, s2 string) string {
		return s1 + s2
	}
	fmt.Println(fun("this", "that"))

	fun = test_functiontype

	fmt.Println(fun("this", "that"))
}

func test_functiontype(s1 string, s2 string) string {
	return s1 + "&" + s2
}
