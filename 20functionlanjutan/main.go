package main

import "fmt"

func sayHello(nama string) {
	fmt.Println("Hello", nama)
}

// contoh fungsi yang penerima paramter dalam bentuk fungsi
func test(test2 func(int) int) {
	fmt.Println(test2(5))
}

// contoh function closures yang mengembalikan hasil dalam bentuk function
func myFunc(nama string) func() {
	return func() {
		fmt.Println(nama)
	}
}

// contoh function closures yang mengembalikan hasil dalam bentuk function (closure) dan juga menerima parameter
func myFunc2(angka1, angka12 int) func(int) int {
	return func(angka3 int) int {
		return (angka1 * angka12) + angka3
	}
}

func main() {
	// referensi function sayHello ke variable x
	a := sayHello
	a("Jay")
	sayHello("Jay")

	// anonymous function (sebuah fungsi tanpa nama)
	b := func(nama string) {
		fmt.Println("Test", nama)
	}
	b("Ting")

	// anonymous function
	func(nama string) {
		fmt.Println("Testing", nama)
	}("Oke")

	// anonymous function
	test3 := func(x int) int {
		return x * 5
	}
	test(test3)

	// function closures
	myFunc("Test")()
	c := myFunc("Test")
	c()

	d := myFunc2(2, 2)
	fmt.Println(d(2))
}
