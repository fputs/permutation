/*
Package permutation implements Heap's algorithm for generating permutations of
lists. It uses Go 1.18's new generics feature to provide a generic interface
*/
package permutation

// GenSlice is a generic slice of data
type GenSlice[T any] []T

/*
Permutations uses Heap's Algorithm to generate a list of all possible
permutations of the provided list. Most slices will automatically coerce
their type into a GenSlice[T] with no casting required
*/
func Permutations[T any](arr GenSlice[T]) []GenSlice[T] {
	var helper func(GenSlice[T], int)
	var res []GenSlice[T]

	helper = func(arr GenSlice[T], n int) {
		if n == 1 {
			var tmp GenSlice[T]
			for _, i := range arr {
				tmp = append(tmp, i)
			}
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
