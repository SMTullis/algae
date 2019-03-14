package sort

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
    for _, value := range unsorted {
        holes[value] = append(holes[value], value)
    }
    var out []string
    for _, each := range holes {
        for _, x := range each {
            out = append(out, x)
        }
    }
    return out
}
