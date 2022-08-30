package main

import "fmt"

// func isoddEven(num int) string {
// 	if num%2 == 0 {
// 		return fmt.Sprintf("The number of %v is even %v ", num, num)
// 	} else {
// 		return fmt.Sprintf("The number of %v is odd %v ", num, num)
// 	}
// }

func printOddEven(n int) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Printf("The number of  %v is even! \n", i)
		} else {
			fmt.Printf("The number of  %v is odd! \n", i)
		}
	}
}

func main() {

	// fmt.Println(isoddEven(4))
	printOddEven(10)

}
