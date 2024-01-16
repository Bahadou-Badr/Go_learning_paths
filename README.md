# Go Learning Paths
## Some Infos About Go


- The Default Value of all int, float, and rune is `0`
- The default value of strings is an empty  string `''`
- The default value of a boolean is `false`
- The default value of errors is `nil`
---------------------------------------------------
- when declaring constants we have to initialize a value
  
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
## Arrays
Fixed length, Same type, Indexable, and Contiguous in memory. 

- Declaring array and assigned values

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
	
    fmt.Println(ages) // output : [25 35 45]
```

## Slices

```go

    //declaring slice
    scores := []int{20, 30, 40}
    //change values
    scores[1] = 35
    //append a new value 
    scores = append(scores, 50)
    //slice ranges
    rangeOne := scores[1:3] //from the position 1 to 3 but not including the value of position 3
    rangeTwo := scores[1:]  //from the position 1 to the end
    rangeThree := scores[:3]  //from the start to the position 3 but not including the value of 3

    fmt.Println(rangeOne) // output : [35 40]
    fmt.Println(rangeTwo) // output : [35 40 50]
    fmt.Println(rangeThree) // output : [20 35 40]
	
```

### Loops
- For-each range loop

```go 
strings := []string{"hello", "world"}
for i, s := range strings {
	fmt.Println(i, s)
}
```
- Three-component loop
```go 
sum := 0
for i := 1; i < 5; i++ {
	sum += i
}
fmt.Println(sum) // 10 (1+2+3+4)
```

- While loop
```go 
n := 1
for n < 5 {
	n *= 2
}
fmt.Println(n) // 8 (1*2*2*2)
```

## Structs
You can think of structs as a way to define your own type
```go

// Create a struct, struct can hold mixed type
type gasEngine struct{
	mpg uint8
	gallons uint8
}

// [2] Create a function which is directly tied to the struct instance itself 
func (e gasEngine) milesLift() uint8 {
	return e.gallons * e.mpg
}

func main() {
	// [1] Create a variable with gasEngine type
	var myEngine gasEngine
	var myEngine gasEngine = gasEngine{25, 15}
	myEngine.mpg = 20
	fmt.Println(myEngine.mpg, myEngine.gallons)

	// [2]
	var myEngine gasEngine = gasEngine{25, 15}
	fmt.Printf("Total miles lift in tank : %v", myEngine.milesLift())
}
```
---------------------------------------------------
---------------------------------------------------
# Web Application:
- Module Paths for Downloadable Packages
If you’re creating a package or application which can be downloaded and
used by other people and programs, then it’s good practice for your
module path to equal the location that the code can be downloaded
from.
For instance, if your package is hosted at https://github.com/foo/bar
then the module path for the project should be github.com/foo/bar.

### Web Application Basics
Now that everything is set up correctly let’s make the first iteration of our
web application. We’ll begin with the three absolute essentials:
-The first thing we need is a handler. If you’re coming from an MVC-
background, you can think of handlers as being a bit like controllers.
They’re responsible for executing your application logic and for
writing HTTP response headers and bodies.
-The second component is a router (or servemux in Go terminology).
This stores a mapping between the URL patterns for your application
and the corresponding handlers. Usually you have one servemux for
your application containing all your routes.
-The last thing we need is a web server. One of the great things about
Go is that you can establish a web server and listen for incoming
requests as part of your application itself. You don’t need an external
third-party server like Nginx or Apache.
Let’s put these components together in the main.go file to make a
working application.

```go
package main
import (
"log"
"net/http"
)
// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
w.Write([]byte("Hello from Snippetbox"))
}
func main() {
// Use the http.NewServeMux() function to initialize a new servemux, then
// register the home function as the handler for the "/" URL pattern.
mux := http.NewServeMux()
mux.HandleFunc("/", home)
// Use the http.ListenAndServe() function to start a new web server. We pas
// two parameters: the TCP network address to listen on (in this case ":4000
// and the servemux we just created. If http.ListenAndServe() returns an er
// we use the log.Fatal() function to log the error message and exit.
log.Println("Starting server on :4000")
err := http.ListenAndServe(":4000", mux)
log.Fatal(err)
}
```
