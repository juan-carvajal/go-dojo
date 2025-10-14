package interfaces

import "fmt"

// Example_interfaceNilComparison shows how nil interface comparisons work.
// An interface is nil only when both its type and value are nil.
func Example_interfaceNilComparison() {
	var i interface{}
	fmt.Println(i == nil) // true: both type and value are nil

	var p *int
	i = p
	fmt.Println(i == nil) // false: type is *int, value is nil
	fmt.Println(p == nil) // true: pointer itself is nil

	// Output:
	//true
	//false
	//true
}

// Example_interfaceEquality shows basic interface equality comparisons.
// Two interfaces are equal if they have identical dynamic types and equal dynamic values.
func Example_interfaceEquality() {
	var a, b interface{}

	a = 42
	b = 42
	fmt.Println(a == b) // true: same type (int) and same value (42)

	a = "hello"
	b = "hello"
	fmt.Println(a == b) // true: same type (string) and same value

	a = 42
	b = "42"
	fmt.Println(a == b) // false: different types (int vs string)

	// Output:
	//true
	//true
	//false
}

// Example_interfaceComparablePanic shows that comparing interfaces containing
// non-comparable types causes a runtime panic.
func Example_interfaceComparablePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()

	var a, b interface{}
	a = []int{1, 2, 3}
	b = []int{1, 2, 3}

	// This will panic because slices are not comparable
	fmt.Println(a == b)

	// Output:
	//panic: runtime error: comparing uncomparable type []int
}

// Example_interfacePointerComparison shows how pointer values are compared
// when stored in interfaces.
func Example_interfacePointerComparison() {
	x := 42
	y := 42

	var a, b interface{}
	a = &x
	b = &x
	fmt.Println(a == b) // true: same pointer address

	a = &x
	b = &y
	fmt.Println(a == b) // false: different pointer addresses

	// Output:
	//true
	//false
}

// Example_interfaceTypeAssertion shows how type assertions interact with
// comparisons and nil checking.
func Example_interfaceTypeAssertion() {
	var i interface{} = (*int)(nil)

	fmt.Println(i == nil) // false: interface holds a type

	p, ok := i.(*int)
	fmt.Println(ok)       // true: type assertion succeeds
	fmt.Println(p == nil) // true: the underlying pointer is nil

	// Output:
	//false
	//true
	//true
}

// Example_interfaceStructComparison shows how structs are compared when
// stored in interfaces.
func Example_interfaceStructComparison() {
	type Point struct {
		X, Y int
	}

	var a, b interface{}
	a = Point{1, 2}
	b = Point{1, 2}
	fmt.Println(a == b) // true: same type, same field values

	a = Point{1, 2}
	b = Point{2, 1}
	fmt.Println(a == b) // false: different field values

	// Output:
	//true
	//false
}

// Example_interfaceDifferentTypes shows comparison behavior with different
// concrete types stored in interfaces.
func Example_interfaceDifferentTypes() {
	type MyInt int

	var a, b interface{}
	a = int(42)
	b = MyInt(42)
	fmt.Println(a == b) // false: different types (int vs MyInt)

	a = 42
	b = 42.0
	fmt.Println(a == b) // false: different types (int vs float64)

	// Output:
	//false
	//false
}

// Example_interfaceEmptyVsNil shows the distinction between empty interfaces
// holding values versus being nil.
func Example_interfaceEmptyVsNil() {
	var i interface{}
	fmt.Println(i == nil) // true

	i = 0
	fmt.Println(i == nil) // false: holds int(0)
	fmt.Println(i == 0)   // true: value is 0

	i = ""
	fmt.Println(i == nil) // false: holds string("")
	fmt.Println(i == "")  // true: value is ""

	// Output:
	//true
	//false
	//true
	//false
	//true
}
