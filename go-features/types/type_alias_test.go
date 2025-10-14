package types

import (
	"fmt"

	"github.com/juan-carvajal/go-dojo/go-features/types/subtype"
)

type aString string

type bString = aString

func (b bString) String() string {
	return string(b)
}

func (b bString) Len() int {
	return len(b)
}

type cString = subtype.SString

// Example_typeAlias shows how method sets behave for type aliases.
//
// Please observe that this will not compile, because cString is an external type.
//
//	 func (c cString) Len(){
//		 return len(c)
//	 }
func Example_typeAlias() {
	a := aString("hello a")
	fmt.Println(a.String(), a.Len()) // `String()` and `Len()` is available here, even tho they are not directly attached to aString

	b := bString("hello b")
	fmt.Println(b.String(), b.Len())

	c := cString("hello c")
	fmt.Println(c.String()) // Methods from external aliased type are available.

	// Output:
	//hello a 7
	//hello b 7
	//hello c
}

func Example_typeAliasComparison() {
	a := aString("hello")
	fmt.Println(a.String(), a.Len()) // `String()` and `Len()` is available here, even tho they are not directly attached to aString

	b := bString("hello")
	fmt.Println(b.String(), b.Len())

	c := cString("hello c")
	fmt.Println(c.String())

	fmt.Println(a == b) // can be compared this way because they are considered the same type and underlying type is string (which is comparable)
	// fmt.Println(c == b) This does not even compile, because they are not considered the same type.

	// Output:
	//hello 5
	//hello 5
	//hello c
	//true
}
