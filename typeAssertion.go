package main

import "fmt"

func toString() interface{} {
	return "ma boyyyy"
}

func main() {
	stringMaboyyy := toString()
	stringMaboyyy2 := stringMaboyyy.(string)

	fmt.Println(stringMaboyyy2)
}
