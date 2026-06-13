package main

import (
	"fmt"
	//"github.com/johannesalke/experimentalgo/stdlib"
	"time"
)

func main() {

	//experiment_interfaces()
	//experiment_reader()
	//experiment_readwriter()
	//experiment_offsets()
	//experiment_embedding()
	//experiment_gorutine()
	//experiment_channel()
	//experiment_select()
	//experiment_goroutinekill()
	//experiment_repeater()
	//experiment_timedcall()
	//<>experiment_mutex()
	//experiment_tcp_test()
	//experiment_tcp_writer()
	//experiment_functiontype()	//Followup!
	//experiment_tcp_readwriter() //Followup!

	//stdlib.Experiment_scanner()
	//stdlib.Experiment_customSplit_scanner()
	//experiment_command_input()
	//experiment_command_input_pipe()
	//experiment_command_pipelineing()
	//experiment_command_shell_from_file()
	//experiment_command_send_output_to_another_terminal()

	//go experiment_goroutine_limit_A()
	//time.Sleep(time.Second)
	//experiment_goroutine_limit_B()

	//experiment_format_preexisting_string()

	//experiment_channels_closing()
	//experiment_printing_runes()

	//experiment_reflect_type_and_kind()

	//experiment_generic_function([]int{1, 2, 4, 5})
	//experiment_constrained_generic(struct256(0b01010101))
	//experiment_union_constrained_generic(intProxy(2))
	//experiment_union_constrained_generic(float64(2.1241))
	experiment_generic_interface()

	fmt.Print("\n")

	time.Sleep(time.Second)
	fmt.Print("|experiment concluded|\n")

}
