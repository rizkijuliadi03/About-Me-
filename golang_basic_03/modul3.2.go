package main

import "fmt"

func main() {
	var a, b, c string
	fmt.Scanf("%s,%s,%s", &a, &b, &c)
	if (a == b) || (a == c) || (b == c) {
		println(true)
	} else {
		println(false)
	}
}
