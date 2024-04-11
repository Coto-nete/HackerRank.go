package main

import (
    "fmt"
    "strconv"
)

func main() {
    var arr []int
    for c := 0; c < 3; c++ {
        var input string
        fmt.Printf("Digite o %d* número:\n", c+1)
        fmt.Scanln(&input)
        num, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Erro ao converter a entrada para número inteiro.")
            return
        }
        arr = append(arr, num)
    }
    
    // Initialize max and min with the first element of the array
    max := arr[0]
    min := arr[0]
    
    // Iterate through the rest of the array to find the actual min and max
    for _, v := range arr {
        if v > max {
            max = v
        }
        if v < min {
            min = v
        }
    }
    
    fmt.Println("Menor = ", min)
    fmt.Println("Maior = ", max)
}
