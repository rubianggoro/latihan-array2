package main

import "fmt"

func main() {
	a := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var genap, ganjil, prima []int
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			genap = append(genap, a[i])
		}
	}
	fmt.Printf("Bilangan genap : %+v, jumlah = %d\n", genap, len(genap))
	fmt.Println("---------------------------------------------------------------------")

	for i := 0; i < len(a); i++ {
		if a[i]%2 == 1 {
			ganjil = append(ganjil, a[i])
		}
	}
	fmt.Printf("Bilangan ganjil : %+v, jumlah = %d\n", ganjil, len(ganjil))
	fmt.Println("---------------------------------------------------------------------")

	for bil := 1; bil < len(a); bil++ {
		i := 0
		for bag := 1; bag < len(a); bag++ {
			if bil%bag == 0 {
				i++
			}
		}
		if (i == 2) && (bil != 1) {
			prima = append(prima, bil)
		}
	}
	fmt.Printf("Bilangan prima : %+v, jumlah = %d\n", prima, len(prima))
	fmt.Println("---------------------------------------------------------------------")
}
