package main

import (
    "fmt"
    "math/rand"
)

type RandomizedSet struct {
    vals map[int]bool
}


func Constructor() RandomizedSet {
    var r RandomizedSet
    r.vals = make(map[int]bool)
    return r
}

func (this *RandomizedSet) Insert(val int) bool {
    _, exists := this.vals[val]
    if exists {
        return false
    } else {
        this.vals[val] = true
        return true
    }
}

func (this *RandomizedSet) Remove(val int) bool {
    _, exists := this.vals[val]
    if exists {
        delete(this.vals, val)
        return true
    } else {
        return false
    }
}

func (this *RandomizedSet) GetRandom() int {
    rand_idx := rand.Intn(len(this.vals))
    i := 0
    for k, _ := range(this.vals) {
        if i == rand_idx {
            return k
        }
        i++
    }
}



func main() {
    rs := Constructor()
    rs.Insert(0)
    rs.Insert(1)

    rs.Remove(0)
    rs.Insert(2)
    rs.Insert(1)
    fmt.Println(rs.GetRandom())
    fmt.Println(rs)
}
