package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func experiment_command_input() {
	cmd := exec.Command("cat")
	cmd.Stdin = strings.NewReader("'this small snippet' > test.txt\n")
	cmd.Stdout = os.Stdout
	cmd.Start()
}

func experiment_command_input_pipe() {
	cmd := exec.Command("cat")

	cmd.Stdout = os.Stdout
	stdin, _ := cmd.StdinPipe()

	fmt.Fprintf(stdin, "'this small snippet'\n > test.txt")
	cmd.Start()
}

func experiment_command_pipelineing() { // Can you link commands as if by | pipeline? Yes!
	cat := exec.Command("cat")
	grep := exec.Command("grep", "line")

	file, _ := os.Open("assets/test.txt")
	cat.Stdin = file

	grep.Stdin, _ = cat.StdoutPipe()
	grep.Stdout = os.Stdout

	cat.Start()
	grep.Start()

}
