package main

import "fmt"

func main() {

	nilai := 50

	// fmt.Println("Di print sebelum if")
	// // lebih besar sama dengan 90
	// if nilai >= 90 {
	// 	fmt.Println("Sangat Memuaskan")
	// } else {
	// 	fmt.Println("Memuaskan")
	// }

	// fmt.Println("Di print setelah if")

	// lebih besar sama dengan 90
	if nilai >= 90 {
		fmt.Println("Sangat Memuaskan")
	} else if nilai >= 80 && nilai < 90 {
		fmt.Println("Memuaskan")
	} else if nilai >= 70 && nilai < 80 {
		fmt.Println("Baik")
	} else if nilai >= 60 && nilai < 70 {
		fmt.Println("Cukup")
	} else if nilai >= 0 && nilai < 60 {
		fmt.Println("Tidak lulus")
		fmt.Printf("Kamu butuh %d tambahan nilai untuk lulus \n", 60-nilai)
	} else {
		fmt.Println("Nilai tidak diketahui")
	}

}
