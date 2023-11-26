package main

import (
    "fmt"
)

func majorityElement(nums []int) int {
    var counts map[int]int
    counts = make(map[int]int)
    for i:=0; i<len(nums); i++ {
        counts[nums[i]]++
    }
    fmt.Println(counts)

    k := 0
    for key,val := range(counts) {
        if val > counts[k] {
            k = key
        }
    }
    return k
}

func main() {
    a := []int{2,2,1,1,1,2,2}
    //a := []int{3,2,3}

    majority := majorityElement(a)
    fmt.Println(majority)
}
