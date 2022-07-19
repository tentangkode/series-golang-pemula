package main

import "fmt"

func main() {

	var numbers []int = []int{2, 5, 1, 4, 6, 7, 2, 7, 2, 7}

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("value %d \n", numbers[i])
	}

	// for i, element := range numbers {
	// 	fmt.Printf("index %d: %d \n", i, element)
	// }

	// for _, element := range numbers {
	// 	fmt.Printf("value %d \n", element)
	// }

	// var buah map[string]int = map[string]int{
	// 	"apel":   10,
	// 	"pisang": 20,
	// 	"jeruk":  20,
	// 	"mangga": 10,
	// }

	// for key, item := range buah {
	// 	fmt.Printf("%s ada %d \n", key, item)
	// }
}
