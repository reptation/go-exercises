package main

import (
   "fmt"
)

func maxProfit(prices []int) int {
    maxProfit := 0
    for i:=0; i<len(prices); i++ {
        for j:=i+1; j<len(prices); j++ {
            if prices[j] <= prices[i] {
                continue
            }
            profit := prices[j] - prices[i]
            if profit > maxProfit {
                maxProfit = profit
            }
        }
    }
    return maxProfit
}

func main() {
    a := []int{7,1,5,3,6,4}
    fmt.Println(maxProfit(a))
}
