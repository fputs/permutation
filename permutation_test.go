package permutation

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	a := []int{1, 2, 3, 4}
	p := Permutations(a)
	fmt.Println(p)
}
