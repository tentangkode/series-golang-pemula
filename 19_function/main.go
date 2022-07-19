package main

import "fmt"

// function adalah sebuah kumpulan kode yang bisa digunakan berkali kali

// func hello() {
// 	fmt.Printf("Hello world")
// }

// membuat function dengan parameter
func hello(name string) {
	fmt.Println("Hello " + name)
}

// membuat function dengan parameter dan mengembalikan value (hasil)
// func penjumlahan(angka1 int, angka2 int) int {
// 	return angka1 + angka2
// }

// meringkas penulisan tipe data yang sama dari parameter disebuah function
func penjumlahan(angka1, angka2 int) int {
	return angka1 + angka2
}

// contoh function yang mengembalikan value lebih dari satu
// func tambah_kurang(angka1, angka2 int) (int, int) {
// 	return angka1 + angka2, angka1 - angka2
// }

// contoh fungsi dengan membuat nama atau label dari return
func tambah_kurang(angka1, angka2 int) (tambah int, kurang int) {
	tambah = angka1 + angka2
	kurang = angka1 - angka2
	return
}

// meringkas penulisan tipe data yang sama dari return
func tambah_kurang_kali(angka1, angka2 int) (tambah, kurang, kali int) {
	tambah = angka1 + angka2
	kurang = angka1 - angka2
	kali = angka1 * angka2
	return
}

// membuat sebuah function dengan variadic parameter
func nilaiRataRata(nilai ...int) int {

	total := 0
	for _, nil := range nilai {
		total += nil
	}

	return total / len(nilai)

}

func main() {
	hello("Jay")
	hello("World")

	penjumlahan(1, 4)

	hasil := penjumlahan(1, 4)
	fmt.Println(hasil)

	hasil1, hasil2 := tambah_kurang(10, 5)
	fmt.Printf("Hasil penambahan: %d \n", hasil1)
	fmt.Printf("Hasil pengurangan: %d \n", hasil2)

	hasilA, hasilB, hasilC := tambah_kurang_kali(10, 5)
	fmt.Printf("Hasil penambahan: %d \n", hasilA)
	fmt.Printf("Hasil pengurangan: %d \n", hasilB)
	fmt.Printf("Hasil perkalian: %d \n", hasilC)

	rata2 := nilaiRataRata(70, 80, 100, 40, 60, 74, 78, 94)
	fmt.Println(rata2)
}
