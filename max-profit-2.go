package main

import (
    "fmt"
)

func maxProfit(prices []int) int {
    totalProfit := 0
    min := prices[0]
    for i:=0; i< len(prices); i++ {
        if prices[i] < min {
            min = prices[i]
        } else {
            totalProfit += prices[i] - min
            min = prices[i]
        }
    }
    return totalProfit
}

func main() {
    a := []int{7,5,1,6,7,3,6,4}
    fmt.Println(maxProfit(a))

    a = []int{1,2,3,4,5}
    fmt.Println(maxProfit(a))

    a = []int{7,6,4,3,1}
    fmt.Println(maxProfit(a))

    a = []int{7,6,1,2,4,7,3,5,1}
    fmt.Println(maxProfit(a))
}
