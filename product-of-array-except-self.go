package main

import (
    "fmt"
)


func productExceptSelf(nums []int) []int {
    numc := make([]int, len(nums))
    for i:=0; i<len(numc); i++ {
        numc[i]=1
    }
    for i:=0; i<len(numc); i++ {
        for j:=0; j<len(nums); j++ {
            if j!=i {
                numc[i] = numc[i] * nums[j]
            }
        }
    }
    return numc
}

func main() {
    a := []int{1,2,3,4}
    fmt.Println(productExceptSelf(a))
}
