package main

import "fmt"

func partition(v []int, low int, high int) int {
    pivot := v[high]

    // Index of smaller element and indicates 
    // the right position of pivot found so far
    var i int
    i = low - 1

    // Traverse arr[low..high] and move all smaller
    // elements on left side. Elements from low to 
    // i are smaller after every iteration
    for j := low; j <= high - 1; j++ {
        if v[j] < pivot {
            i++;
            v[i], v[j] = v[j], v[i]
        }
    }

    // Move pivot after smaller elements and
    // return its position
    v[i + 1], v[high] = v[high], v[i + 1]
    return i + 1
}


func quickSort(v []int, low int, high int) []int {
    if low < high {
      
        // pi is the partition return index of pivot
        pi := partition(v, low, high)

        // Recursion calls for smaller elements
        // and greater or equals elements
        quickSort(v, low, pi - 1)
        quickSort(v, pi + 1, high)
    }

    return v
}

func main() {
    v := []int{2, 6, 1, 8, 6, 7}
    fmt.Println(v)
    fmt.Println(quickSort(v, 0, len(v)-1))
}
