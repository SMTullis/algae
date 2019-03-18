package sort

import (
    "fmt"
    "testing"
)

func Test_MergeSort(t *testing.T) {
    sorted := MergeSort([]int{3,5,7,4,10,6,8,2,9,1})
    expect := [10]int{1,2,3,4,5,6,7,8,9,10}
    for index := range sorted {
        if sorted[index] != expect[index] {
            fmt.Printf("Not sorted: %v != %v", sorted[index], expect[index])
        }
    }
}

func Test_PigeonSort(t *testing.T) {
    sorted := PigeonSort([]string{"six", "five", "one", "five", "five", "six", "one", "five", "five",  "one", "five"})
    expect := [11]string{"five", "five", "five", "five", "five", "five", "one", "one", "one", "six", "six"}
    for index := range sorted {
        if sorted[index] != expect[index] {
            fmt.Printf("Not sorted: %v != %v", sorted[index], expect[index])
        }
    }
}
