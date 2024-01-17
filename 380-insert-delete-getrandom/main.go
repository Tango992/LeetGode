package main

import (
	"math/rand"
	"slices"
)

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
*/

func main() {
	obj := Constructor()
	obj.Insert(1)
	obj.Remove(2)
	obj.Insert(2)
}

type RandomizedSet struct {
    Val []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{Val: []int{}}
}


func (rs *RandomizedSet) Insert(val int) bool {
    if slices.Contains(rs.Val, val) {
        return false
    }

    rs.Val = append(rs.Val, val)
    return true
}


func (rs *RandomizedSet) Remove(val int) bool {
    if !(slices.Contains(rs.Val, val)) {
        return false
    }

    index := slices.Index(rs.Val, val)
    rs.Val = slices.Delete(rs.Val, index, index+1)
    return true
}


func (rs *RandomizedSet) GetRandom() int {
    index := rand.Intn(len(rs.Val))
    return rs.Val[index]
}