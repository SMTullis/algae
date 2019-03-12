package algae_test

import (
    "fmt"
    "testing"
    "github.com/SMTullis/algae/sort"
)

func Test_MergeSort(t *testing.T) {
    sorted := algae.MergeSort([]int{3,5,7,4,10,6,8,2,9,1})
    expect := [10]int{1,2,3,4,5,6,7,8,9,10}
    for index := range sorted {
        if sorted[index] != expect[index] {
            panic("Not sorted")
        }
    }
}
