package main

import "fmt"

func main() {

	// operator aritmatika adalah yang digunakan untuk melakukan operasi matematika

	/*
		+ (penambahan)
		- (pengurangan)
		* (perkalian)
		/ (pembagian)
		% (modulus) => sisa bagi
		++ increment operator, untuk menambahkan value dari sebuah integer (ditambah 1)
		-- decrement operator, untuk mengurangi value dari sebuah integer (dikurang 1)
	*/

	var angka1 int = 12
	var angka2 int = 3

	hasil := angka1 - angka2

	fmt.Println(hasil)

	// sisi kanan dan kiri harus sama tipe data nya
	// var angka3 float32 = 12
	// var angka4 int = 3

	// hasil2 := angka3 - angka4

	// fmt.Println(hasil2)

	nomor := 10
	nomor++
	fmt.Println(nomor)
}
