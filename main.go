package main

import (
	"fmt"
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
	experiment_goroutinekill()
	fmt.Print("\n")

	time.Sleep(time.Second)
	fmt.Print("|experiment concluded|\n")

}
