package main

import "fmt"

func main() {
	name := "nabil"

	if name == "nabil" {
		fmt.Println("Halo Nabil")
	}else {
		fmt.Println("Kenalan lah")
	}

	val := "hungariana"
	switch val {
	case "hungarian":
		fmt.Println("mantab")
		default:
			fmt.Println("beda ini")
	}

	var length = len(name)
	switch {
	case length > 5 :
		fmt.Println("Panjang")
	case length < 3 :
		fmt.Println("Pendek banget")
	default :
		fmt.Println("Cukup")
	}

}