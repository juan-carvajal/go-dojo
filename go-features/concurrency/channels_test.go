package concurrency

import "fmt"

// Example_readingFromClosedChannel shows how you can read from closed channels without a problem,
// but it can be confusing, so using the second return variable to check if the channel is open is recommended.
func Example_readingFromClosedChannel() {
	ch := make(chan int)
	close(ch)

	v := <-ch
	fmt.Println(v) // It will allow to read zero-values from closed channel.

	v, ok := <-ch
	fmt.Println(v, ok) // Using the second return variable is recommended to know if you are reading a real value.
	// Output:
	//0
	//0 false
}

// Example_readingFromChannelInForLoop shows how to use the `range` keywork to read from a channel until closed safely.
func Example_readingFromChannelInForLoop() {
	ch := make(chan int)
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch { // The loop will continue until the channel is closed and all the values have been read.
		fmt.Println(v)
	}
	// Output:
	//0
	//1
	//2
	//3
	//4
}
