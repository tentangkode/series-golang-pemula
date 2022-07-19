package main

import (
	"fmt"
	"strconv"
)

func main() {

	//menampilkan binary dari sebuah angka (base 2)
	//menampilkan octal dari sebuah angka (base 4)
	//menampilkan angka decimal (base 10)
	//menampilkan hexadecimal (base 16)

	var stringAngka string = "10"

	angkaInt, _ := strconv.ParseInt(stringAngka, 10, 64)
	fmt.Printf("Data tipe: %T value: %v \n", angkaInt, angkaInt)

	var stringAngka2 string = "10.12"
	angkaFloat, _ := strconv.ParseFloat(stringAngka2, 64)
	fmt.Printf("Data tipe: %T value: %v \n", angkaFloat, angkaFloat)

	var stringAngka3 string = "0"
	bl, _ := strconv.ParseBool(stringAngka3)
	fmt.Printf("Data tipe: %T value: %v \n", bl, bl)
}
