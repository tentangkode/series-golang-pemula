package main

import "fmt"

func main() {
	// operan itu adalah object atau sesuatu yang akan dibandingkan
	// Operator Relasional / Kesetaraan adalah operator yang digunakan
	// untuk menentukan relasi atau hubungan dari dua buah operand

	/*
		== 	It checks if the values of two operands are equal or not;
			if yes, the condition becomes true.
			(A == B) is not true.
		!=	It checks if the values of two operands are equal or not;
			if the values are not equal, then the condition becomes true.
			(A != B) is true.
		>	It checks if the value of left operand is greater than the value of right operand;
			if yes, the condition becomes true.	(A > B) is not true.
		<	It checks if the value of left operand is less than the value of the right operand;
			if yes, the condition becomes true.	(A < B) is true.
		>=	It checks if the value of the left operand is greater than or equal to the value of the right operand;
			if yes, the condition becomes true.	(A >= B) is not true.
		<=	It checks if the value of left operand is less than or equal to the value of right operand;
			if yes, the condition becomes true.	(A <= B) is true.
	*/

	// contoh
	fmt.Printf("True atau False: %t \n", 1 != 1)
}
