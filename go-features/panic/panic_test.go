package panic

import (
	"fmt"
	"sync"
)

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

// Example_recoverFromPanic shows how panics can be created explicitly and recovered using the `recover` function.
func Example_recoverFromPanic() {
	defer recoverFromPanic()
	panic("Oops! Something went wrong!")
	// Output: Recovered from panic: Oops! Something went wrong!
}

// Example_panicStoppingExecution shows how a panic will not allow the program to continue after it has been triggered (except with deferred functions).
func Example_panicStoppingExecution() {
	defer recoverFromPanic()
	panic("initial panic")
	panic("second panic") // This will never run
	// Output: Recovered from panic: initial panic
}

// Example_panicChain shows how panic chains work with defer functions.
func Example_panicChain() {
	defer recoverFromPanic()
	defer func() {
		fmt.Println("inside first defer")
		panic("panic inside first defer")
	}()
	panic("initial panic")

	// Output:
	// inside first defer
	// Recovered from panic: panic inside first defer
}

// Example_panicChain2 shows how the deferred functions up to the point where the program panic initially will always have a chance to run, even if another deferred function panics.
func Example_panicChain2() {
	defer recoverFromPanic()
	defer func() {
		fmt.Println("first")
		panic("panic first")
	}()
	defer func() {
		fmt.Println("second")
		panic("panic second")
	}()
	defer func() {
		fmt.Println("third")
	}()
	panic("original")

	// Output:
	// third
	// second
	// first
	// Recovered from panic: panic first
}

// Example_panicInGoroutine shows code that would crash because the `recover` mechanism works at the goroutine level only. Each goroutine needs to handle it's panic stack.
func Example_panicInGoroutine() {
	//defer recoverFromPanic()
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("Worker goroutine starting...")
	//	panic("Something went wrong in the worker!") // This will cause a panic
	//}()
	//wg.Wait()
	fmt.Println("don't uncomment this, it will crash")
	// Output: don't uncomment this, it will crash
}

// Example_panicInGoroutineGracefully shows how to handle panics inside a goroutine.
func Example_panicInGoroutineGracefully() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer recoverFromPanic()
		fmt.Println("Worker goroutine starting...")
		panic("Something went wrong in the worker!") // This will cause a panic
	}()
	wg.Wait()
	fmt.Println("done!")
	// Output:
	// Worker goroutine starting...
	// Recovered from panic: Something went wrong in the worker!
	// done!
}
