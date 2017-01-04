package main

import "fmt"

func main() {

	//variable decalartions
	var (
		xs          = []int{}
		i, j, total int
		//starttime   string
	)

	i = 1
	j = 1
	for j < 4000000 {
		fmt.Println(i)
		if i%2 == 0 { //Check if term value is even, append to slice
			xs = append(xs, i)
			total += i
		}
		i, j = i+j, i
	}
	fmt.Println("Even terms in Febonacci is -", xs)
	fmt.Println("Sum of even is - ", total)
}
