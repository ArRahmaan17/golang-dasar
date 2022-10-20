package main

import "fmt"

// kontrak interface
type goodBye interface {
	GoodBye() string
}

// struct tipe data seperti object yang berada di pemrogramman oop
// type Customer struct {
// 	Name, Address string
// 	Age           int
// 	Married       bool
// }

type tempatKerja struct {
	Name string
}

// func (ardhi Customer) SayGoodBye() {
// 	fmt.Println("Hiii I am", ardhi.Name, "Good Bye Cikajogja")
// }

func (tempatkerja tempatKerja) GoodBye() string {
	return tempatkerja.Name
}

func SayGoodBye(goodbye goodBye) {
	fmt.Println("good bye", goodbye.GoodBye())
}

// func interfaceKosong(i int) interface{} {
// 	if i == 1 {
// 		return i
// 	} else if i == 2 {
// 		return false
// 	} else {
// 		return "Telo Gaming"
// 	}
// }

func main() {
	// ardhi := Customer{
	// 	Name:    "Maman Recing",
	// 	Address: "Jambangan",
	// 	Age:     19,
	// 	Married: false,
	// }

	// fmt.Println(ardhi)
	// ardhi.SayGoodBye()
	tempatKerjaToxic := tempatKerja{
		Name: "CikaJogja.Com",
	}
	SayGoodBye(tempatKerjaToxic)

	// interfaceKosong := interfaceKosong(3)

	// fmt.Println(interfaceKosong)
}
