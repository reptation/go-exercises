package main

import (
    "fmt"
)

func jump(nums []int) int {
    minJumps := 0
    for i:=0; i<len(nums)-1; {
        if i == len(nums)-1 {
            break
        }

        if i + nums[i] >= len(nums)-1 {
            minJumps++
            break
        }
        localMax := 0
        next := i+1
        for j:=i+1; j<=nums[i]+i; j++ {
            if nums[j] + nums[i] - (nums[i]-j) >= localMax {
                localMax = nums[j] + nums[i] - (nums[i]-j)
                next = j
            }
        }
        //case where just jumping using current value makes sense
        if nums[i] >= next + nums[next] {
            next = i + nums[i]
        }
        i = next
        minJumps++
    }
    return minJumps
}

func main() {
    a := []int{2,3,1,1,4}
    fmt.Println(jump(a))

    a = []int{2,3,0,1,4}
    fmt.Println(jump(a))

    a = []int{1,2,1,1,1}
    fmt.Println(jump(a))

    a = []int{1,2,3}
    fmt.Println(jump(a))

    a = []int{10,9,8,7,6,5,4,3,2,1,1,0}
    fmt.Println(jump(a))

    a = []int{7,0,9,6,9,6,1,7,9,0,1,2,9,0,3}
    // should be 2
    fmt.Println(jump(a))
}
