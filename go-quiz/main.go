// // package main

// // import (
// // 	"fmt"
// // 	"os"
// // 	"sync"
// // )

// // func main() {
// // 	// ask name , 5 questions with 30 seconds timeout , atlast greeting with name and score
// // 	fmt.Println("Welcome to GO-Quiz")
// // 	var username string
// // 	fmt.Println("Enter your Name")
// // 	fmt.Scan(&username)

// // 	type questionandans struct {
// // 		question string
// // 		answer   int
// // 	}

// // 	var qanda = []questionandans{
// // 		{"1+2", 3},
// // 		{"9x9", 81},
// // 		{"7/9", 7},
// // 		{"4 times 4", 16},
// // 		{"Value of pi x 100", 314},
// // 	}

// // 	var answers []int
// // 	var wg sync.WaitGroup
// // 	for _, val := range qanda {
// // 		// timer := time.After(time.Second * 30)

// // 		wg.Add(1)
// // 		var userInput int
// // 		fmt.Println(val.question)
// // 		go func() {
// // 			_, err := fmt.Scan(&userInput)
// // 			if err != nil {
// // 				fmt.Println("Error: Please enter a valid integer.")
// // 				os.Exit(1)
// // 			}
// // 			defer wg.Done()
// // 		}()
// // 		// select {
// // 		// case <-timer:
// // 		// 	fmt.Println("Times up")
// // 		// 	answers = append(answers, 0)
// // 		// 	continue
// // 		// case <-time.After(time.Second * 30):

// // 		// }

// // 		answers = append(answers, userInput)
// // 		wg.Wait()
// // 	}

// // 	result := 0
// // 	wg.Add(1)
// // 	go func() {
// // 		if qanda[0].answer == answers[0] {
// // 			result++
// // 		}
// // 		if qanda[1].answer == answers[1] {
// // 			result++
// // 		}
// // 		if qanda[2].answer == answers[2] {
// // 			result++
// // 		}
// // 		if qanda[3].answer == answers[3] {
// // 			result++
// // 		}
// // 		if qanda[4].answer == answers[4] {
// // 			result++
// // 		}
// // 		defer wg.Done()
// // 	}()

// // 	// fmt.Println(answers)
// // 	fmt.Println(result)
// // 	wg.Wait()
// // }

// // package main

// // import (
// // 	"fmt"
// // 	"sync"
// // 	"time"
// // )

// // type questionandans struct {
// // 	question string
// // 	answer   int
// // }

// // var answers []int

// // var qanda = []questionandans{
// // 	{"1+2", 3},
// // 	{"9x9", 81},
// // 	{"7/9", 7},
// // 	{"4 times 4", 16},
// // 	{"Value of pi x 100", 314},
// // }

// // // channel := make(chan int)
// // var wg sync.WaitGroup
// // var timer = time.After(time.Second * 5)
// // var mutex sync.Mutex

// // func main() {

// // 	for i := 0; i < len(qanda); i++ {
// // 		wg.Add(1)

// // 		go quiz(i, &wg)

// // 	}

// // 	wg.Wait()

// // 	result := 0
// // 	for i := 0; i < len(qanda); i++ {
// // 		if i < len(answers) && qanda[i].answer == answers[i] {
// // 			result++
// // 		}
// // 	}

// // 	fmt.Println("your score is ", result)

// // }

// // func quiz(i int, wg *sync.WaitGroup) {

// // 	var response int
// // 	userInput := make(chan int, 1)
// // 	fmt.Println(qanda[i].question)
// // 	// wg.Add(1)
// // 	go func() {
// // 		fmt.Scan(&response)
// // 		userInput <- response
// // 		// defer wg.Done()
// // 		time.Sleep(time.Second * 5)
// // 	}()

// // 	select {
// // 	case <-timer:
// // 		fmt.Println("Time's up. Appending 0.")
// // 		mutex.Lock()
// // 		answers = append(answers, 0)
// // 		mutex.Unlock()
// // 	case response := <-userInput:
// // 		fmt.Println("nnnnn.")
// // 		mutex.Lock()
// // 		answers = append(answers, response)
// // 		mutex.Unlock()
// // 	}

// // 	defer wg.Done()

// // }

// // package main

// // import "fmt"

// // func main() {
// // 	type questionandans struct {
// // 		question string
// // 		answer   int
// // 	}

// // 	// var answers []int

// // 	var qanda = []questionandans{
// // 		{"1+2", 3},
// // 		{"9x9", 81},
// // 		{"45/9", 5},
// // 		{"4 times 4", 16},
// // 		{"Value of pi x 100", 314},
// // 	}

// // 	for i := 0; i < len(qanda); i++ {
// // 		fmt.Println(qanda[i].question)

// // 	}
// // }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type questionandans struct {
// 	question string
// 	answer   int
// }

// var answers []int

// var qanda = []questionandans{
// 	{"1+2", 3},
// 	{"9x9", 81},
// 	{"45/9", 5},
// 	{"4 times 4", 16},
// 	{"Value of pi x 100", 314},
// }

// func main() {

// 	var wg sync.WaitGroup

// 	for i := 0; i < len(qanda); i++ {
// 		wg.Add(1)
// 		go quiz(i, &answers, &wg)
// 	}

// 	wg.Wait()

// 	fmt.Println("Collected answers:", answers)
// }

// func quiz(i int, answers *[]int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	var response int
// 	userInput := make(chan int, 1)
// 	timer := time.After(5 * time.Second)

// 	fmt.Println(qanda[i].question)

// 	go func() {
// 		fmt.Scan(&response)
// 		userInput <- response
// 	}()

// 	select {
// 	case <-timer:
// 		fmt.Println("Time's up. Appending 0.")
// 		*answers = append(*answers, 0)
// 	case response := <-userInput:
// 		fmt.Println("Received user input:", response)
// 		*answers = append(*answers, response)
// 	}
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

type questionandans struct {
	question string
	answer   int
}

var answers []int

var qanda = []questionandans{
	{"1+2", 3},
	{"9x9", 81},
	{"45/9", 5},
	{"4 times 4", 16},
	{"Value of pi x 100", 314},
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < len(qanda); i++ {
		wg.Add(1)
		go quiz(i, &answers, &wg)
		wg.Wait() // Wait for user input before proceeding to the next question
	}

	fmt.Println("Collected answers:", answers)
}

func quiz(i int, answers *[]int, wg *sync.WaitGroup) {
	defer wg.Done()

	var response int
	userInput := make(chan int, 1)
	timer := time.After(5 * time.Second)

	fmt.Println(qanda[i].question)

	go func() {
		fmt.Scan(&response)
		userInput <- response
	}()

	select {
	case <-timer:
		fmt.Println("Time's up. Appending 0.")
		*answers = append(*answers, 0)
	case response := <-userInput:
		fmt.Println("Received user input:", response)
		*answers = append(*answers, response)
	}
}
