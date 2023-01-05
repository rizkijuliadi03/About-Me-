package main

import "fmt"

func main() {
	var x1, x2, x3, x4 int
	fmt.Scan(&x1)
	fmt.Scan(&x2)
	fmt.Scan(&x3)
	fmt.Scan(&x4)
	if x1%4 == 0 || x1%400 == 0 && x1%100 != 0 {
		print(true)
	} else {
		print(false)
	}
	if x2%4 == 0 || x2%400 == 0 && x2%100 != 0 {
		print(true)
	} else {
		print(false)
	}
	if x3%4 == 0 || x3%400 == 0 && x3%100 != 0 {
		print(true)
	} else {
		print(false)
	}
	if x4%4 == 0 || x4%400 == 0 && x4%100 != 0 {
		print(true)
	} else {
		print(false)
	}

}
