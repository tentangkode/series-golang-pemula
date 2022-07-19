package main

import "fmt"

func main() {

	// membuat array
	var arraySatu [5]int
	// default [0 0 0 0 0]
	arraySatu[0] = 100
	arraySatu[3] = 300

	fmt.Println(arraySatu)

	var arrayDua [5]string
	// default [     ]
	arrayDua[0] = "Golang"
	arrayDua[1] = "JavaScript"
	fmt.Println(arrayDua)

	var arrayTiga [5]float32
	// default [0 0 0 0 0]
	arrayTiga[0] = 12.12
	fmt.Println(arrayTiga)

	// membuat array dan langsung memasukkan value
	//var arrayEmpat [5]string = [5]string{"Go", "Js", "PHP", "Python", "Java"}
	//var arrayEmpat = [5]string{"Go", "Js", "PHP", "Python", "Java"}
	arrayEmpat := [5]string{"Go", "Js", "PHP", "Python", "Java"}
	fmt.Println(arrayEmpat[1])
	fmt.Println(len(arrayEmpat))

	for i := 0; i < len(arrayEmpat); i++ {
		fmt.Println(arrayEmpat[i])
	}

	// arrayEmpat[2] = "Ruby"
	// for _, lang := range arrayEmpat {
	// 	fmt.Println(lang)
	// }

	// array 2 dimensi
	arr2dimensi := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	arr2dimensi[0][1] = 10
	fmt.Println(arr2dimensi)

	// for index, items := range arr2dimensi {
	// 	fmt.Printf("Index ke %v isinya: \n", index)
	// 	for _, item := range items {
	// 		fmt.Println(item)
	// 	}
	// }
}
