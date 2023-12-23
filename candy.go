package main

import (
    "fmt"
)

func candy(ratings []int) int {
    candy := make([]int, len(ratings))
fmt.Println(candy)
    total := len(ratings)
    for i:=1; i<len(ratings); i++ {
        if ratings[i] > ratings[i-1] {
            candy[i] = candy[i-1] + 1
        }
    }

fmt.Println(candy)
    for i:=len(ratings)-2; i>=0; i-- {
        if ratings[i] > ratings[i+1] {
            candy[i] = max(candy[i+1]+1,candy[i])
        }
    }

fmt.Println(candy)
    for i:=0; i<len(ratings); i++ {
        total += candy[i]
    }

    return total
}

func main() {
c := []int{1,0,2}
// 5
fmt.Println(candy(c))

c = []int{1,0,2}
// 4
fmt.Println(candy(c))

c = []int{1,3,2,2,1}
// 7
fmt.Println(candy(c))

c = []int{1,2,87,87,87,2,1}
// 13
fmt.Println(candy(c))

c = []int{1,3,4,5,2}
// 11
fmt.Println(candy(c))

c = []int{1,6,10,8,7,3,2}
// 18
fmt.Println(candy(c))
}
