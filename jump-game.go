package main

import (
    "fmt"
)

func canJump(nums []int) bool {
    // last number can be 0
    for i:=0; i<len(nums)-1; i++ {
        if nums[i] >= len(nums[i:])-1 {
            return true
        }
        if nums[i] == 0 {
            gtg := false
            for j:=i-1; j>=0; j-- {
                if nums[j] > i-j {
                    gtg = true
                    break
                }
            }
            if gtg == false {
                return false
            }
        }
    }
    return true
}

func main() {
    a := []int{2,3,1,1,4}
    fmt.Println(canJump(a))

    a = []int{3,2,1,0,4}
    fmt.Println(canJump(a))

    a = []int{2,0,0}
    fmt.Println(canJump(a))

    a = []int{1,2,0,1}
    fmt.Println(canJump(a))

    a := []int{3,0,8,2,0,0,1}
    fmt.Println(canJump(a))
}
