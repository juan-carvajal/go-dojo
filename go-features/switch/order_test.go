package _switch

import (
	"fmt"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

// Example_switchOrder helps understand the order of execution of switch cases
func Test_switchOrder(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		weekday := time.Now().Weekday()
		switch weekday {
		case time.Saturday:
			require.FailNow(t, "saturday")
		case time.Sunday:
			require.FailNow(t, "sunday")
		default:
			require.Equal(t, weekday, time.Friday)
		}
	})
}

// Example_switchOrder also shows the linear order of the switch case without a default case. Also, we can see that the fallthrough keyword can be used to continue running the next case.
func Example_switchOrder() {
	switch 2 {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}

	switch 2 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}

	// Output:
	// 2
	// 2
	// 3
}

func Foo(n int) int {
	fmt.Println(n)
	return n
}

// ExampleFoo shows the order of execution from left to right and top to bottom, and also shows the fallthrough keyword.
func ExampleFoo() {
	switch Foo(2) {
	case Foo(1), Foo(2), Foo(3):
		fmt.Println("First case")
		fallthrough
	case Foo(4):
		fmt.Println("Second case")
	}
	// Output:
	// 2
	// 1
	// 2
	// First case
	// Second case
}

// Example_break shows how to use the explicit break. Golang's switch cases break implicitly.
func Example_break() {
	argv := []any{"cat"}
	switch argv[0] {
	case "echo":
		fmt.Print(argv[1:]...)
	case "cat":
		if len(argv) <= 1 {
			fmt.Println("Usage: cat <filename>")
			break
		}
	default:
		fmt.Println("Unknown command; try 'echo' or 'cat'")
	}
	// Output: Usage: cat <filename>
}
