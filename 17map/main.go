package main

import "fmt"

func main() {
	// map mirip seperti dictionary di python
	// map mirip seperti object di javascript

	// di python, key nya bisa dicampur tipe data nya, integer atau string
	// sedangkan di golang, key nya harus tipe data yang sama

	// var buah map[string]int = map[string]int{
	// 	"apel":   10,
	// 	"pisang": 15,
	// 	"jeruk":  5,
	// }

	// var buah = map[string]int{
	// 	"apel":   10,
	// 	"pisang": 15,
	// 	"jeruk":  5,
	// }

	// buah := map[string]int{
	// 	"apel":   10,
	// 	"pisang": 15,
	// 	"jeruk":  5,
	// }
	// fmt.Println(buah)

	menu := map[string]float32{
		"kopi":       5000.00,
		"kopisusu":   10000.00,
		"nasigoreng": 20000.00,
	}
	// mengubah value berdasarkan key
	menu["kopi"] = 10000.00

	// menghapus item berdasarkan
	delete(menu, "kopisusu")
	//fmt.Println(menu)
	for nama, harga := range menu {
		fmt.Printf("Menu: %-15v Rp.%.2f \n", nama, harga)
	}

	// CONTOH KEY DENGAN ANGKA
	contohMap := map[int]string{
		1: "Satu",
		2: "Dua",
		3: "Tiga",
	}
	fmt.Println(contohMap)

	// CONTOH MAP DENGAN TIPE DATA VALUE CAMPURAN
	contohMapValueCampuran := map[string]interface{}{
		"channel":     "tentangkode",
		"isSubscribe": false,
	}
	fmt.Println(contohMapValueCampuran)

	// CONTOH MAP DENGAN TIPE DATA VALUE CAMPURAN DAN ADA ANAKAN
	contohMapValueCampuranAnakan := map[string]interface{}{
		"status":  true,
		"message": "Data berhasil disimpan",
		"data": map[string]interface{}{
			"id":         1,
			"nama":       "Tentangkode",
			"subscriber": 100,
		},
	}

	fmt.Println(contohMapValueCampuranAnakan)

	// membuat map kosong
	mapLain := make(map[string]int)
	mapLain["subscriber"] = 10

	fmt.Println(mapLain)

}
