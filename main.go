package main

import "fmt"

func main() {
	a := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21}
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			fmt.Println(a[i], "genap")
		}
	}
	// fmt.Println("---------------------------------------------------------------------")

	for i := 0; i < len(a); i++ {
		if a[i]%2 == 1 {
			fmt.Println(a[i], "ganjil")
		}
	}
	// fmt.Println("---------------------------------------------------------------------")

	for bil := 1; bil < len(a); bil++ {
		i := 0
		for bag := 1; bag < len(a); bag++ {
			if bil%bag == 0 {
				i++
			}
		}
		if (i == 2) && (bil != 1) {
			fmt.Println(bil, "Prima")
		}
	}
}
