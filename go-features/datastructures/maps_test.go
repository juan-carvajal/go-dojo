package datastructures

import "fmt"

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

// Example_allowedMapOperations shows what operations are allowed on a nil map.
func Example_allowedMapOperations() {
	var m map[string]string // nil map
	delete(m, "")           // allowed on nil map

	for k, v := range m { // iteration over key-values, allowed on nil map
		fmt.Println(k, v)
	}

	k, ok := m["a"] // explicit read, allowed on nil map
	fmt.Println(k, ok)

	defer recoverFromPanic()
	m["a"] = "b" // write operations not allowed on nil map, this will panic

	// Output:
	// false
	//Recovered from panic: assignment to entry in nil map
}
