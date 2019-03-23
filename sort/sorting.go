package sorting

import "sort"

func MergeSort(arr []int) []int {
    length := len(arr)
    if length == 1 {
        return arr
    }

    midpoint := int(length / 2)
    return merge(
        MergeSort(arr[:midpoint]),
        MergeSort(arr[midpoint:]))
}

func merge(front, back []int) []int {
    out := make([]int, 0)
    outerLoop:
    for _, x := range front {
        innerLoop:
        for _, y := range back {
            if x < y {
                out = append(out, x)
                continue outerLoop
            } else {
                out = append(out, y)
                continue innerLoop
            }
        }
    }
    return out
}

func PigeonSort(unsorted []string) []string {
    holes := make(map[string][]string)
    sorted_keys := make([]string, 0)
    for _, value := range unsorted {
        holes[value] = append(holes[value], value)
        if !isInArray(sorted_keys, value) {
            sorted_keys = append(sorted_keys, value)
        }
    }
    sort.Strings(sorted_keys)
    var out []string
    for _, each := range sorted_keys {
        for _, x := range holes[each] {
            out = append(out, x)
        }
    }
    return out
}

func isInArray(array []string, value string) bool {
    for _, each := range array {
        if each == value { return true }
    }
    return false
}
