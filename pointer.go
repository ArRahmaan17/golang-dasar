package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	rahmaanAdress1 := Address{
		"Sleman", "D.I.Yogyakarta", "Indonesia",
	}

	// change all reference to new value
	rahmaanAdress2 := &rahmaanAdress1
	fmt.Println("address 1", rahmaanAdress1)
	fmt.Println("changing address 1")
	*rahmaanAdress2 = Address{
		"Surakarta", "Jawa Tengah", "Indonesia",
	}
	rahmaanAdress2.City = "Surakarta"
	rahmaanAdress2.Province = "Jawa Tengah"
	fmt.Println("address 1", rahmaanAdress2)
	fmt.Println("address 2", rahmaanAdress1)

	// pass by reference
	// fmt.Println(rahmaanAdress1)
	// rahmaanAdress2.City = "Surakarta"
	// rahmaanAdress2.Province = "Jawa Tengah"
	// fmt.Println(rahmaanAdress2)
	// fmt.Println(rahmaanAdress1)

	// pass by value
	// rahmaanAdress2 := rahmaanAdress1
	// fmt.Println(rahmaanAdress1)
	// rahmaanAdress2.City = "Surakarta"
	// rahmaanAdress2.Province = "Jawa Tengah"
	// fmt.Println(rahmaanAdress2)
}
