package consts

import "fmt"

// Example_iotaIdentity shows the basic usage of iota.
func Example_iotaIdentity() {
	const (
		c0 = iota // c0 == 0
		c1 = iota // c1 == 1
		c2 = iota // c2 == 2
	)
	fmt.Println(c0)
	fmt.Println(c1)
	fmt.Println(c2)
	// Output:
	//0
	//1
	//2
}

// Example_iotaSkip shows that you can skip iota "rows", meaning not use the iota in that row and override the with something else.
// The count will be incremented as normal in the next row.
func Example_iotaSkip() {
	const (
		a = 1 << iota // a == 1  (iota == 0)
		b = 1 << iota // b == 2  (iota == 1)
		c = 3         // c == 3  (iota == 2, unused)
		d = 1 << iota // d == 8  (iota == 3)
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	// Output:
	//1
	//2
	//3
	//8
}

// Example_iotaCustomIncrement shows that you can use a custom expression to increment, for example a combination of bitwise + multiplication.
func Example_iotaCustomIncrement() {
	type byteSize float64

	const (
		_           = iota // ignore first value by assigning to blank identifier
		KB byteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	// Output:
	//1024
	//1.048576e+06
	//1.073741824e+09
	//1.099511627776e+12
	//1.125899906842624e+15
	//1.152921504606847e+18
}

// Example_iotaMixedTypes shows that you can define mixed types with iota expressions.
func Example_iotaMixedTypes() {
	const (
		u         = iota * 42 // u == 0     (untyped integer constant)
		v float64 = iota * 42 // v == 42.0  (float64 constant)
		w         = iota * 42 // w == 84    (untyped integer constant)
	)
	fmt.Println(u)
	fmt.Println(v)
	fmt.Println(w)
	// Output:
	//0
	//42
	//84
}
