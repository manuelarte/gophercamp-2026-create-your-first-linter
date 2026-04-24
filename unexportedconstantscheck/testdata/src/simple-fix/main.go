package simple

const myconstant = "myconstant" // want `unexported constant "myconstant" should be prefixed with _`

const Rate = 0

const errNotFound = "not found"

const _errNotFound = "not found" // want `unexported err constant "_errNotFound" should not be prefixed with _`

const (
	group        = "group" // want `unexported constant "group" should be prefixed with _`
	Of           = "Of"
	errConstants = "error"
)

func aFunction(input int) int {
	output := input * 2
	return output
}
