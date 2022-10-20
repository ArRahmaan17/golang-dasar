package main

import "fmt"

func main() {
	var angka = [3]int8{
		66,
		67,
		21,
	}
	angka[2] = 36
	fmt.Println(angka)
	fmt.Println(angka[0])
	fmt.Println(angka[1])
}
