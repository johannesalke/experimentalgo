package main

import (
	"bufio"
	"fmt"
	"time"

	"net"
	"os"
)

func experiment_tcp_test() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":42069") //Listen, I didn't choose this port number. The course I built this client in used it, and now I'm just stuck with it. Have mercy on this poor soul, will you?
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())

	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())

	}
	msg := []byte("This is a test message.")
	conn.Write(msg)
	conn.Close()
	fmt.Println("Connection closed")

}

func experiment_tcp_writer() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":21512") //Listen, I didn't choose this port number. The course I built this client in used it, and now I'm just stuck with it. Have mercy on this poor soul, will you?
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())

	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())

	}
	fmt.Print("Connection opened!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		msg := []byte(scanner.Text() + "\n")
		if string(msg) == "exit\n" {
			break
		}
		conn.Write(msg)

	}

	conn.Close()
	fmt.Println("Connection closed")

}

func experiment_tcp_readwriter() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":21512") //Listen, I didn't choose this port number. The course I built this client in used it, and now I'm just stuck with it. Have mercy on this poor soul, will you?
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())

	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())

	}
	fmt.Println("Connection opened!")
	inputScanner := bufio.NewScanner(os.Stdin)
	responseScanner := bufio.NewScanner(conn)

	for {
		fmt.Print("> ")
		inputScanner.Scan()
		msg := []byte(inputScanner.Text() + "\n")
		if string(msg) == "exit\n" {
			break
		}
		conn.Write(msg)
		for responseScanner.Scan() {
			line := responseScanner.Text()
			if line == "---END---" {
				break
			}
			fmt.Println(line)

		}

		/*
			var resp []byte
			resp, err = io.ReadAll(conn)
			if err != nil {
				println("Network error:", err.Error())

			}
			fmt.Println(string(resp))
			resp = []byte("test")
			fmt.Println(string(resp))
		*/
	}

	conn.Close()
	fmt.Println("Connection closed")

}

func moreData(conn *net.TCPConn) bool {
	conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
	defer conn.SetReadDeadline(time.Time{}) // clear deadline after check
	buf := make([]byte, 1)
	n, err := conn.Read(buf)
	return n > 0 && err == nil
}
