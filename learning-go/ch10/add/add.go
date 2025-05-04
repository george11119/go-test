package add

import "golang.org/x/exp/constraints"

// adds a and b, returns their sum
// [https://www.mathsisfun.com/numbers/addition.html]
func Add[T constraints.Integer](a, b T) T {
	return a + b
}
