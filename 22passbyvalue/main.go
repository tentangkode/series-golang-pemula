package main

import "fmt"

func updateAngka(y int) {
	y = 100
}

func updateBuah(b map[string]int) {
	b["apel"] = 100
}

func main() {
	// Go => adalah bahasa pemrograman yang sifat nya Pass by Value
	// Pass by Value ini artinya misal kita memasukkan sebuah variable kedalam sebuah parameter function/variable,
	// maka go akan membuat sebuah copy-an dari variable yang dimasukkan

	// Jadi ada dua grup tipe data di Golang:

	// Mutable dan Immutable data type
	// => Mutable => nilai dari sebuah variable bisa diubah (Slices, Maps, Functions)
	// => immutable => nilai dari sebuah variable sifat nya tetap (Strings, Ints, Floats, Booleans, Arrays dan Structs)

	// dan ada yang menyebutkan
	// => Non Pointer Values => Strings, Ints, Floats, Booleans, Arrays dan Structs
	// => Pointer Wrapper Values => Slices, Maps, Functions

	// => Non Pointer Values => Strings, Ints, Floats, Booleans, Arrays dan Structs
	// x := 10
	// y := x
	// y = 100
	// fmt.Println(x, y)

	// updateAngka(x)
	// fmt.Println(x)

	// => Pointer Wrapper Values => Slices, Maps, Functions
	buah := map[string]int{
		"apel":   10,
		"jeruk":  15,
		"mangga": 5,
	}

	// x := buah
	// x["apel"] = 100

	// fmt.Println(buah)
	// fmt.Println(x)

	updateBuah(buah)
	fmt.Println(buah)
}
