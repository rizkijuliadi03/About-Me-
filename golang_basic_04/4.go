package main

import "fmt"

func main() {
	var n, i int
	var k1, k2, k3, k4 bool
	var status bool
	fmt.Scan(&n)
	status = true
	i = 0
	for i < n && status {
		fmt.Scan(&k1, &k2, &k3, &k4)
		status = k1 && k2 && k3 && k4
		i += 1
	}
	fmt.Println(status)
}
