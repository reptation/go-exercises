package main

import (
    "fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
    currentGas := 0
    start := 0
    totalGas := 0

    for i:=0; i<len(gas); i++ {
        currentGas += gas[i] - cost[i]
        totalGas += gas[i] - cost[i]

        if currentGas < 0 {
            currentGas = 0
            start = i+1
        }
    }

    if totalGas < 0 {
        return -1
    }
    return start

}

func main() {
   g := []int{1,2,3,4,5}
   c := []int{3,4,5,1,2}
   // 3
   fmt.Println(canCompleteCircuit(g,c))
}
