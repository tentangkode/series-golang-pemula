package main

import "fmt"

func main() {
	// Operator Logika adalah operator yang digunakan
	// untuk membandingkan 2 kondisi logika,
	// yaitu logika benar (TRUE) dan logika salah (FALSE).

	/*
		&&	Called Logical AND operator. If both the operands are non-zero, then condition becomes true.	(A && B) is false.
		||	Called Logical OR Operator. If any of the two operands is non-zero, then condition becomes true.	(A || B) is true.
		!	Called Logical NOT Operator. Use to reverses the logical state of its operand. If a condition is true then Logical NOT operator will make false.
	*/

	fmt.Printf("True atau False: %t \n", 1 == 2 || 2 == 2)

	fmt.Printf("True atau False: %t \n", !false)
}
