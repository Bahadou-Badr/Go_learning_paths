# Go Learning Paths
## Working with primitive data type
- Declaring Variables
```go
    var a int
    a = 42
    fmt.Println(a)
    
    // other way
    var b int = 55
    fmt.Println(b)
    
    // other way
    name := "Badr"
    fmt.Println(name) 
	
	//we have also float32 / float64 /string / bool / uint
	
```
- Pointers
```go
	var name *string = new(string)

	*name = "Badr"
	fmt.Println(*name)
```

## Functions

```go
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
```
## Array
- Declaring array

```go
    //declaring array with assigned values (3 ways)
	var ages [3]int = [3]int{20, 30, 40}
    var ages = [3]int{20, 30, 40}
    ages := [3]int{20, 30, 40}
	
    //declaring array 
	//then assigned values
    ages := [3]int{}
    ages[0] = 25
    ages[1] = 35
    ages[2] = 45
	
    fmt.Println(ages)
```