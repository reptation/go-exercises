package main

import (
    "fmt"
)

func candy(ratings []int) int {
    trend := "equal"
    newTrend := "equal"
    total := len(ratings)
    trendNum := 0
    for i:=0; i<len(ratings)-1; i++ {
//fmt.Println("i is ", i)
        if ratings[i+1] == ratings[i] {
            newTrend = "equal"
        } else if ratings[i+1] < ratings[i] {
            newTrend = "decreasing"
        } else if ratings[i+1] > ratings[i] {
            newTrend = "increasing"
        }

        if newTrend == trend && trend != "equal" {
            trendNum++
        } else if newTrend == "equal" {
            trendNum = 0
        }  else {
            trendNum = 1
        }
fmt.Println("i is ", i)
fmt.Println("val is ", ratings[i])
fmt.Println("next val is", ratings[i+1])
fmt.Println("trendNum is ", trendNum)
fmt.Println("trend and newTrend is ", trend, newTrend)
fmt.Println("total is ", total)
fmt.Println("")
        total += trendNum
        trend = newTrend
   }
    return total
}

func main() {
    c := []int{1,0,2}
//    // 5
//    fmt.Println(candy(c))
//
//    c = []int{1,2,2}
//    // 4
//    fmt.Println(candy(c))

//c = []int{1,3,2,2,1}
//// 7
//fmt.Println(candy(c))
//
//c = []int{1,2,87,87,87,2,1}
//// 13
//fmt.Println(candy(c))

//c = []int{1,3,4,5,2}
//// 11? - 1,2,3,4,1
//fmt.Println(candy(c))

c = []int{1,6,10,8,7,3,2}
// 18
fmt.Println(candy(c))
}
