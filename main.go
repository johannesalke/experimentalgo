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
	experiment_command_pipelineing()

	fmt.Print("\n")

	time.Sleep(time.Second)
	fmt.Print("|experiment concluded|\n")

}
