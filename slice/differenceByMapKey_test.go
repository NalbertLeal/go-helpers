package slice

import (
    "testing"
)

func TestDiffenrenceByMapKey(t *testing.T) {
    arr1 := []map[string]int{ {"x": 1, "y": 1}, {"x": 2, "y": 2} }
    arr2 := []map[string]int{ {"x": 1, "y": 4} }

    keyComp := []string{ "x" }

    result := DifferenceByMapKey(arr1, arr2, keyComp)
    for _, v := range result {
        if v["x"] == 1 {
            t.Errorf("The function DiffenrenceByMapKey didnt filtered the map correct")
            return
        }
    }
}
