package main

import (
    "fmt"
)

func rotate(nums []int, k int) {
    // If k is greater or equal than len(nums) there are problems
    if k >= len(nums) {
        k = k % len(nums)
    }
    var dupeNums = make([]int, len(nums))
    copy(dupeNums, nums)
    copy(nums[:k], dupeNums[len(dupeNums)-k:])
    copy(nums[k:], dupeNums[:len(dupeNums)-k])
    fmt.Println(nums)
}

func main() {
    a := []int{1,2,3,4,5,6,7}
    k := 3
    rotate(a, k)

    a = []int{1}
    k = 1
    rotate(a,k)

    a = []int{1,2}
    k = 2
    rotate(a, k)

    a = []int{-1}
    k = 2
    rotate(a,k)

    a = []int{1,2,3,4,5,6}
    k = 1
    rotate(a,k)
}
