package main

import "fmt"

func main() {
	// array
	months := [...]string{
		"Januari",
		"februari",
		"Maret",
		"April",
	}
	// slice
	sliceMonth := months[2:3];

	// fmt.Println(sliceMonth);
	
	sliceMonth[0] = "Sekarang Maret";
	sliceMonth[1] = "Lalu April";
	
	sliceMonth2 := append(sliceMonth, "Mei");
	// hati hati saat merubah slice hati hati dengan capacity array
	// kalau slice mengacu ke index yang jauh dari capacity hati2 
	// fmt.Println(sliceMonth);
	// fmt.Println(len(sliceMonth));
	// fmt.Println(cap(sliceMonth));
	fmt.Println(months);
	fmt.Println(sliceMonth2);
}