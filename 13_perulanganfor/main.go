package main

import "fmt"

func main() {

	// mendeklarasikan inisial diluar for dan increment didalam for
	// i := 0
	// for i <= 10 {
	// 	fmt.Printf("Tentangkode ke %v \n", i)
	// 	i++
	// }

	// mendeklarasikan inisial, kondisi dan increment di statement for
	// for i := 0; i <= 10; i += 2 {
	// 	fmt.Printf("Tentangkode ke %v \n", i)
	// }

	// for range
	// membuat slice
	var namaHari = []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	// for i := range namaHari {
	// 	fmt.Println(namaHari[i])
	// }

	// mengambil index dan nama hari
	for index, hari := range namaHari {
		fmt.Printf("Index ke %v nama hari: %v \n", index, hari)
	}

	// mengabaikan index, dan mengambil nama hari
	for _, hari := range namaHari {
		fmt.Printf("nama hari: %v \n", hari)
	}
}
