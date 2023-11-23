package main

import (
    "fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
    //fmt.Println(nums1)
    if nums1[0] == 0 {
        //fmt.Println("equals 0")
        //fmt.Println(nums1)
        //fmt.Println(nums2)
        nums1 = nums2
    } else {
        if nums2[0] < nums1[0] {
            copy(nums1[1:], nums1[:len(nums1)-1])
            nums1[0] = nums2[0]
            merge(nums1[1:], m, nums2[1:], n)
            //fmt.Println(nums1)
            //fmt.Println(nums1[1:])
        } else {
            merge(nums1[1:], m, nums2, n)
        }
    }
    //fmt.Println(nums1)
    // if n < m, insert m and bump m down
    //fmt.Println(nums1[:len(nums1)-1])
    // base case is if element of m is zero, insert remaining elements of n into it
    // if not base case, compare current element of both. if m is greater, insert n into it and shift m right
    // if n is greater do not insert, call merge with slice of m decremented by one and n
}

func main() {
    a := []int{1,3,0,0,0,0}
    //a := []int{1,0}
    b := []int{2,7,9,11}

    merge(a, 2, b, len(b))

}
