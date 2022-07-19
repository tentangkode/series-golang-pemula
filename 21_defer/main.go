package main

import "fmt"

// defer => sebuah statement untuk menjalankan suatu baris kode tertentu diakhir operasi
// defer => last in first out (LIFO)

func main() {

	kumpulanDefer()

	defer fmt.Println("Tiga")
	defer fmt.Println("Dua")
	defer fmt.Println("Satu")

	fmt.Println("Hello world")
	defer fmt.Println("Nol")
}

func kumpulanDefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println("Angka:", i)
	}
}
