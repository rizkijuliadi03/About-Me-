package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	kg := x / 1000
	pon := x / 453.592
	ons := x / 28.3495
	fmt.Print(kg, pon, ons)
}
