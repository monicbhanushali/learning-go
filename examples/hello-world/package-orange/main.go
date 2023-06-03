package orange

// since we have named the package something else then main it wouldn't run and instead would only compile
// the code in this can be used by other as a module

import "fmt"

func greet() {
	fmt.Println("Hi there....this function won't execute since it is not main :)")
}
