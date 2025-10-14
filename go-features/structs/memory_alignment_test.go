package structs

import (
	"fmt"
	"unsafe"
)

// Example_structMemoryAlignment shows how the memory alignment constraints will affect struct size at runtime.
// Paraphrasing this [Memory Layout Article], Golang has some guarantees when it comes to storing fields in structs:
//
// 1. Alignment guarantees:
//
//	type                      alignment guarantee
//	------                    ------
//	bool, uint8, int8         1
//	uint16, int16             2
//	uint32, int32             4
//	float32, complex64        4
//	arrays                    depend on element types
//	structs                   depend on field types
//	other types               size of a native word
//
// 2. Type size guarantees (more details for other types in [Value copy cost article]):
//
//	type                    size in bytes
//	------                  ------
//	uint8, int8             1
//	uint16, int16           2
//	uint32, int32, float32  4
//	uint64, int64           8
//	float64, complex64      8
//	complex128              16
//	uint, int               implementation-specific,generally 4 on 32-bit architectures, and 8 on 64-bit architectures.
//	uintptr                 implementation-specific, large enough to store the uninterpreted bits of a pointer value.
//
// Additionally and according [Alignment and size guarantees], to following minimal alignment properties are guaranteed:
//
//  1. For a variable x of any type: unsafe.Alignof(x) is at least 1.
//  2. For a variable x of struct type: unsafe.Alignof(x) is the largest of all the values unsafe.Alignof(x.f) for each field f of x, but at least 1.
//  3. For a variable x of array type: unsafe.Alignof(x) is the same as the alignment of a variable of the array's element type.
//
// In practice, this means that total struct sizes are not a simple sum of its parts in all cases.
//
// [Memory Layout Article]: https://go101.org/article/memory-layout.html
// [Value copy cost article]: https://go101.org/article/value-copy-cost.html#value-sizes
// [Alignment and size guarantees]: https://golang.org/ref/spec#Size_and_alignment_guarantees
func Example_structMemoryAlignment() {
	type t1 struct {
		a int8

		// On 64-bit architectures, to make field b
		// 8-byte aligned, 7 bytes need to be padded
		// here. On 32-bit architectures, to make
		// field b 4-byte aligned, 3 bytes need to be
		// padded here.

		b int64
		c int16

		// To make the size of type T1 be a multiple
		// of the alignment guarantee of T1, on 64-bit
		// architectures, 6 bytes need to be padded
		// here, and on 32-bit architectures, 2 bytes
		// need to be padded here.
	}
	// The size of T1 is 24 (= 1 + 7 + 8 + 2 + 6)
	// bytes on 64-bit architectures and is 16
	// (= 1 + 3 + 8 + 2 + 2) on 32-bit architectures.

	type t2 struct {
		a int8

		// To make field c 2-byte aligned, one byte
		// needs to be padded here on both 64-bit
		// and 32-bit architectures.

		c int16

		// On 64-bit architectures, to make field b
		// 8-byte aligned, 4 bytes need to be padded
		// here. On 32-bit architectures, field b is
		// already 4-byte aligned, so no bytes need
		// to be padded here.

		b int64
	}
	// The size of T2 is 16 (= 1 + 1 + 2 + 4 + 8)
	// bytes on 64-bit architectures, and is 12
	// (= 1 + 1 + 2 + 8) on 32-bit architectures.

	fmt.Println("Size of t1:", unsafe.Sizeof(t1{}))
	fmt.Println("Size of t2:", unsafe.Sizeof(t2{}))
	// Output:
	//Size of t1: 24
	//Size of t2: 16
}

// Example_structMemoryAlignmentOptimized shows how by ordering the fields from largest to the smallest will minimize the
// need for padding between fields, thus reducing the total size of the struct.
func Example_structMemoryAlignmentOptimized() {
	type t1 struct {
		b int64
		c int16
		a int8
	}

	fmt.Println("Size of t1:", unsafe.Sizeof(t1{}))
	// Output:
	//Size of t1: 16
}
