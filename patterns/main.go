package main

import "fmt"

func main() {
	seventhPattern(4)
}

// func firstPattern(n int) {
// 	stringg := ""
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			stringg += "*"
// 		}
// 		fmt.Println(stringg)
// 		stringg = ""
// 	}
// }

// func secondPattern(n int) {
// 	for i := 0; i < n; i++ {
// 		for j := 0; j <= i; j++ {
// 			fmt.Printf("*")
// 		}
// 		fmt.Println()
// 	}
// }

// func thirdPattern(n int) {
// 	for i := 0; i < n; i++ {
// 		for j := 0; j <= i; j++ {
// 			fmt.Print(j + 1)
// 		}
// 		fmt.Println()
// 	}
// }

// func fourthPattern(n int) {
// 	for i := 1; i <= n; i++ {
// 		for j := 0; j < i; j++ {
// 			fmt.Print(i)
// 		}
// 		fmt.Println()
// 	}
// }

// func fifthPattern(n int) {
// 	for j := 0; j < n; j++ {
// 		for i := n; i > j; i-- {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}

// }

// func sixthPattern(n int) {
// 	for i := 0; i < n; i++ {
// 		for j := 1; j <= n-i; j++ {
// 			fmt.Print(j)
// 		}
// 		fmt.Println()
// 	}
// }

// xmas tree pattern
func seventhPattern(n int) {

	num := 1

	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < num; k++ {
			fmt.Print("*")
		}
		for j := 0; j < n-i; j++ { // optional space after wards
			fmt.Print(" ")
		}
		num += 2
		fmt.Println()
	}
}
