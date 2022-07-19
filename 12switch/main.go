package main

import "fmt"

func main() {

	// angka := 8

	// switch angka {
	// case 1:
	// 	fmt.Println("Angka 1")
	// case 2:
	// 	fmt.Println("Angka 2")
	// case 3:
	// 	fmt.Println("Angka 3")
	// default:
	// 	fmt.Println("Angka lain")
	// }

	// angka2 := 1

	// switch {
	// case angka2%2 == 0:
	// 	fmt.Println("Angka Genap")
	// case angka2%2 != 0:
	// 	fmt.Println("Angka Ganjil")
	// default:
	// 	fmt.Println("Tidak diketahui")
	// }

	nilai := 20
	switch {
	case nilai >= 90 && nilai <= 100:
		fmt.Println("Sangat memuaskan")
	case nilai >= 80 && nilai < 90:
		fmt.Println("memuaskan")
	case nilai >= 70 && nilai < 80:
		fmt.Println("Baik")
	case nilai >= 60 && nilai < 70:
		fmt.Println("Cukup")
	case nilai < 60:
		fmt.Println("Tidak lulus")
		fmt.Printf("Kamu butuh %d tambahan nilai untuk lulus \n", 60-nilai)
	default:
		fmt.Println("Nilai tidak diketahui")
	}

}
