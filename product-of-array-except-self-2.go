package main

import (
    "fmt"
)


func productExceptSelf(nums []int) []int {
    numc := make([]int, len(nums))


    for i:=0; i<len(numc); i++ {
        numc[i]=1
    }
    backward := make([]int, len(nums))
    copy(backward, numc)
    forward :=  make([]int, len(nums))
    copy(forward, numc)

//yang's forward and backward implementation
forward[0] = nums[0]
for i:=1; i<len(forward)-1; i++{
	forward[i] = forward[i-1]*nums[i]
}

backward[len(backward)-1] = nums[len(numc)-1]
for j:=len(backward)-2; j>0; j--{
	backward[j]= backward[j+1]*nums[j]
}

numc[0] = backward[1]
numc[len(numc)-1]=forward[len(forward)-2]
for k:=1; k<len(numc)-1;k++{
	numc[k]=forward[k-1]*backward[k+1]
}
//fmt.Println(forward)
//fmt.Println(backward)
return numc
}

func main() {
    a := []int{1,2,3,4}
    // {24, 12, 8, 6}
    fmt.Println(productExceptSelf(a))
}



