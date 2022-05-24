package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//Empty struct is used when we want to save memories.
	//This is because they do not consume any memory for the values. The syntax for an empty struct is:
	emptyStruct := struct{}{}
	var emptyInterface interface{}

	// The size of empty struct would return 0 when using:
	fmt.Println(unsafe.Sizeof(emptyStruct))    // -> 0
	fmt.Println(unsafe.Sizeof(emptyInterface)) // -> 16

	// The important use of empty struct is to show the developer that we do not have any value.
	// The purpose is purely informational. Some of the examples where the empty struct is useful
	// are as follows:

	// 1- While implementing a data set: We can use the empty struct to implement a dataset.
	// Consider an example as shown below.

	mapVariable := make(map[string]struct{})

	for _, val := range []string{"field1", "field2", "field3"} {
		mapVariable[val] = emptyStruct
	}
	// Here, we are initializing the value of a key to an empty struct
	// and initializing the map_obj to an empty struct.
	fmt.Println(mapVariable) // -> map[field1:{} field2:{} field3:{}]

	// 2- When a channel needs to send a signal of an event without the need for sending any data.
	// From the below piece of code, we can see that we are sending a signal using sending empty struct
	// to the channel which is sent to the workerRoutine.

	goroutinFunction := func(ch chan struct{}) {
		// Receive message from main program.
		<-ch
		fmt.Println("signal Recived from main")

		// Send a message to the main program.
		close(ch)
	}

	//Create channel
	emptyStructChan := make(chan struct{})

	go goroutinFunction(emptyStructChan)

	// Send signal to worker goroutine
	emptyStructChan <- struct{}{}

	// Receive a message from the workerRoutine.
	<-emptyStructChan
	fmt.Println("signal Recived from GoRoutin")

}
