package main

import (
    "fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
    nums3 := make([]int, m+n)
    j := 0

    for i, _ := range nums1 {
        if len(nums2) == 0 {
            nums3[i] = nums1[j]
            j++
        } else if nums1[j] == 0 {
            nums3[i] = nums2[0]
            nums2 = nums2[1:]
            j++
        } else if nums1[j] < nums2[0] {
            nums3[i] = nums1[j]
            j++
        } else if nums1[j] >= nums2[0] {
            nums3[i] = nums2[0]
            nums2 = nums2[1:]
        }
    }
nums1 = nums3
fmt.Println(nums1)

}
func main() {
    a := []int{1,2,3,0,0,0}
    b := []int{2,5,6}

    merge(a, 3, b, 3)

}
