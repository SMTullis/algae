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
