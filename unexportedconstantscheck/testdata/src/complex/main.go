package complex

import "fmt"

// rate the case
const rate = 0 // want `unexported constant "rate" should be prefixed with _`

const one, two = "ONE", "TWO" // want `unexported constant "one" should be prefixed with _` `unexported constant "two" should be prefixed with _`

func main() {
	var a = 1

	fmt.Print(a)
}
