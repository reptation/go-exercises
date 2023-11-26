package main

import (
    "fmt"
)

func removeDuplicates(nums []int) int {
    k := len(nums)
    seen := 0
    var last int
    fmt.Println(last)
    for i:=len(nums)-1; i>=0; i-- {
        cur := nums[i]
        if cur != last || i==len(nums)-1 {
            seen = 0
            last = cur
        } else {
            seen++
        }
        if seen > 1 {
            k--
            copy(nums[i:len(nums)-1], nums[i+1:])
        }

    }
    //fmt.Println(nums)
    return k
}

func main() {
    a := []int{0,0,0,0,0}
    //a := []int{1,1,1,2,2,3}

    fewerDupes := removeDuplicates(a)
    fmt.Println(fewerDupes)
}
