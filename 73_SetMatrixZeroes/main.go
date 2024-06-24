package main

import(
    "fmt"
)


func setZeroes(matrix [][]int)  {
    rows := len(matrix)
    cols := len(matrix[0])

    // Create row map, all to false
    is0col := make(map[int]bool)
    is0row := make(map[int]bool)

    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            if(matrix[row][col] == 0) {
                fmt.Println("matrix[row][col] == 0: ", row, col)
                is0row[row] = true
                is0col[col] = true
            }
        }
    }

    for i := range is0col {
        for row := 0; row < rows; row++ {
            matrix[row][i] = 0
        }
    }
    for i := range is0row {
        for row := 0; row < cols; row++ {
            matrix[i][row] = 0
        }
    }

}



func main() {
    matrix := [][]int{
        {1, 1, 1},
        {1, 0, 1},
        {1, 1, 1},
    }

    fmt.Println("First Matrix")
    // Printing the matrix
    for _, row := range matrix {
        fmt.Println(row)
    }

    setZeroes(matrix)

    fmt.Println("Second Matrix")
    for _, row := range matrix {
        fmt.Println(row)
    }
}
