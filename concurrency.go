package main

import (
	"fmt"
	"os"
	"time"
)

func experiment_gorutine() { //For a program like this, having a time.Sleep() call is essential, or else the program will exit before the goroutine have a chance to run.
	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")

	time.Sleep(time.Second)

}

func experiment_channel() { //As it happens, you can use channels to ferry information and events between two background goroutines.

	c := make(chan int)
	go func() {
		for i := 1; i < 6; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for i := range c { //This will remain active until the channel closes. If there is no closing statement, this may cause a runtime error.
			fmt.Print(i, "\n")
		}
	}()

	//--------------------------------------------

	sc := make(chan struct{})
	go experiment_channel_signalunblock(sc)
	<-sc //This blocks until the channel is closed.
	fmt.Print("channel closed\n")
}

func experiment_channel_signalunblock(sc chan<- struct{}) { //Using the closing of a channel as a signal that a certain task is finished and other threads can continue working.

	time.Sleep(5 * time.Second)
	fmt.Print("closing channel...\n")
	close(sc)
}

func experiment_select() {

	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() {
		for i := 1; i < 6; i++ {
			ch1 <- i
			time.Sleep(2 * time.Second)
		}
		//close(ch1)
	}()
	go func() {
		for i := 1; i < 6; i++ {
			ch2 <- "-o-"
			time.Sleep(3 * time.Second)
		}
		//close(ch2)

	}()

	for {
		select {
		case i := <-ch1:
			fmt.Printf("Received from ch1: %d\n", i)
		case s := <-ch2:
			fmt.Printf("Received from ch2: %s\n", s)
		case <-time.After(5 * time.Second): // The time.After(t) case happens if none of the other channels supply an item within duration t.
			os.Exit(0) // It can be used as an imperfect proxy for the channels being done with their purpose, so long they weren't closed (which would send infinite 0 signals)
		}

	} //Alternate method for waiting for a group of go-routines to finish: https://yourbasic.org/golang/wait-for-goroutines-waitgroup/
}

func experiment_goroutinekill() {
	quit := make(chan bool)
	ch := make(chan int)
	go func() {
		for i := 1; ; i++ {
			select {
			case <-quit:
				return
			default:
				ch <- i
				time.Sleep(1 * time.Second)
			}
		}
	}()

	for i := range ch {
		fmt.Printf("Number: %d\n", i)
		if i == 7 {
			quit <- true
			break
		}

	}

}
