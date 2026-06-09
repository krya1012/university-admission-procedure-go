package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	mean := float64(a+b+c) / 3.0
	fmt.Println(mean)
	fmt.Println("Congratulations, you are accepted!")
}
