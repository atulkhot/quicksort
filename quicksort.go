package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(a []int) int {
	hi := len(a) - 1
	pivot := a[hi]
	i := -1

	for j := 0; j < hi; j++ {
		if a[j] <= pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	i++
	a[i], a[hi] = a[hi], a[i]
	return i
}

func Make_random_array(num_items, max int) []int {
	var rand_array []int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num_items; i++ {
		rand_array = append(rand_array, rand.Intn(max))
	}

	return rand_array
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Print_array(arr []int, num_items int) {
	num_elems := Min(len(arr), num_items)
	fmt.Println(arr[:num_elems])
}

func quicksort(a []int) {
	if len(a) <= 1 {
		return
	}
	p := partition(a)
	quicksort(a[:p])
	quicksort(a[p:])
}

func main() {
	arr := Make_random_array(20, 25)
	Print_array(arr, len(arr)+1)
	quicksort(arr)
	Print_array(arr, len(arr)+1)
}
