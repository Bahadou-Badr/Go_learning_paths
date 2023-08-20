package main

import "fmt"

func main() {

	//var ages [3]int = [3]int{20, 30, 40}
	//var ages = [3]int{20, 30, 40}
	//ages := [3]int{20, 30, 40}
	ages := [3]int{}
	ages[0] = 25
	ages[1] = 35
	ages[2] = 45
	fmt.Println(ages)

	//Slices

	//declaring slice
	scores := []int{20, 30, 40}
	//change values
	scores[1] = 35
	//append a new value
	scores = append(scores, 50)
	//slice ranges
	rangeOne := scores[1:3]  //from the position 1 to 3 but not including the value of position 3
	rangeTwo := scores[1:]   //from the position 1 to the end
	rangeThree := scores[:3] //from the start to the position 3 but not including the value of 3

	fmt.Println(rangeOne)   // output : [35 40]
	fmt.Println(rangeTwo)   // output : [35 40 50]
	fmt.Println(rangeThree) // output : [20 35 40]
}
