package array

import (
	"math"
)

func chunck[data any](arr []data, others ...int) [][]data {
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
