package main

import "fmt"

func main() {
	biodata := map[string]string{
		"name" : "Ardhi Rahmaan",
		"address" : "Sleman",
		"age" : "19",
	}

	biodata["status"] = "Programmer";

	fmt.Println(biodata)
	fmt.Println(biodata["name"])
	fmt.Println(biodata["status"])
}