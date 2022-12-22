package array

import (
    "testing"
    "fmt"
)

func TestDifference(t *testing.T) {
    fmt.Println("OKOKOK")
    a := []int{1, 2}
    b := []int{2, 1}

    result := Difference(a, b)
    fmt.Println(result)
    for _, v := range result {
        if v != 3 {
            t.Errorf("Difference didnt add 3")
        }
    }
}
