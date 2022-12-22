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

    result := DiffenrenceByIterrator(arr1, arr2, iteratee)

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

func TestDiffenrenceByMapKey(t *testing.T) {
    arr1 := []map[string]int{ {"x": 1, "y": 1}, {"x": 2, "y": 2} }
    arr2 := []map[string]int{ {"x": 1, "y": 4} }

    keyComp := []string{ "x" }

    result := DiffenrenceByMapKey(arr1, arr2, keyComp)
    for _, v := range result {
        if v["x"] == 1 {
            t.Errorf("The function DiffenrenceByMapKey didnt filtered the map correct")
            return
        }
    }
}
