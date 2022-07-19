package main

import (
	"errors"
	"fmt"
)

func penjumlahan(angka1, angka2 int) (int, error) {
	if angka1 > 10 || angka2 > 10 {
		return 0, errors.New("Angka melebihi limit")
	}
	return angka1 + angka2, nil
}

func main() {

	hasil, _ := penjumlahan(15, 10)

	fmt.Println(hasil)
}
