package structs

import "fmt"

type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

// Example_methodSets shows how method sets behave at a basic level.
// Some important considerations from https://go.dev/wiki/MethodSets
//
// Method Sets: A type may have a method set associated with it. The method set of an interface type is its interface.
// The method set of any other named type T consists of all methods with receiver type T.
// The method set of the corresponding pointer type *T is the set of all methods with receiver *T or T
// (that is, it also contains the method set of T). Any other type has an empty method set.
// In a method set, each method must have a unique name.
//
// Calls: A method call x.m() is valid if the method set of (the type of) x contains m and the argument list can
// be assigned to the parameter list of m. If x is addressable and &x’s method set contains m,
// x.m() is shorthand for (&x).m().
//
// In practice, this means that for structs at least, `T` pointer receivers will have access to `*T` pointer receivers
// because `T` is addressable, so any call to a method `m()` of `*T` coming from `T` is translated to
// (&x).m()
//
// IMPORTANT: Keep in mind that mixing value and pointer receiver in the same struct is not recommended by Go docs.
func Example_methodSets() {
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)\n", lst, lst.Len())

	// A pointer value
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)\n", plst, plst.Len())

	// Output:
	//[1] (len: 1)
	//&[2] (len: 1)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}
func (p *Person) SetAge(age int) {
	p.Age = age
}

type Singer struct {
	Person // extends Person by embedding it
	works  []string
}

// Example_embeddings shows how field embedding works too, as stated by the Go specs (https://go.dev/ref/spec#Struct_types):
// Given a struct type S and a type name T, promoted methods are included in the method set of the struct as follows:
//
// If S contains an embedded field T, the method sets of S and *S both include promoted methods with receiver T.
// The method set of *S also includes promoted methods with receiver *T.
// If S contains an embedded field *T, the method sets of S and *S both include promoted methods with receiver T or *T.
func Example_embeddings() {
	var gaga = Singer{Person: Person{"Gaga", 30}}
	gaga.PrintName() // Name: Gaga
	gaga.Name = "Lady Gaga"
	(&gaga).SetAge(31)
	(&gaga).PrintName()   // Name: Lady Gaga
	fmt.Println(gaga.Age) // 31
	// Output:
	//Name: Gaga
	//Name: Lady Gaga
	//31
}

// Example_embeddings shows that even tho the Go specs clearly states that `Singer` should not have access to the
// `*Person.SetAge` method of `*Person`, but again, in practice it has access because behind the scenes, the compiler
// will take the address before calling the method.
// https://go101.org/article/type-embedding.html and https://go101.org/article/method.html#call elaborate on the concepts
// described here.*/
func Example_embeddings2() {
	var gaga = Singer{Person: Person{"Gaga", 30}}
	gaga.PrintName() // Name: Gaga
	gaga.Name = "Lady Gaga"
	gaga.SetAge(31)
	gaga.PrintName()      // Name: Lady Gaga
	fmt.Println(gaga.Age) // 31
	// Output:
	//Name: Gaga
	//Name: Lady Gaga
	//31
}
