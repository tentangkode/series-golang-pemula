package main

import "fmt"

func main() {

	// break => untuk menghentikan perulangan ditengah perulangan
	// i := 1
	// for i <= 10 {
	// 	// break
	// 	if i == 5 {
	// 		fmt.Println("Perulangan sudah mencapai ke angka 5, proses distop")
	// 		break
	// 	}
	// 	i++

	// }
	// fmt.Println("Proses akan langsung loncat ke baris setelah statement for loop")

	// continue => untuk melanjutkan perulangan setelah kondisi tertentu
	// i := 1
	// for i <= 10 {

	// 	// break
	// 	if i == 5 {
	// 		fmt.Printf("Perulangan sudah sampai ke %v, proses perulangan akan loncat langsung ke perulangan ke 7 \n", i)
	// 		i += 2
	// 		continue
	// 	}

	// 	fmt.Printf("Perulangan ke %v \n", i)
	// 	i++
	// }

	// goto => untuk keluar dari perulangan
	i := 0
	for i <= 10 {

		// break
		if i == 5 {
			goto tampilkan_ini
		}

		fmt.Printf("Perulangan ke %v \n", i)
		i++
	}

tampilkan_ini:
	fmt.Println("Jangan lupa subscribe guys")
}
