package main

import "fmt"

func main() {
	letterMap := make(map[rune]bool)

	string := "abcdefghi jklmopqrstuvuuw  xyz"

	for _, char := range string {
		if 'a' <= char && 'z' >= char {
			letterMap[char] = true
		}
	}

	fmt.Println(len(letterMap))

}
