package main

import (
	"fmt"
	"math"
	"math/rand"
)

func h(x, i int) int{
	return x % (i + 1) + x % (m - i)
}

var m = 100000
var k = int(math.Log(float64(m)))
var a = make([]bool, m)

func insert(x int){
	for i:= 0; i < k; i++{
		a[h(x, i)] = true
	}
}

func find(x int) bool{
	for i:= 0; i < k; i++{
		if !a[h(x, i)]{
			return false
		}
	}
	return true
}

func main() {
	for i:=0; i < 20000; i++{
		insert(rand.Int())
	}
	for i:=0; i < 10; i++{
		fmt.Println(find(i))
	}
}