package main

import (
	fmt "fmt"
	"math"
	"math/rand"
	foo "net/http" //changes the import
	"slices"
	"time"

	//to be a named import
	"errors"
	"maps"
	"os"
)

func main() {
	fmt.Println("This is a simple Go program.")
	resp, err := foo.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)

	// variables
	// 	integers : 0
	// floats
	// copmlex
	// bools : false
	// strings : ""
	// constants
	// arrays
	// pointers, slices, maps, functions, and structs : nil
	// channels
	// json
	// text/html templates

	// var age int //uninitialized
	var age_jenny int = 15 //used for initialized with value
	age_mike := 16         //type is inferred by the go compiler when using this operator
	age_jenny = 16         //reassigns

	fmt.Println("Jenny's Age: ", age_jenny)
	fmt.Println("Mike's Age: ", age_mike)

	//examples of naming conventions
	const MAXRETRIES = 5

	//EmployeeGoogle would be how you name it with multiple
	type Employee struct {
		FirstName string
		LastName  string
		Age       int
	}

	type EmployeeApple struct {
		FirstName string
		LastName  string
		Age       int
	}

	var employeeID = 1001
	fmt.Println("Employee ID: ", employeeID)

	//constants
	const SOMECONST int = 5
	const SOMECONST_2 = "yellow"

	//consts are evaluated at compile time
	//they can be typed or untyped
	const (
		CONST_1 = 5
		CONST_2 = 5.555
		CONST_3 = "blue"
		CONST_5
	)
	//overflow
	var uMaxInt uint64 = (uint64(math.Pow(2, 64) - 1))
	fmt.Println(uMaxInt)

	//underflow because precision
	var smallFloat float64 = 1.0e-23
	fmt.Println(smallFloat / math.MaxFloat64)

	//for loops
	//walrus first, simple over a range
	for i := 1; i <= 5; i++ {
		fmt.Println((i))
	}

	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		//continue keyword
		if index == 2 {
			continue
		}
		//%v is general and %d is specific for numbers
		fmt.Printf("Index %d, Value %d\n", index, value)
		if index > 3 {
			//break keyword
			break
		}
	}
	//simple way to declare a range
	for i := range 10 {
		fmt.Println("shortened range: ", 10-i)
	}

	//for as while looop
	i := 1
	for i <= 5 {
		fmt.Println("Iteration: ", i)
		i++
	}

	//Infinite loop
	// for {
	// 	fmt.Println("Hello")
	// }

	//for as while loop with break
	sum := 0
	for {
		sum += 10
		fmt.Println("Sum: ", sum)
		if sum >= 50 {
			break
		}
	}

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	//generate a random number between 1 and 100
	target := random.Intn(100) + 1
	fmt.Println("Welcome to the guessing game!\nI've chosen a number between 1 and 100\nCan you guess what it is?\n")
	tries := 0
	for {
		if tries > 9 {
			fmt.Println("Maximum attempts tried. The correct answer was: ", target)
			break
		}
		if tries > 0 {
			tries += 1
			continue
		}
		guess := 0
		fmt.Scan(&guess)
		//if else , if
		if guess == target {
			fmt.Println("Correct! You win!")
			break
		} else if guess != target {
			fmt.Println("Try again!")
			tries += 1
			fmt.Printf("Try %d/%d\n", tries+1, 10)
		}
	}

	//allowable operators
	// !
	// ||
	// &&
	// & |
	// ^
	// &^ which is x and not y
	// called the "bit clear" operator
	// <<
	// >>
	// ==  != < > <= >=
	// % modulus
	// NO exponentiation operator

	//switch
	a := true
	b := false
	switch a == b {
	case true:
		//no break needed
		fmt.Println("They are equal.")
	case false:
		fmt.Println("It's false, going to the default case.")
		fallthrough
	default:
		fmt.Println("Using the default case.")
	}

	number := 15
	switch {
	case number < 10:
	case number > 10 && number < 20:
		switch number {
		case 13, 14, 15:
			fmt.Println("Number is between 13 and 15.")
		default:
			fmt.Println("Number is not between 13 and 15 but lies between 10 and 20")
		}
	default:
		fmt.Println("Number is greater than 20.")
	}

	//you can access the type
	//type switch does not allow fallthrough
	var basicType = func(s interface{}) float64 {
		switch s.(type) {
		case int, float64, string:
			fmt.Println("type is int, float64, OR string")
		case nil:
			fmt.Println("interface is nil")
		default:
			fmt.Println("interface is not int, float64, string, or nil")

		}
		return 1e-20
	}
	basicType(15)
	basicType(nil)

	//arrays
	//var arrayName [size] type
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
		fmt.Println("Array position: ", arr[i])
	}
	//different way to declare array
	fruits := []string{"apple", "banana", "octopus", "jellybean"}
	fmt.Println("Fruits array: ", fruits)

	//underscore is a blank identifier
	//no use for the value it just discards
	//basically skips the load instruction for an alias
	for _, v := range fruits {
		fmt.Printf(" Value: %s\n", v)
	}

	//copy by address
	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int
	copiedArray = &originalArray
	copiedArray[2] = 100

	fmt.Printf("Copied Array[2]: %d\nOriginal Array[2]: %d\n", copiedArray[2], originalArray[2])

	// slices are same as array but mention the type
	//var sliceName[]ElementType
	// var numbers_slice []int
	// var numbers1 = []int{1, 2, 3}
	// numbers2 := []int{9, 8, 7}

	//will initialize an array on the backend but return the reference to that array
	slice := make([]int, 5)
	//array
	a2 := [5]int{1, 2, 3, 4, 5}
	//slice, start to stop before
	slice = a2[1:4]

	fmt.Println(slice)
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	fmt.Println("Slices: %d %d\n", slice[1], sliceCopy[1])
	slice[1] = 52
	fmt.Println("Slices new%d %d\n", slice[1], sliceCopy[1])

	//not referencing any array
	//nilslice
	var nilSlice []int
	nilSlice = []int{1, 2, 3, 4}

	//more convenient method to iterate over slice
	for i, v := range nilSlice[1:] {
		fmt.Println(i, v)
	}

	if slices.Equal(slice, sliceCopy) {
		fmt.Println("slice 1 == sliceCopy")
	} else {
		fmt.Printf("slice and copy are not the same, lengths are: %d %d\n", len(slice), len(sliceCopy))
		fmt.Printf("Slices: \n%v\n%v\n", slice, sliceCopy)
	}

	//multidimensional slice/array
	twoD := make([][]int, 3)
	alpha := []int{1, 3, 5}
	beta := []int{1e1, 1e2, 1e3, 1e4}
	for i, v1 := range alpha {
		innerLen := len(beta)
		//allocating an array
		twoD[i] = make([]int, innerLen)
		for j, v2 := range beta {
			twoD[i][j] = v1 + v2
		}
	}
	fmt.Println("twoD= ", twoD)
	//semantic error
	// twoD_slice := twoD[0:2][0:2]
	// fmt.Println("twoD_slice: ", twoD_slice)
	//want a 2x2 matrix, so have to slice in a loop similar to the above multidimensional slice/array example
	dimen_1 := 2
	dimen_2 := 3
	twoD_slice := make([][]int, dimen_1)
	for i := 0; i < dimen_1; i++ {
		twoD_slice[i] = twoD[i][0:dimen_2]
	}
	fmt.Println(twoD_slice)
	fmt.Println("Capacity of the underlying array of twoD: ", cap(twoD[0:2]))

	// maps (key:value)
	// var mapVariable map[keyType]valueType
	var m = make(map[int]string)
	var m2 = map[int]string{1: "1"}
	m2[2] = "2"
	m[1] = "1"
	fmt.Println(m, m2)
	//to delete a key from a map
	delete(m2, 2)
	fmt.Println(m2)
	//to get a bool to see if a key has a value
	_, ok := m2[2]
	fmt.Println("Is m2[2] there?: ", ok)
	if maps.Equal(m, m2) {
		fmt.Println("Maps are equal.")
	}
	for _, v := range m2 {
		fmt.Println(v)
	}
	//make a nil map, because there's no constructor of the ADT behind the map
	var m3 map[string]string
	//meanwhile this will make the map with the ADT behind it
	m4 := make(map[string]string)
	m4["hello"] = "goodbye"

	if m3 == nil {
		fmt.Println("This map is nil")
	} else {
		fmt.Println("This map is not nil")
	}

	//RANGE KEYWORD for iterables
	fmt.Printf("add: %d \nminus: %d\n", innerFunc(7, 5, add), innerFunc(7, 5, minus))

	//variadic function call with ellipsis
	var variadicArray []float32 = make([]float32, 3)

	variadic(1, 2, variadicArray...)

	deferExample()
	defer fmt.Println("Deferred on main for panic example")
	// panicExample(-1)

	recoverExample(-1)
	fmt.Println("Returned from process.")

	//os exit example
	osExitExample()
}

// functions must be at top level
func add(v int, w int) int   { return v + w }
func minus(v int, w int) int { return v - w }
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// generic example
func must[T any](t T, err error) T {
	if err != nil {
		panic(fmt.Errorf("Error in ..."))
	}
	return t
}

// multiple return values
func multipleReturnValue[T any, U any, V any](t T) (T, error) {
	switch any(t).(type) {
	case int:
		return t, nil
	default:
		return t, errors.New("Unable to cast to an integer.")
	}
}

// functions can be passed in as parameters here
func innerFunc(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// variadic functions
// ... ellipsis
func variadic(a int, b int, c ...float32) {
	//to access all members of c:
	var total float64 = 0
	x := float64(a)
	y := float64(b)
	for _, v := range c {
		total += (x + y) * float64(v)
	}
}

// defer example
func deferExample() {
	defer fmt.Println("This is the deferred.")
	//deferred statements are LIFO so they live on a stack
	defer fmt.Println("This is the second deferred.")
	var x int = 5
	//evaluated immediately
	defer fmt.Println("This is the third defer with the argument: x=", x)
	x = 6
	fmt.Println("Normal execution statement.")
}

// panic exmaple
func panicExample(input int) {
	if input < 0 {
		panic("Input is less than 0. It must be positive.")
	} else if input == 2048 {
		fmt.Println("Specific panic example to be noted for example usage. Please change input to -1 for true panic")
	}
	fmt.Println("Processing input: ", input)
}

// recover example
func recoverExample(r int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic from value: ", r)
		}
	}()

	fmt.Println("Start process")
	panic("Something went wrong!")
	fmt.Println("End process")
}

func osExitExample() {
	fmt.Println("os exit example.")

	//will not execute the defer
	defer fmt.Println("Deferred os.exit example")

	//exit with code 1
	os.Exit(1)

	fmt.Println("Wont go because exit process.")
}

func init() {
	fmt.Println("Initializing package...")
}

func init() {
	fmt.Println("Initializing package 1...")
}
