package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func add(x int, y int) int {
	return x + y
}

// [2]
func incrementSend(sendsSoFar int, sendsToAdd int) int {
	return sendsSoFar + sendsToAdd
}

// [3]
func getNames() (string, string) {
	return "BDR", "BAHADOU"
}

func main() {
	fmt.Println(concat("BDR", "Happy birthday"))
	fmt.Println(add(10, 4))
	// [2] Passing variables by value
	sendsSoFar := 400
	const sendsToAdd = 25
	sendsSoFar = incrementSend(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent ", sendsSoFar, "messages")

	// [3] Ignoring return values (by using _)
	firstName, _ := getNames()
	fmt.Println("my first name is", firstName)

}
