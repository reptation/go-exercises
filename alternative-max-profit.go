package main

import (
    "fmt"
)

func maxProfit(prices []int) int {
    maxProfit := 0
    min := prices[0]
    for i:=0; i<len(prices); i++ {
        if prices[i] < min {
            min = prices[i]
        }
        if maxProfit < prices[i] - min {
            maxProfit = prices[i] - min
        }
    }
    return maxProfit
}

func main() {
    a := []int{7,6,4,3,1}
    fmt.Println(maxProfit(a))
    a = []int{7,1,5,3,6,4}
    fmt.Println(maxProfit(a))
}
