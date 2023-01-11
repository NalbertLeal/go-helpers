package slice

import (
	"math"
)

// Divide an slice elements into chuncks, each chunck is represented by a
// integer number that idicate the number of elements to be inserted into
// that chunck. Example:
//		Chunck(someSlice, []int{3, 9, 5})
// the example above divide the elements of slice **someSlice** into at least
// 3 chunks, the first have 3 elements, the second 9 and the third 5. If the
// slice **someSlice** have more elements than the number alocated (in the
// example 3+9+5 elements) then all the non-allocated elements are inserted
// into another chunck (in the example a fourth chunck).
func Chunck[data any](arr []data, others ...int) [][]data {
	var size int
	if len(others) > 0 {
		size = others[0]
	} else {
		return [][]data{arr[0:]}
	}

	elemsPerChunck := int(math.Ceil(float64(len(arr)) / float64(size)))
	result := [][]data{}
	i := 0
	for i = 0; i+elemsPerChunck < len(arr); i = i + elemsPerChunck {
		result = append(result, arr[i:i+elemsPerChunck])
	}
	if i < len(arr) {
		result = append(result, arr[i-elemsPerChunck:])
	}
	return result
}
