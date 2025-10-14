package concurrency

import (
	"context"
	"fmt"
	"time"
)

func sendUntilCancelled(ctx context.Context, ch chan<- int, tickerDuration time.Duration) {
	defer close(ch)
	t := time.NewTicker(tickerDuration)
	defer t.Stop()
	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			ch <- i
		}
		i += 1
	}
}

func sendNTimes(ctx context.Context, ch chan<- int, tickerDuration time.Duration, nTimes int) {
	defer close(ch)
	t := time.NewTicker(tickerDuration)
	defer t.Stop()
	for i := range nTimes {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			ch <- i
		}
	}
}

// Example_makingBlockingCalls shows how to make a blocking call non-blocking by calling inside a goroutine
func Example_makingBlockingCalls() {
	ch := make(chan int, 1) // The one-sized buffered channel is useful to avoid blocking the producer if the reader is slow.
	go sendNTimes(context.Background(), ch, time.Millisecond*50, 5)
	for v := range ch {
		fmt.Println(v)
	}

	// Output:
	//0
	//1
	//2
	//3
	//4
}

// Example_makingBlockingCallsWithCancelledContext shows how to make a blocking call non-blocking by calling inside a goroutine
// this non-blocking producer functions can be controlled with a context
func Example_makingBlockingCallsWithCancelledContext() {
	ch := make(chan int, 1) // The one-sized buffered channel is useful to avoid blocking the producer if the reader is slow.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()
	go sendUntilCancelled(ctx, ch, time.Second*10)
	for v := range ch {
		fmt.Println(v) // This wil never run because the context is timeouts before any message is sent.
	}

	// Output:
}
