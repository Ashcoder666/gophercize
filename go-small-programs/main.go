// pangram
// package main

// import "fmt"

// func main() {
// 	letterMap := make(map[rune]bool)

// 	string := "abcdefghi jklmopqrstuvuuw  xyz"

// 	for _, char := range string {
// 		if 'a' <= char && 'z' >= char {
// 			letterMap[char] = true
// 		}
// 	}

// 	fmt.Println(len(letterMap))

// }

// reverse a string
package main

import "fmt"

func main() {
	stringg := "njan nee"
	rev := ""
	for i := len(stringg) - 1; i >= 0; i-- {
		// fmt.Println(string(stringg[i]))

		rev += string(stringg[i])
	}

	fmt.Println(rev)
}
