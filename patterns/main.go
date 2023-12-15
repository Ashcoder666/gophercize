package main

import "fmt"

func main() {
	thirdPattern(5)
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

func thirdPattern(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j + 1)
		}
		fmt.Println()
	}
}
