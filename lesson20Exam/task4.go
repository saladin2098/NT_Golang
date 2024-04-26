package main

import "fmt"

func main(){
	slice := []int{1, 2, 3, 4,5,7,8}
	for i := 0; i < len(slice)-1; i++ {
		if slice[i+1] - slice[i] != 1 {
			fmt.Println(slice[i]+1)
		}
	}
	
}