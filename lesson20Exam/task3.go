package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	maks := slice[0]
	for _, v := range slice {
		if maks < v {
			maks = v
		}
	}
	fmt.Println(maks)
}