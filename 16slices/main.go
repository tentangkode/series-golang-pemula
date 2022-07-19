package main

import (
	"fmt"
)

func main() {

	// perbedaan slice dengan array,
	// pada slice kita tidak perlu menentukan panjang atau jumlah elemen yang akan ditampung didalam slice
	// dan di slice kita juga bebas untuk menambah item mau berapa saja,
	// ukuran memori akan otomatis dialokasikan kembali tergantung dari jumlah item yang ada didalam slice
	//var bahasaPemrograman []string = []string{"Go", "Python", "Java"}
	//var bahasaPemrograman = []string{"Go", "Python", "Java"}
	bahasaPemrograman := []string{"Go", "Python", "Java"}
	fmt.Println(bahasaPemrograman)

	//untuk menambahkan item kedalam slice
	bahasaPemrograman = append(bahasaPemrograman, "PHP", "C")
	fmt.Println(bahasaPemrograman)

	// [start:end]
	// mengambil item dari index 0 sampai dengan item index 1 (2 tidak termasuk)
	// 0 => inclusive
	// 2 => Exclusive
	bahasaPemrograman = append(bahasaPemrograman[0:2])
	fmt.Println(bahasaPemrograman)

	// contoh lain make slice dengan alokasi memori untuk 4 item element
	nilai := make([]int, 4)

	nilai[0] = 70
	nilai[1] = 80
	nilai[2] = 90
	nilai[3] = 69
	// jika ditambahkan index 4, maka akan error karena sudah melebihi alokasi memori
	// nilai[4] = 100

	// tetapi ketika menggunkan append(),
	// maka alokasi memori akan di alokasikan ulang menyesuikan jumlah item yang akan ditambah
	nilai = append(nilai, 100, 70, 88)
	fmt.Println(nilai)

	// cara menghapus item didalam slice berdasarkan index
	var mataKuliah = []string{"Pemrograman Terstruktur", "Pemrograman Web", "Pemrograman Berorientasi Object"}
	fmt.Println(mataKuliah)

	fmt.Println(mataKuliah[:1])
	fmt.Println(mataKuliah[2:])

	// ... => variadic parameter
	mataKuliah = append(mataKuliah[:1], mataKuliah[2:]...)

	fmt.Println(mataKuliah)
}
