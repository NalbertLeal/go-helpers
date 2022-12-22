package array

import (
    "testing"
)

func TestConcat(t *testing.T) {
    a := []int{1, 2, 3}
    b := []int{4, 5, 6}

    result := Concat(a, b)
    for i, v := range a {
        if v != result[i] {
            t.Errorf("Elements of the 1st array is different")
            return
        }
    }

    for i, v := range b {
        if v != result[len(a)+i] {
            t.Errorf("Elements of the 1st array is different")
            return
        }
    }
}
