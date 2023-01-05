package main

import "fmt"

func main() {
	var b, faktor, nfaktor int
	fmt.Scan(&b)
	fmt.Print("Faktor: ")
	nfaktor = 0
	for faktor = 1; faktor <= b; faktor++ {
		if b%faktor == 0 {
			fmt.Print(faktor, " ")
			nfaktor += 1
		}
	}
	fmt.Println("\nPrima:", nfaktor == 2)
}
