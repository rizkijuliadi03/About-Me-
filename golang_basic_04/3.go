package main

import (
	"fmt"
)

func main() {
	var n, i, bilangan, jumlah int
	fmt.Scan(&n)
	jumlah = 0
	i = 1
	for i <= n {
		fmt.Scan(&bilangan)
		for bilangan < 0 || bilangan > 9 {
			fmt.Scan(&bilangan)
		}
		jumlah += bilangan
		i += 1
	}
	fmt.Println(jumlah)
}
