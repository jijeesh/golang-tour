package main

import "fmt"

func main() {
	i, j := 10, 20
	p := &i
	fmt.Println(*p)
	*p = 30
	fmt.Println(i)
	p = &j
	fmt.Println(*p)
	*p = 40
	fmt.Println(j)
}
