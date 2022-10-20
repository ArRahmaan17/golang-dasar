package main

import "fmt"

func main()  {
	perhitunganNilai(90);
}
func perhitunganNilai(nilai int8)  {
	switch {
	case nilai < 70:
		fmt.Println("Perlu ada usaha lagi yaa")
	case nilai < 50:
		fmt.Println("Perlu ada peningkatan nilai nya")
	default :
		fmt.Println("Di pertahankan terus ya nilainya")
	}
}