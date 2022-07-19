package main

import "fmt"

func main() {

	// mendapatkan tipe data
	fmt.Printf("Tipe data: %T \n", 10)
	// mendapatkan value data
	fmt.Printf("Tipe data: %v \n", 10)
	//print tanda persen
	fmt.Printf("Persen: %v%% \n", 10)
	//memasukkan hasil printf ke variable dengan menggunakan sprintf
	var myString string = fmt.Sprintf("Tipe data: %T", 10)
	fmt.Println(myString)

	//mengambil boolean
	fmt.Printf("True atau False: %t \n", 1 > 2)

	//menampilkan binary dari sebuah angka (base 2)
	fmt.Printf("Binari: %b \n", 10)
	//menampilkan octal dari sebuah angka (base 4)
	fmt.Printf("Octal: %o \n", 10)
	//menampilkan angka decimal (base 10)
	fmt.Printf("Desimal: %d \n", 10)
	//menampilkan hexadecimal (base 16)
	fmt.Printf("Hexadesimal: %x \n", 1000)

	// notasi ilmiah
	fmt.Printf("Hasil: %e \n", 2.812748916249)
	// menampilkan angka decimal tetapi dipotong dan dibulatkan
	fmt.Printf("Hasil: %F \n", 2.812748916249)
	// menampilkan angka decimal tanpa dipotong dan dibulatkan
	fmt.Printf("Hasil: %g \n", 2.812748916249)

	// menampilkan string default
	fmt.Printf("Bahasa: %s \n", "GO")

	// menampilkan string dalam tanda kutip ganda
	fmt.Printf("Bahasa: %q \n", "GO")

	/* width dan percision */

	// panjang nya default dan presisi juga default
	fmt.Printf("%f \n", 100.000)

	// panjang ke kiri
	fmt.Printf("Bahasa: %12s keren\n", "Go")

	// panjang ke kanan kanan
	fmt.Printf("Bahasa: %-12q keren\n", "Go")

	// panjangnya default dan presisi 2
	fmt.Printf("%.2f \n", 10.000)

	// panjangnya 9 dan presisi 2 (kiri)
	fmt.Printf("Angka: %9.2f keren\n", 10.000)

	// panjangnya 9 dan presisi 2 (kanan)
	fmt.Printf("Angka: %-9.2f keren\n", 10.000)

	// panjangnya 9 dan presisi 0
	fmt.Printf("Angka: %9.f keren coy\n", 10.000)

	// padding
	var noUrut string = fmt.Sprintf("No.Urut A%04d di Loket A\n", 12)
	fmt.Println(noUrut)
}
