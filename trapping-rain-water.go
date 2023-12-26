package main

import (
    "fmt"
    "slices"
)

func trap(height []int) int {
     water := 0
     for i:=0; i<slices.Max(height); i++ {
         layer := make([]int, len(height))

         hasWall := false
         waterThisLayer := 0
         for j:=0; j<len(height); j++ {
             res := height[j]-i
             if res > 0 {
                 hasWall = true
                 layer[j] = 1
             } else if hasWall == false && res <= 0 {
                 layer[j] = 2
             }
         }

         hasWall = false
         for j:=len(height)-1; j>=0; j-- {
             res := height[j]-i
             if res > 0 {
                 hasWall = true
                 layer[j] = 1
             } else if hasWall == false && res <= 0 {
                 layer[j] = 2
             }
         }

         for j:=0; j<len(height); j++ {
             if layer[j] == 0 {
                 waterThisLayer++
             }
         }
         //fmt.Println(layer)
         if waterThisLayer == 0 {
             break
         }
         water += waterThisLayer
     }
     return water
}



func main() {
h := []int{0,1,0,2,1,0,1,3,2,1,2,1}
fmt.Println(trap(h))

//h = []int{4,2,0,3,2,5}
//fmt.Println(trap(h))

}
