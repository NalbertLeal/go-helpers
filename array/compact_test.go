package array

import (
    "testing"
)

func TestCompactToInt(t *testing.T) {
    arr := []int{0, 1, 2, 3, 0, 4, 5, 0, 6, 0, 7, 8, 0, 9}

    result := Compact(arr)
    for _, v := range result {
        if v == 0 {
            t.Errorf("Failed to remove 0 from integer array.")
            return
        }
    }
}

type A struct {
    a int
    b string
}

func TestCompactToStruct(t *testing.T) {
    a1 := &A{ 1, "a1" }
    a2 := &A{ 2, "a2" }

    arr := []*A{a1, nil, a2}

    result := Compact(arr)
    for _, v := range result {
        if v == nil {
            t.Errorf("Failed to remove 0 from integer array.")
            return
        }
    }
}
