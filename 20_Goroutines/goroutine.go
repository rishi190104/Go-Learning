//A goroutine is a lightweight thread managed by the Go runtime.

package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task", id)
}

func main() {

	//now without the go run it will in a synchronous way so till the time the further program stays unexecutable

	for i := 0; i < 10; i++ {
		// task(i)
	}

	//so to make the program multithreading we use go routines so the program runs on multi thread concurrently

	//for making the function concurrent we just need to write the go keyword before the function
	for i := 0; i < 10; i++ {
		go task(i)
	}

	//the above didn't run because the code started running on another thread concurrently and moved to
	//further code but there is no code further so the main function stop its executing and when the code stop
	//executing and when code execution gets to the end the goroutines also stop

	//Because the main() function finished before the goroutine got a chance to run.******
	//When main() exits, the entire program exits, including all goroutines.******

	//we can add here a small sleep time

	//the output is not in the sequence because everthing runs concurrently
	//The order is not guaranteed.
	//The Go scheduler decides when each goroutine runs.

	//we do this with anonymous function as well
	//anonymous goroutines

	for i := 0; i <= 10; i++ {
		func() {
			fmt.Println(i)
		}()
	} //this is also a example of a closures

	time.Sleep(time.Second * 2)

	//How goroutines work internally

	//if a computer has a 4cpu cores
	//thread 1
	//thread 2
	//thread 3
	//thread 4

	//each thread is relatively expensive

	//goroutines are much lighter

	//Thread 1
	//Go routine 1
	//Go routine 2
	//Go routine 3
	//Go routine 4
	//Go routine 5
	//Go routine 6

	//The Go runtime schedules thousands (or even more) goroutines across a smaller number of operating system threads.
	//This is why goroutines are lightweight.

}
