package main

import "fmt"

func main(){
	oper := []string{"--X","X++","X++","++X"}
    Calculate(&oper)
}
func Calculate(slice *[]string) {
	sum := 0
	for _, op := range *slice {
		if op == "--X" || op == "X--" {
			sum -= 1
		} else if op == "++X" || op == "X++" {
			sum += 1
		}
	}
	fmt.Println(sum)

}
