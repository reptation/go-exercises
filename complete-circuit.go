package main

import (
    "fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
    currentGas := 0
    start := 0
    done := false
    for i:=0; i<len(gas); i++ {
        if cost[i] > gas[i] {
            continue
        }
        isBreak := false
        for j:=i; j<len(gas); j++ {
            currentGas += gas[j] - cost[j]
            if currentGas < 0 {
                isBreak = true
                currentGas = 0
                break
            }
        }
        if isBreak {
            continue
        }
        if i == 0 {
            return i
        }
        // we made to the end of the array. check the first entries
        for k:=0; k<i; k++ {
            currentGas += gas[k] - cost[k]
            if currentGas < 0 {
                return -1
            }
            done = true
        }
        if done {
            start = i
            break
        }
    }
    if done {
        return start
    } else {
        return -1
    }
}

func main() {
   g := []int{1,2,3,4,5}
   c := []int{3,4,5,1,2}
   // 3
   fmt.Println(canCompleteCircuit(g,c))
}
