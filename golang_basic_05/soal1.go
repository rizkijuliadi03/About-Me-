package main

import "fmt"

func main() {
	var marcus, degea, i, N, tendang int
	fmt.Scan(&N)
	marcus = 0
	degea = 0
	for i = 0; i < N; i++ {
		fmt.Scan(&tendang)
		if tendang%2 == 0 {
			fmt.Println("Tendangan terlalu kesamping")
			degea = degea + 1
		} else if tendang%5 == 0 {
			fmt.Println("Tendangan terlalu keatas")
			degea = degea + 1
		} else {
			fmt.Println("Tendangan tepat sasaran")
			marcus = marcus + 1
		}
	}
	println(degea, marcus)
}
