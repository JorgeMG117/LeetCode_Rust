package main

import (
    "fmt"
)

func merge(arr []int, left, mid, right int) {
    n1 := mid - left + 1
    n2 := right - mid

    // Create temp vectors
    L := make([]int, n1)
    R := make([]int, n2)

    // Copy data to temp vectors L[] and R[]
    for i := 0; i < n1; i++ {
        L[i] = arr[left + i]
    }
    for j := 0; j < n2; j++ {
        R[j] = arr[mid + 1 + j]
    }

    i, j := 0, 0
    k := left

    // Merge the temp vectors back 
    // into arr[left..right]
    for i < n1 && j < n2 {
        if (L[i] <= R[j]) {
            arr[k] = L[i]
            i++;
        } else {
            arr[k] = R[j]
            j++;
        }
        k++;
    }

    // Copy the remaining elements of L[], 
    // if there are any
    for i < n1 {
        arr[k] = L[i]
        i++
        k++
    }

    // Copy the remaining elements of R[], 
    // if there are any
    for j < n2 {
        arr[k] = R[j]
        j++
        k++
    }
}

// begin is for left index and end is right index
// of the sub-array of arr to be sorted
func mergeSort(arr []int, left int, right int) {
    if (left >= right) {
        return
    }

    mid := left + (right - left) / 2
    mergeSort(arr, left, mid)
    mergeSort(arr, mid + 1, right)
    merge(arr, left, mid, right)
}

func main() {
    v := []int{2, 6, 1, 8, 6, 7}
    fmt.Println(v)
    mergeSort(v, 0, len(v)-1)
    fmt.Println(v)
}
