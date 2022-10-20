package main

import "fmt"

type data interface {
	getDataSlogan() string
}

type Kelurahan struct {
	Nama, Slogan                        string
	IdKecamtan, IdKabupaten, IdProvinsi int
}

func (kelurahan Kelurahan) getDataSlogan() string {
	return kelurahan.Slogan
}

func fetchDataKelurahan(data data) {
	fmt.Println("Slogan Kabupaten Ini adalah", data.getDataSlogan())
}

func main() {
	karanganyar := Kelurahan{
		Nama:   "karangayar",
		Slogan: "Karanganyar Tentrem",
	}
	fetchDataKelurahan(karanganyar)
}
