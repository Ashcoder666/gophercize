package main

import (
	"fmt"
	"strconv"
)

func main() {
	palindrome(242)
}

func firstPattern(n int) {
	stringg := ""
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			stringg += "*"
		}
		fmt.Println(stringg)
		stringg = ""
	}
}

func secondPattern(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func thirdPattern(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j + 1)
		}
		fmt.Println()
	}
}

func fourthPattern(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

func fifthPattern(n int) {
	for j := 0; j < n; j++ {
		for i := n; i > j; i-- {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func sixthPattern(n int) {
	for i := 0; i < n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

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

func eighthPattern(n int) {
	num := n*2 - 1
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < num; k++ {
			fmt.Print("*")
		}
		for j := 1; j < i; j++ {
			fmt.Print(" ")
		}
		num -= 2
		fmt.Println()
	}
}

func ninthPattern() {
	num1 := 1
	num2 := 9
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < num1; k++ {
			fmt.Print("*")
		}
		num1 += 2
		fmt.Println()

	}
	for a := 1; a <= 5; a++ {
		for b := 1; b < a; b++ {
			fmt.Print(" ")
		}
		for c := 0; c < num2; c++ {
			fmt.Print("*")
		}
		num2 -= 2
		fmt.Println()
	}
}

func tenthPattern(n int) {
	for i := 1; i <= 2*n-1; i++ {
		stars := i
		if i > (2*n)/2 {
			stars = 2*n - i // magic happens here
		}
		for j := 0; j < stars; j++ {
			fmt.Print("*")

		}
		fmt.Println()
	}
}

func eleventhPattern(n int) {
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if j%2 == 0 {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		}
		fmt.Println()
	}
}

func twelthPattern(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		for k := 0; k < 2*n-2*i; k++ {
			fmt.Print(" ")
		}
		for l := i; l > 0; l-- {
			fmt.Print(l)
		}

		fmt.Println()
	}
}

func thirteenthPattern(n int) {
	num := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(num, " ")
			num++
		}
		fmt.Println()
	}
}

func fourteenthPattern(n int) {
	alpha := []string{"A", "B", "C", "D", "E"}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(alpha[j])
		}
		fmt.Println()
	}
}

func fifteenthPattern(n int) {
	alpha := []string{"A", "B", "C", "D", "E"}
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(alpha[j])
		}
		fmt.Println()
	}
}

func sixteenthPattern(n int) {
	alpha := []string{"A", "B", "C", "D", "E"}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(alpha[i])
		}
		fmt.Println()
	}
}

func seventeenthPattern(n int) {
	alpha := []string{"A", "B", "C", "D", "E"}
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Print(alpha[k])
		}
		for l := i - 1; l > 0; l-- {
			fmt.Print(alpha[l-1])
		}
		fmt.Println()
	}
}

func eighteenthPattern(n int) {
	alpha := []string{"A", "B", "C", "D", "E"}

	for i := 1; i <= n; i++ {
		for j := n - i; j <= i; j++ {
			fmt.Print(alpha[n-j-1])
		}
		fmt.Println()
	}
}

//Given an integer N, write a program to count the number of digits in N.

func countDigits(digit int) int {
	count := 0
	//first methos
	for digit > 0 {
		digit /= 10
		count++
	}

	//second method
	newDig := strconv.Itoa(digit)
	fmt.Println(len(newDig))
	return count
}

//Given a number N reverse the number and print it.

func reverse(n int) {
	num := 0
	for n != 0 {
		digit := n % 10

		num = num*10 + digit

		n = n / 10
	}

	fmt.Println(num)
}

func palindrome(n int) {
	temp := n
	num := 0

	for n != 0 {
		digit := n % 10

		num = num*10 + digit

		n = n / 10

	}

	if temp == num {
		fmt.Println("palindrome")
	} else {
		fmt.Println("not palindrome")
	}
}
