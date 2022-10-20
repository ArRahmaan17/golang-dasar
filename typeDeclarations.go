package main

import "fmt"

func main() {
	type kalimat string; 
	type kesimpulan bool;

	var(
		name kalimat = "lorem ipsum";
		nikah kesimpulan = false;
	);

	fmt.Println(name)
	fmt.Println(nikah)
}