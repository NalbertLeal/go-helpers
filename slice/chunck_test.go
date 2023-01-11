package slice

import (
	"testing"
)

func TestChunckDefaultChunckSize(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	result := Chunck(arr)
	if len(result) != 1 {
		t.Errorf("Wrong number of chuncks. Got %d chuncks.", len(result))
	}
}

func TestChunck3Chuncks(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	result := Chunck(arr, 3)
	if len(result) != 3 {
		t.Errorf("Wrong number of chuncks. Got %d chuncks.", len(result))
	}
}

func TestChunckSameNumberChuncksAndElements(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	result := Chunck(arr, 18)
	if len(result) != 10 {
		t.Errorf("Wrong number of chuncks. Got %d chuncks.", len(result))
	}
}
