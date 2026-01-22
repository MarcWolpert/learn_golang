package main

import (
	"bufio"
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"html/template"
	"io"
	"math"
	"math/big"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	RWX = 0755
)

func main() {
	closureExample(0, 10, 200)

	sequence := closureExample2()
	for range 5 {
		//iterates 5 times, i will still be available
		fmt.Println(sequence())
	}

	fmt.Println("Program ending.")

	//pointer example
	var ptr *int
	var a int = 10
	ptr = &a
	fmt.Println("This is a: ", a)
	fmt.Println("This is the address of a using ptr: ", ptr)
	//to get the actual value at the address
	fmt.Println("Value at ptr address: ", *ptr)

	for i, char := range "hello" {
		fmt.Printf("Character at %v position: %c with value: %d\n", i, char, char)
	}

	//utf-8 for runes (which is just an extension of strings beyond ascii)
	fmt.Println("Rune count: ", utf8.RuneCountInString("hello"))

	str_example := "hello"
	//doesn't work
	// str_example[0] = `g`
	rune_example := []rune(str_example)
	//example of rune literal
	rune_example[0] = 'g'
	var ch rune = 'a'
	ch = 'あ'
	fmt.Println(ch)

	// var name string
	// var age int

	// fmt.Println("What is your name and age?")
	// fmt.Scanln(&name, &age)

	// fmt.Printf("scanln - \nName: %v, Age: %v\n", name, age)

	// fmt.Println("What is your name and age?")

	// fmt.Scan(&name, &age)

	// fmt.Printf("scan - \nName: %v, Age: %v\n", name, age)

	// fmt.Println("What is your name and age?")
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Printf("scanf - \nName: %v, Age: %v\n", name, age)

	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	p := Person{
		firstName: "Marc",
		lastName:  "Wolpert",
		age:       29,
	}
	fmt.Printf("p values: %v %v %v\n", p.firstName, p.lastName, p.age)
	q := struct {
		firstName string
		lastName  string
		age       int
	}{firstName: "Josiah", lastName: "Jehosaphat", age: 56}
	fmt.Printf("Q values: %v %v %v\n", q.firstName, q.lastName, q.age)

	c := Car{
		make:  "Chevrolet",
		model: "Malibu",
		year:  2001,
	}
	fmt.Println(c.fullDetails())
	d := &c
	fmt.Println(d.fullDetailsPointer())
	fmt.Println("Incrementing year by 1 on c")
	c.year += 1
	fmt.Printf("c year: %d\nd year: %d\n", c.year, d.year)
	//#65 interfaces
	r := rect{width: 3, height: 4}
	cir := circle{radius: 5}
	measure(r)
	measure(cir)

	x, y := 1, 2
	fmt.Printf("X: %v\tY: %v\n", x, y)
	x, y = swap(x, y)
	fmt.Printf("X: %v\tY: %v\n", x, y)

	//#67: Generics
	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)

	printAll(intStack.elements)
	item, _ := intStack.pop()
	fmt.Println(item)
	printAll(intStack.elements)
	for !intStack.isEmpty() {
		item, _ = intStack.pop()
	}
	fmt.Println(intStack.isEmpty())

	//#70: Strings functions
	str5 := "Hello, 123 Go! 11"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)

	str6 := "Hello, 世界"
	fmt.Println(utf8.RuneCountInString(str6))

	//string builder for efficient string handling
	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	//convert builder to string
	result := builder.String()
	fmt.Println(result)

	//using Writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you")

	result = builder.String()
	fmt.Println(result)

	//reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	result = builder.String()
	fmt.Println(result)

	// htmlTmpl := htmlTemplate.New("example")

	tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")
	if err != nil {
		panic(err)
	}
	//define data for the welcome message template
	data := map[string]interface{}{
		"name": "John",
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	//however that is a clumsy way, this will do all the logic up until the Parse() and panic because it might throw an error
	tmpl = template.Must(template.New("example").Parse("Welcome, {{.name}}! Howsa are you doing?\n"))

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	//accept input from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	//will accept anything before the newline character as the input
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//now to put it in the template
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! We're glad you joined.",
		"notification": "{{.nm}}, you have a new notification {{.ntf}}",
		"error":        "Oops! An error occurred: {{.em}}",
	}

	//parse and store templates
	parsedTemplates := make(map[string]*template.Template)

	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	isExit := false
	for {
		//show menu
		fmt.Println("\nMenu: ")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit: ")
		fmt.Println("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template
		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": name}
		case "2":
			fmt.Println("Enter your notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{"nm": name, "ntf": notification}
		case "3":
			fmt.Println("Enter your error message: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{"nm": name, "em": errorMessage}
		case "4":
			fmt.Println("Exiting...")
			isExit = true
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
			continue
		}
		if isExit {
			break
		}
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template:", err)
		}

	}
	//#73: Regular expressions regex
	fmt.Println("He said, \n\"I am great\"")
	fmt.Println(`He said, "I am great".`)

	//compile a regex pattern to match email address
	re1 := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	//test stirngs
	emails := [2]string{"user@email.com", "invalid_email"}

	for i, v := range emails {
		//match strings with regex
		fmt.Printf("Email%v: %v\n", i, re1.MatchString(v))
	}

	date := "2024-07-30"

	//capturing groups
	re1 = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	//find all submatches
	submatches := re1.FindStringSubmatch(date)
	fmt.Println(submatches)
	for _, v := range submatches[1:] {
		fmt.Printf("%v\n", v)
	}

	//replace characters in target string
	str := "Hello World"
	re1 = regexp.MustCompile(`[aeiou]`)
	result = re1.ReplaceAllString(str, "*")
	fmt.Println(result)

	//i - case insensitive
	//m - multi line model
	//s - dot matches all

	//to accept flags, precede with question mark
	re = regexp.MustCompile(`(?i)go`)
	text := "Golang is going great"

	//match
	fmt.Println("Match:", re.MatchString(text))

	//#74: Time
	//gets the local time
	fmt.Println(time.Now())

	//gets specific time
	specificTime := time.Date(2024, time.July, 24, 12, 5, 5, 5, time.UTC)
	fmt.Println("Specific time: ", specificTime)

	//parse time
	parsedTime, err := time.Parse("2006-01-02", "2020-05-01") //Mon Jan 2 15:04:05 MST 2006
	if err != nil {
		panic("Error during time parsing 00")
	}
	parsedTime1, err := time.Parse("06-01-02", "20-05-01") //Mon Jan 2 15:04:05 MST 2006
	if err != nil {
		panic("Error during time parsing 01")
	}
	parsedTime2, err := time.Parse("06-1-2", "20-5-1") //Mon Jan 2 15:04:05 MST 2006
	if err != nil {
		panic("Error during time parsing 02")
	}
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)

	//formatting time on time type
	t := time.Now()
	//uses the reference format as an example
	fmt.Println("Formatted time: ", t.Format("Monday 06-01-02 04-15"))

	//manipulating time
	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	//rounding time
	fmt.Println("Rouded Time: ", t.Round(time.Hour))

	//converting time zones
	loc, _ := time.LoadLocation("Asia/Kolkata")
	t = time.Date(2024, time.July, 8, 14, 16, 40, 00, time.UTC)
	//convert this to specific time zome
	tLocal := t.In(loc)
	fmt.Println(tLocal)

	//similar to round but just does floor operation
	fmt.Println("Truncated Time: ", t.Truncate(time.Hour))

	//#75: Epochs - Specific count in terms of seconds or milliseconds from a defined starting point (which is UTC january 1, 1970 not counting leap seconds)
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix Time: ", unixTime)
	//converting unix time back to utc
	t = time.Unix(unixTime, 0)
	fmt.Println(t)
	fmt.Println("Time: ", t.Format("2006-01-02"))

	//#76: Time Formatting
	//Mon Jan 2 15:04:05 MST 2006 #Reference
	//parse an ISO 8601 format
	layout := "2006-01-02T15:04:05Z07:00"
	str = "2024-07-04T14:30:18Z"
	t, err = time.Parse(layout, str)
	if err != nil {
		fmt.Printf("Error parsing time %v", err)
		return
	}
	fmt.Println(t)

	str1 := "Jul 03, 2024 03:18 PM"
	layout1 := "Jan 02, 2006 03:04 PM"
	t1, err := time.Parse(layout1, str1)

	fmt.Println(t1)
	//always consider time zones when accessing

	//#77: Random Numbers
	//in Go, there's by default an auto-seed so determinism is by default but only if you access that seed can you see it (useful for debugging)
	random := rand.Intn(100)
	//this rand.Intn() is [0,n) so if you want the lower bound to be higher you have to add to it
	lowerBound := 5
	upperBound := 27
	random = rand.Intn(upperBound-lowerBound) + lowerBound
	fmt.Println(random)

	//to manually seed
	val := rand.New(rand.NewSource(1337))
	fmt.Println(val.Intn(101))

	//generate a crypto rand package
	//use when needing security, not just randomness
	bigCryptoInt, _ := crand.Int(crand.Reader, big.NewInt(100))
	fmt.Println(bigCryptoInt.Int64())

	//#78: Number Parsing
	numStr := "12345"
	num, _ := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(num)

	//to convert to an integer, base, bit size
	base10ThirtyTwobit, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(base10ThirtyTwobit)
	//parse a hex
	hexString := "FFFF"
	hexNum, err := strconv.ParseInt(hexString, 16, 32)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(hexNum)

	//#79: URL Parsing
	rawURL := "https://example.com:8080/path?query=param&name=marc&age=15#fragment"
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}
	fmt.Println("Scheme: ", parsedUrl.Scheme)
	fmt.Println("Host: ", parsedUrl.Host)
	fmt.Println("Port: ", parsedUrl.Port())
	fmt.Println("Path: ", parsedUrl.Path)
	fmt.Println("Query: ", parsedUrl.Query())
	fmt.Println("Fragment: ", parsedUrl.Fragment)
	//NOTE: Query returns a map with all the parsed strings
	for i, v := range parsedUrl.Query() {
		fmt.Println(i, v)
	}
	//to get a specific value
	queryParams := parsedUrl.Query()
	fmt.Println("Name: ", queryParams.Get("name"))

	//to build a URL
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "google.com",
		Path:   "/",
	}

	query := baseURL.Query()
	// expected format for key/value pairings
	query.Set("key", "value")
	query.Set("big", "true")
	baseURL.RawQuery = query.Encode()

	fmt.Println(baseURL.String())

	//setting it with a map instead
	values := url.Values{
		"key": []string{"value"},
		"big": []string{"false"},
	}
	query.Del("key")
	query.Del("big")
	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)
	baseURL.RawQuery = encodedQuery
	fmt.Println(baseURL)

	//#79: Buffered io
	//just a reader when not enclosed
	reader = bufio.NewReader(strings.NewReader("Hello, bufio packageeeee!\n"))
	//now you can generate some output with chunks
	dataByte := make([]byte, 20)
	n, err := reader.Read(dataByte)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Read %d bytes: |%s|\n", n, dataByte[:n])

	//reads a string, reads to a delimiter
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Read string: ", line)

	// struct writer interface={}

	//#82:Hashing/Cryptography
	password := "password123"
	hash := sha256.Sum256([]byte(password))
	fmt.Printf("Password: %v == %v == %x\n", password, hash, hash)
	//salting is an extra layer of security by combining an extra random value to the hash
	//protects against security attacks like rainbow tables
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	hashedPassword, err := hashPassword(password, salt)

	//store the salt + password in database
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt: ", saltStr)
	fmt.Println("Hashed password: ", hashedPassword)
	//above is simulated storing in database
	//below is how to decode and retrieve
	decodeSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	//verify password
	loginHash, err := hashPassword(password, decodeSalt)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	//compare stored hash with login hash
	if hashedPassword == loginHash {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Login failed. Please check user credentials.")
	}
	fmt.Printf("Original Salt: %x\n", salt)

	//#83:Writing Files
	filename := "test.txt"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file: ", err)
	}
	defer file.Close()
	//write data to file
	var someBytes []byte = []byte("Hello World\n")
	_, err = file.Write(someBytes)
	if err != nil {
		fmt.Println("Error writing bytes: ", err)
	}
	fmt.Println("Data has been written to file successfully.")

	//create a file and write a string to it
	filename = "writeString.txt"
	file, err = os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file: ", err)
	}
	defer file.Close()
	//write data to file
	someBytes = []byte("Hello World\n")
	_, err = file.WriteString(string(someBytes))
	if err != nil {
		fmt.Println("Error writing bytes: ", err)
	}
	_, err = file.WriteString("a\nb\nc\nd\ne\n")
	if err != nil {
		fmt.Println("Error writing string: ", err)
		return
	}
	fmt.Println("Data has been written to file successfully.")

	//#84:Reading Files
	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %v: %v", filename, err)
	}
	defer func() {
		fmt.Println("Closing open file. ")
		file.Close()
	}()
	fmt.Println("File opened successfully.")

	// //now reading the contents of the file
	// data_file := make([]byte, 1024)
	// //reads file into buffer
	// _, err = file.Read(data_file)
	// if err != nil {
	// 	fmt.Println("Error reading from file: ", err)
	// 	return
	// }
	// fmt.Printf("|File content: %v|\n", string(data_file))

	//read file line by line
	//create a scanner to read line by line
	scanner := bufio.NewScanner(file)
	//read line by line until EOF character
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line: |%v|\n", line)
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	//#85: Line Filters
	filename = "example.txt"
	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %v: %v\n", filename, err)
		return
	}
	defer file.Close()

	lineNumber := 1

	scanner = bufio.NewScanner(file)
	//keyword to filter lines
	keyword := "important"
	//read and filter lines
	for scanner.Scan() {
		line = scanner.Text()
		if strings.Contains(line, keyword) {
			//replace
			line = strings.ReplaceAll(line, keyword, "dougie")
			fmt.Printf("[Line# %3d]Filtered line: %v\n", lineNumber, line)
			lineNumber++
		}
	}
	err = scanner.Err()
	if err != nil {
		fmt.Printf("Error scanning file: %v", err)
		return
	}

	// #86: File Paths
	relativePath := "./data/file.txt"
	absolutePath, _ := filepath.Abs(".")
	joined := filepath.Join(absolutePath, relativePath)
	fmt.Printf("Joined path: %v to %v \n %v \n", absolutePath, relativePath, joined)

	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println(normalizedPath)
	filepath_x := "/home/user/docs/file.txt"
	dir_x, file_x := filepath.Split(filepath_x)
	//base is the last component of a filepath, can be directory or filename
	fmt.Printf("dir: %v\npath: %v\nbase: %v\n", dir_x, file_x, filepath.Base(filepath_x))
	//to get a bool is filepath.isAbs(string)
	//to get extension use filepath.Ext(string)
	no_extension := strings.TrimSuffix(filepath_x, filepath.Ext(filepath_x))
	fmt.Println("Without suffix: ", no_extension)

	//to get the relative position from string1 to string2, in terms of path, use the filepath.Rel(str1,str2)

	// #87: Directories
	checkError(os.Mkdir("subdir", RWX))
	fmt.Println("Directory made successfully.")
	// defer os.RemoveAll("subdir")

	os.WriteFile("subdir/file.txt", []byte(""), RWX)

	//recursively makes these
	checkError(os.MkdirAll("subdir/parent/child", RWX))

	//to read all the files in a directory
	result_x, err := os.ReadDir("subdir/parent")
	checkError(err)

	for _, entry := range result_x {
		fmt.Println(entry, entry.Name(), entry.IsDir(), entry.Type())
	}

	checkError(os.Chdir("subdir/parent/child"))

	result_x, err = os.ReadDir(".")
	checkError(err)

	for _, entry := range result_x {
		fmt.Println(entry, entry.Name(), entry.IsDir(), entry.Type())
	}

	//os.Getwd()
	//os.Chdir() uses relative path

	//filepath.WalkDir
	//reads all files and directories in the tree, including root and calls a function defined on each of them

}

// can put structs within structs, both at a basic nesting level and you can define them within each other using anonymous level
type Truck struct {
	details   Car
	rimSize   int
	rustLevel struct {
		oxidization int
		material    string
	}
}
type Car struct {
	make  string
	model string
	year  int
}

// this syntax is called a method receiver
// adds a method onto a struct
func (c_inside Car) fullDetails() string {
	return fmt.Sprintf("%d %v %v", c_inside.year, c_inside.make, c_inside.model)
}

func (c_inside *Car) fullDetailsPointer() string {
	return fmt.Sprintf("%d %v %v", c_inside.year, c_inside.make, c_inside.model)
}

func closureExample(lower int, upper int, sleepDuration int) {
	//the reasoning here is that this is a closure because the outer function remembers the value of i while the inner function has no access to
	fmt.Println("closure example started")
	for i := lower; i < upper; i++ {
		//the reason this works is because it's a loop variable and is assigned a new address every time the loop iterates. it may have the same alias the the 'i' in the loop control structure, but it's a different address
		i := i
		time.AfterFunc(200*time.Millisecond, func() {
			//depending on the intended use of this function, this could be seen as a bug as it only does the callback after the dealye
			fmt.Println("This is i: ", i)
		})
		//due to Go's concurrent scheduler, this line makes it so that they execute synchronously. The scheduler didn't care that they were scheduled so close together and just picked an arbitrary order.
		time.Sleep(50 * time.Millisecond)
	}
	time.Sleep(time.Duration((upper-lower)*sleepDuration) * time.Millisecond)
}

func closureExample2() func() int {
	i := 0
	fmt.Println("previous value of i", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}

// #64: Methods
func (Car) hello() {
	fmt.Println("This is a car.")
}
func (c *Car) hi() {
}

// #65: Interfaces
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}
func (c circle) perim() float64 {
	return math.Pi * c.radius * 2
}
func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// #67: Generics
// remember any is an alias for an interface and that can be any type
func swap[T any](a, b T) (T, T) {
	return b, a
}

// make a data structure with any of the same uniform type
type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if s.isEmpty() {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func printAll[T any](t []T) {
	if len(t) == 0 {
		fmt.Printf("Zero elements in this array of type %T.", t)
	}
	for _, v := range t {
		fmt.Println(v)
	}
}

// #68: Errors
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math: No real square root can be negative.")
	}
	return x, nil
}

// #69: Custom Errors
type customError struct {
	code    int
	message string
	er      error
}

// error returns the error message. implementing Error() method of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.message)
}
func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong.",
			er:      err,
		}
	}
	return nil
}
func doSomethingElse() error {
	return errors.New("internal error")
}

// #82 Salting
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	//use io to read byte slice directly into salt slice
	//this reads values from the reader which is random in this case, and then puts it in the byte slice up to the point where it's full
	//doesnt return the slice since it's pass by reference, so it's in place memory
	_, err := io.ReadFull(crand.Reader, salt)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	return salt, err
}

func hashPassword(password string, salt []byte) (string, error) {
	//basically appends two byte slices together if theres
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:]), nil
}

// #87: Directories
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
