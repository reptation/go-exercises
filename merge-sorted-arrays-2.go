// Another take on merging 2 sorted arrays. This time do not use a separate 3rd array

package main

import (
    "fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
    for i:=0; i<n; i++ {
        nums1index := 0
        for nums2[i] >= nums1[nums1index] && nums1index < m  {
            nums1index++
        }

        copy(nums1[nums1index+1:], nums1[nums1index:len(nums1)-1])
        nums1[nums1index] = nums2[i]
        m++
    }
fmt.Println(nums1)
}
func main() {
    //a := []int{1,0}
    //b := []int{2}
    //m := 1
    //n := 1

    //a := []int{-1,3,0,0,0,0,0}
    //b := []int{0,0,1,2,3}
    //m := 2
    //n := 5

    //a := []int{-1,0,0,3,3,3,0,0,0}
    //b := []int{1,2,2}
    //m := 6
    //n := 3

    //a := []int{1,2,3,0,0,0}
    //b := []int{2,5,6}
    //m := 3
    //n := 3

    a := []int{0}
    b := []int{1}
    m := 0
    n := 1

    //a := []int{1}
    //b := []int{}
    //m := 1
    //n := 0

    merge(a, m, b, n)

}
