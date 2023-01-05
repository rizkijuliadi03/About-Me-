package main

import "fmt"

func main() {
	var x1, x2, y1, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)
	fmt.Print((y1 - y2) / (x1 - x2))
}
