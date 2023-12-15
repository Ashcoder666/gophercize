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
// package main

// import "fmt"

// func main() {
// 	stringg := "njan nee"
// 	rev := ""
// 	for i := len(stringg) - 1; i >= 0; i-- {
// 		// fmt.Println(string(stringg[i]))

// 		rev += string(stringg[i])
// 	}

// 	fmt.Println(rev)
// }

// Remove Duplicates from Sorted Array:

// Given a sorted array, remove the duplicates in-place such that each element appears only once and return the new length. Do not allocate extra space for another array, and modify the input array in-place.
// Example:

// makefile
// Copy code
// Input: nums = [1, 1, 2]
// Output: [1, 2]

// package main

// import "fmt"

// func main(){
// 	arr := []int{1,2,2,6,7,7,9,11} // 1,2,6,7,9,11 --> output

// 	for i:=0 ; i< len(arr) ; i++{
// 		for j:=i+1 ; j<len(arr) ; j++{
// 			if arr[i] == arr[j]{
// 				arr = append(arr)
// 			}
// 		}
// 	}
// }

// Initialize two pointers
slow := 0

// Iterate through the array with the fast pointer
for fast := 1; fast < len(nums); fast++ {
	// If the current element is different from the previous one, update the slow pointer and copy the element
	if nums[fast] != nums[slow] {
		slow++
		nums[slow] = nums[fast]
	}
}

// The length of the new array is the position of the slow pointer + 1
return slow + 1

// Find the Maximum Product of Two Integers:

// Given an array of integers, find the maximum product that can be obtained by multiplying two different integers. The array may contain both positive and negative integers.
// Example:

// vbnet
// Copy code
// Input: nums = [2, 3, -2, 4]
// Output: 6 (as the maximum product can be obtained by multiplying -2 and 3)
// Rotate Array:

// Given an array, rotate the array to the right by k steps, where k is a non-negative integer.
// Example:

// makefile
// Copy code
// Input: nums = [1, 2, 3, 4, 5, 6, 7], k = 3
// Output: [5, 6, 7, 1, 2, 3, 4]
