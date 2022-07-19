package main

import "fmt"

func ubahValue(angka *int) {
	*angka = 100
}

func ubahValue2(angka int) {
	angka = 100
}

func main() {
	// & => ampersand (mengambil pointer)
	// * => asterisk (dereference => mengambil value)

	// x := 7
	// fmt.Println(&x)
	// y := &x
	// *y = 10
	// fmt.Println(x, *y)

	nilai := 10
	fmt.Println(nilai)
	ubahValue2(nilai)
	fmt.Println(nilai)

	var nilai2 *int
	nilai2 = &nilai
	*nilai2 = 100
	fmt.Println(nilai, *nilai2)
}
