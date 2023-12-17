package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
uint8: unsigned 8-bit integers (0 to 255)
uint16: unsigned 16-bit integers (0 to 65535)
uint32: unsigned 32-bit integers (0 to 4294967295)
uint64: unsigned 64-bit integers (0 to 1844674473709551615)
int8: unsigned 8-bit integers (-128 to 127)
int16: unsigned 16-bit integers (-32768 to 32767)
int32: unsigned 32-bit integers (-2147483648 to 2147483647)
int64: unsigned 64-bit integers (-9223372036854775808 to 9223372036854775807)
*/
func main() {
	var choice string
	ex := Examples{}
App:
	for {
		fmt.Print("Enter an example number from 1 to 32 or q to quit: ")
		fmt.Scan(&choice)
		val, _ := strconv.ParseInt(choice, 0, 64)
		num := fmt.Sprintf("%03d", val)
		if choice == "q" {
			fmt.Println("Bye")
			break App
		} else if val >= 1 && val <= 32 {
			fmt.Printf("Running example%v:\n", num)
			reflect.ValueOf(ex).MethodByName("Example" + num).Call(nil)
		} else {
			fmt.Println("Incorrect Entry. Bye!")
			break App
		}
	}
}

type Examples struct{}

func (ex Examples) Example032() {
	pizzaNum := 0
	pizzaName := ""
	_ = pizzaName
	makeDough := func(stringChan chan string) {
		pizzaNum++
		pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
		println("Make Dough and Send for Sauce")
		stringChan <- pizzaName
		time.Sleep(time.Millisecond * 10)
	}
	addSauce := func(stringChan chan string) {
		pizza := <-stringChan
		println("Add Sauce and Send", pizza, "for toppings")
		stringChan <- pizzaName
		time.Sleep(time.Millisecond * 10)
	}
	addToppings := func(stringChan chan string) {
		pizza := <-stringChan
		fmt.Println("Add Toppings to", pizza, "and ship")
	}
	stringChan := make(chan string)
	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)
		time.Sleep(time.Millisecond * 5000)
	}
}
func (ex Examples) Example031() {
	count := func(id int) {
		for i := 0; i < 10; i++ {
			fmt.Println(id, ":", i)
			time.Sleep(time.Millisecond * 1000)
		}
	}
	for i := 0; i < 10; i++ {
		go count(i)
	}
	time.Sleep(time.Millisecond * 11000)
}
func (ex Examples) Example030() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World\n")
	}
	handler2 := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Earth\n")
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)
	http.ListenAndServe(":8080", nil)
}
func (ex Examples) Example029() {
	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "250.5"

	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))
	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt)
	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newFloat)
}
func (ex Examples) Example028() {
	file, err := os.Create("samp.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("This is some random text")
	file.Close()
	stream, err := ioutil.ReadFile("samp.txt")
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)
}
func (ex Examples) Example027() {
	csvString := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ","))
	listOfLetters := []string{"c", "a", "b"}
	sort.Strings(listOfLetters)
	fmt.Println(listOfLetters)
	listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")
	fmt.Println(listOfNums)
}
func (ex Examples) Example026() {
	sampString := "Hello World"
	fmt.Println(strings.Contains(sampString, "lo"))
	fmt.Println(strings.Index(sampString, "lo"))
	fmt.Println(strings.Count(sampString, "l"))
	fmt.Println(strings.Replace(sampString, "l", "x", 3))
}
func (ex Examples) Example025() {
	getArea := func(shape Shape) float64 {
		return shape.area()
	}
	rect := Rectangle1{20, 50}
	circ := Circle{4}
	fmt.Println("Rectangle Area =", getArea(rect))
	fmt.Println("Circle Area =", getArea(circ))
}

type Shape interface {
	area() float64
}
type Rectangle1 struct {
	height float64
	width  float64
}
type Circle struct {
	radius float64
}

func (r Rectangle1) area() float64 {
	return r.height * r.width
}
func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (ex Examples) Example024() {
	rect1 := Rectangle{
		leftX:  0,
		topY:   50,
		height: 10,
		width:  10,
	}
	// Or rect1 := Rectangle{0, 50, 10, 10}
	fmt.Println("Rectangle is", rect1.width, "wide")
	fmt.Println("Area of rectangle is", rect1.area())
}

type Rectangle struct {
	leftX  float64
	topY   float64
	height float64
	width  float64
}

func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}
func (ex Examples) Example023() {
	changeXVal := func(x int) {
		x = 2
	}
	x := 0
	changeXVal(x)
	fmt.Println("x =", x)
	changeXValNow := func(x *int) {
		*x = 2
	}
	changeXValNow(&x)
	fmt.Println("x =", x)
	fmt.Println("Memory Address for x =", &x)
	changeYValNow := func(yPtr *int) {
		*yPtr = 100
	}
	yPtr := new(int)
	changeYValNow(yPtr)
	fmt.Println("y =", *yPtr)
}
func (ex Examples) Example022() {
	demPanic := func() {
		defer func() {
			fmt.Println(recover())
		}()
		panic("PANIC")
	}
	demPanic()
}
func (ex Examples) Example021() {
	safeDiv := func(num1, num2 int) int {
		defer func() {
			// recover from panic
			fmt.Println(recover())
		}()
		solution := num1 / num2
		return solution
	}
	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))
}
func (ex Examples) Example020() {
	printOne := func() {
		fmt.Println(1)
	}
	printTwo := func() {
		fmt.Println(2)
	}
	defer printTwo()
	printOne()
}
func (ex Examples) Example019() {
	fmt.Println(factorial(3))
}
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
func (ex Examples) Example018() {
	num3 := 3
	doubleNum := func() int {
		num3 *= 2
		return num3
	}
	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
}
func (ex Examples) Example017() {
	subtractThem := func(args ...int) int {
		finalValue := 0
		for _, value := range args {
			finalValue -= value
		}
		return finalValue
	}
	fmt.Println(subtractThem(1, 2, 3, 4, 5))
}
func (ex Examples) Example016() {
	next2Values := func(number int) (int, int) {
		return number + 1, number + 2
	}
	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)
}
func (ex Examples) Example015() {
	listNums := []float64{1, 2, 3, 4, 5}

	addThemUp := func(numbers []float64) float64 {
		sum := 0.0
		for _, val := range numbers {
			sum += val
		}
		return sum
	}

	fmt.Println("Sum:", addThemUp(listNums))
}
func (ex Examples) Example014() {
	presAge := make(map[string]int)
	presAge["TheodoreRoosevelt"] = 42
	fmt.Println(presAge["TheodoreRoosevelt"], len(presAge))
	presAge["John F. Kennedy"] = 43
	fmt.Println(len(presAge))
	delete(presAge, "John F. Kennedy")
	fmt.Println(len(presAge))
}
func (ex Examples) Example013() {
	numSlice := []int{5, 4, 3, 2, 1}
	numSlice2 := numSlice[3:5]
	fmt.Println("numSlice2[0] =", numSlice2[0])
	fmt.Println("numSlice2[:2] =", numSlice2[:2])
	numSlice3 := make([]int, 5, 10)
	copy(numSlice3, numSlice)
	fmt.Println(numSlice3)
	numSlice3 = append(numSlice3, 0, -1)
	fmt.Println(numSlice3)
}
func (ex Examples) Example012() {
	var favNums2 [5]float64
	favNums2[0] = 163
	favNums2[1] = 78557
	favNums2[2] = 691
	favNums2[3] = 3.141
	favNums2[4] = 1.618
	fmt.Println(favNums2[3])
	// Or
	for idx, v := range [5]float64{163, 78557, 691, 3.141, 1.618} {
		favNums2[idx] = v
	}
	fmt.Println(favNums2[3])
	favNums3 := [5]float64{1, 2, 3, 4, 5}
	for i, value := range favNums3 {
		fmt.Println(i, value)
	}
}
func (ex Examples) Example011() {
	yourAge := 5
	switch yourAge {
	case 16:
		fmt.Println("Go Drive")
	case 18:
		fmt.Println("Go Vote")
	default:
		fmt.Println("Go Have Fun")
	}
}
func (ex Examples) Example010() {
	yourAge := 18
	if yourAge >= 16 {
		fmt.Println("You Can Drive")
	} else {
		fmt.Println("You Can't Drive")
	}

	if yourAge >= 16 {
		fmt.Println("You Can Drive")
	} else if yourAge >= 18 {
		fmt.Println("You Can Vote")
	} else {
		fmt.Println("You Can Have Fun")
	}
}
func (ex Examples) Example009() {
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
}
func (ex Examples) Example008() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
func (ex Examples) Example007() {
	fmt.Println("true && false =", true && false)
	fmt.Println("true || false =", true || false)
	fmt.Println("!true =", !true && false)
}
func (ex Examples) Example006() {
	const pi float64 = 3.14159265
	var isOver40 bool = true
	if isOver40 {
		fmt.Printf("%f \n", pi)
		fmt.Printf("%.3f \n", pi)
		fmt.Printf("%T \n", pi)
		fmt.Printf("%t \n", isOver40)
		fmt.Printf("Int 100: %d \n", 100)
		fmt.Printf("Binary 100: %b \n", 100)
		fmt.Printf("Char 44: %c \n", 44)
		fmt.Printf("Hex 17: %x \n", 17)
		fmt.Printf("Scientific pi: %e \n", pi)
	}
}
func (ex Examples) Example005() {
	var (
		varA = 2
		varB = 3
	)
	fmt.Println(varA, varB)
	var myName string = "Derek Banas"
	fmt.Println(len(myName), myName+" is a robot")
}
func (ex Examples) Example004() {
	fmt.Println("6 + 4 =", 6+4)
	fmt.Println("6 - 4 =", 6-4)
	fmt.Println("6 * 4 =", 6*4)
	fmt.Println("6 / 4 =", 6/4)
	fmt.Println("6 % 4 =", 6%4)
}
func (ex Examples) Example003() {
	var numOne = 1.000
	var num99 = .9999
	fmt.Println("Floats are not always 100% accurate:", numOne-num99)
}
func (ex Examples) Example002() {
	var age int = 40
	var favNum float64 = 1.6180339
	randNum := 1
	fmt.Println(age, favNum, randNum)
}
func (ex Examples) Example001() {
	fmt.Println("Hello World!")
}
