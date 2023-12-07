package main

import (
    "fmt"
    "slices"
)

func hIndex(citations []int) int {
    h := 0
    test := 0
    slices.Sort(citations[:])
    for i:=0; i<len(citations); i++ {
        if citations[i] == h {
            continue
        } else {
            test = citations[i]
            for j:=test; j>h; j-- {
                if len(citations[i:]) >= j {
                    h = j
                    break
                }
            }
        }
    }
    return h
}

func main() {
    a := []int{3,0,6,1,5}
    fmt.Println(hIndex(a))

    a = []int{100}
    fmt.Println(hIndex(a))
}
