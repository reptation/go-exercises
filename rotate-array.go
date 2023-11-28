package main

import (
    "fmt"
)

func rotate(nums []int, k int) {
    var rotatedNums = make([]int, len(nums))
    copy(rotatedNums[:k], nums[len(nums)-k:])
    copy(rotatedNums[k+1:], nums[:len(nums)-k])

    nums = rotatedNums
    fmt.Println(nums)
}

func main() {
    a := []int{1,2,3,4,5,6,7}
    k := 3
    rotate(a, k)
}
