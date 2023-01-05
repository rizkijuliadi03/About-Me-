package main

import "fmt"

func main() {
	var x, i, n, a, r int
	fmt.Scan(&n, &a, &r)
	fmt.Print(0)
	i = 1
	for i <= n {
		x = a * i * r
		fmt.Print(" + ", x)
		i++
	}

}
