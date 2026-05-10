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

// It turns out that you can hand a file with \n chars to a sh process and it will be interpreted as seperate commands. Basically just the same as a shell script though.
func experiment_command_shell_from_file() {
	cmd := exec.Command("sh")
	stdin, _ := cmd.StdinPipe()
	cmd.Stdout = os.Stdout
	file, _ := os.Open("assets/commands.txt")
	file.WriteTo(stdin)
	cmd.Start()

}

// By getting the file identifier of another terminal's output with tty, you can redirect cmd.Stdout to another terminal process entirely
// I also checked whether you can write to another terminal's stdin, but that was wisely blocked as a security measure, unless you explicitly unblock it
func experiment_command_send_output_to_another_terminal() {
	out, _ := os.OpenFile(os.Args[1], os.O_WRONLY, 0)

	defer out.Close()

	cmd := exec.Command("sh")
	stdin, _ := cmd.StdinPipe()
	cmd.Stdout = out
	file, _ := os.Open("assets/commands.txt")
	defer file.Close()
	file.WriteTo(stdin)
	cmd.Start()

}
