package main

import "fmt"

func main() {
	sum := 0
	for x := 0; x < 10; x++ {
		sum += x
	}
	fmt.Println(sum)

	wsum := 1
	for wsum < 1000 {
		wsum += wsum
	}
	fmt.Println(wsum)
}
