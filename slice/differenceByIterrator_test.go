package slice

import (
    "testing"
)

func TestDiffenrenceByIterrator(t *testing.T) {
    iteratee := func(v float64) int {
        return int(v)
    }
    arr1 := []float64{1.3, 2.3, 3.3, 4.3, 5.3, 6.3, 7.3}
    arr2 := []float64{4.7, 2.7}

    result := DifferenceByIterrator(arr1, arr2, iteratee)

    contains4 := false
    contains2 := false
    for _, v := range result {
        if int(v) == 4 {
            contains4 = true
            continue
        }
        if int(v) == 2 {
            contains2 = true
            continue
        }
    }
    if contains2 || contains4 {
        t.Errorf("Result of DiffenrenceByIterrator dont contains 2 or 4")
    }
}