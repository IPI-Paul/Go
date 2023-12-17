package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

var pl = fmt.Println
var pf = fmt.Printf
var tws = strings.TrimSpace

/*
Formats:
%d	Integer
%c	Character
%f	Float
%t	Boolean
%s	String
%o	Base 8
%x	Base 16
%v	Gueses based on data type
%T	Type of supplied value
File Settings:
Only one of the following.
O_RDONLY	open the file read-only
O_WRONLY	open the file write-only
O_RDWR		open the file read-write
These can be or'ed
O_APPEND	append data to the file when writing
O_CREATE	create a new file if none exists
O_EXCL		used with O_CREATE, file must not exist
O_SYNC		OPEN FOR SYNCHRONOUS I/O
O_TRUNC		truncate regular writable file when opened
*/

func main() {
	ex := Examples{}
	if len(os.Args) > 1 {
		runExample(os.Args[1], ex)
	} else {
	App:
		for {
			fmt.Print("Enter an example number from 1 to 50 or q to quit: ")
			reader := bufio.NewReader(os.Stdin)
			choice, err := reader.ReadString('\n')
			choice = tws(choice)
			if err != nil {
				fmt.Println("Error Entry!")
				continue
			}
			if !runExample(choice, ex) {
				break App
			}
		}
	}
}

func runExample(choice string, ex Examples) bool {
	val, _ := strconv.ParseInt(choice, 0, 42)
	num := fmt.Sprintf("%03d", val)
	if choice == "q" {
		fmt.Println("Bye")
		return false
	} else if val >= 1 && val <= 42 {
		fmt.Printf("Running example%v:\n", num)
		reflect.ValueOf(ex).MethodByName("Example" + num).Call(nil)
	} else {
		fmt.Println("Incorrect Entry. Bye!")
		return false
	}
	return true
}

type Examples struct{}

func (ex Examples) Example042() {
	cmd := exec.Command("bash", "-c", "(cd webapp; go run webapp.go)")
	stdout, _ := cmd.Output()
	pl(string(stdout))
}
func (ex Examples) Example041() {
	cmd := exec.Command("bash", "-c", "(cd app2; go test -v)")
	stdout, _ := cmd.Output()
	pl(string(stdout))
}
func (ex Examples) Example040() {
	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^\\s]+)", reStr)
	s, _ := regexp.Compile("(ape[^\\s]+)")
	pl(match, s.FindAllString(reStr, -1))
	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("([crmfp]at)")
	pl("MatchString:", r.MatchString(reStr2))
	pl("FindString:", r.FindString(reStr2))
	pl("Index:", r.FindStringIndex(reStr2))
	pl("All String:", r.FindAllString(reStr2, -1))
	pl("1st 2 Strings:", r.FindAllString(reStr2, 2))
	pl("All Submatch Index:", r.FindAllStringSubmatchIndex(reStr2, -1))
	pl(r.ReplaceAllString(reStr2, "Dog"))
}
func (ex Examples) Example039() {
	pl("Factorial 4 =", factorial(4))
}
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
func (ex Examples) Example038() {
	intSum := func(x, y int) int {
		return x + y
	}
	pl("5 + 4 =", intSum(5, 4))
	samp1 := 1
	changeVar := func() {
		samp1 += 1
	}
	changeVar()
	pl("samp1 =", samp1)
	useFunc := func(f func(int, int) int, x, y int) {
		pl("Answer:", (f(x, y)))
	}
	sumVals := func(x, y int) int {
		return x + y
	}
	useFunc(sumVals, 5, 8)
}
func (ex Examples) Example037() {
	var acct Account
	acct.balance = 100
	pl("Balance:", acct.GetBalance())
	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)
}

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}
func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		pl("Not enough money in account")
	} else {
		a.balance -= v
		pf("%d withdrawn. Balance: %d\n", v, a.balance)
	}
}
func (ex Examples) Example036() {
	nums1 := func(channel chan int) {
		channel <- 1
		channel <- 2
		channel <- 3
	}
	nums2 := func(channel chan int) {
		channel <- 4
		channel <- 5
		channel <- 6
	}
	channel1 := make(chan int)
	channel2 := make(chan int)
	go nums1(channel1)
	go nums2(channel2)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel2)
	pl(<-channel2)
	pl(<-channel2)
}
func (ex Examples) Example035() {
	printTo15 := func() {
		for i := 0; i <= 15; i++ {
			pl("Fun 1:", i)
		}
	}
	printTo10 := func() {
		for i := 0; i <= 10; i++ {
			pl("Fun 2:", i)
		}
	}
	go printTo15()
	go printTo10()
	time.Sleep(2 * time.Second)
}
func (ex Examples) Example034() {
	var kitty Animal
	kitty = Cat("kitty")
	kitty.AngrySound()
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	pl("Cat's Name:", kitty2.Name())
}

type Animal interface {
	AngrySound()
	HappySound()
}
type Cat string

func (c Cat) Attack() {
	pl("cat Attacks its Prey")
}
func (c Cat) Name() string {
	return string(c)
}
func (c Cat) AngrySound() {
	pl("Cat says Hisssss")
}
func (c Cat) HappySound() {
	pl("Cat say Purrrrrr")
}
func (ex Examples) Example033() {
	cmd := exec.Command("bash", "-c", "(cd app; go run main.go 2)")
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	pl(string(stdout))
}
func (ex Examples) Example032() {
	tspToML := func(tsp Tsp) ML {
		return ML(tsp * 4.92)
	}
	tbsToML := func(tbs TBs) ML {
		return ML(tbs * 14.79)
	}
	ml1 := ML(Tsp(3) * 4.92)
	pf("3 tsps = %.2f ML\n", ml1)
	ml2 := ML(TBs(3) * 14.79)
	pf("3 TBs = %.2f ML\n", ml2)
	pl("2 tsp + 4 tsp =", Tsp(2)+Tsp(4))
	pl("2 tsp > 4 tsp =", Tsp(2) > Tsp(4))
	pf("3 tsp = %.2f ML\n", tspToML(3))
	pf("3 tbs = %.2f ML\n", tbsToML(3))
	tsp1 := Tsp(3)
	pf("%.2f tsp = %.2f ML\n", tsp1, tsp1.ToMLs())
}

type Tsp float64
type TBs float64
type ML float64

func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}
func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}
func (ex Examples) Example031() {
	con1 := contact{
		"James",
		"Wang",
		"555-1212",
	}
	bus1 := business{
		"ABC Plumbing",
		"234 North St",
		con1,
	}
	bus1.info()
}

type contact struct {
	fName, lName, phone string
}
type business struct {
	name, address string
	contact
}

func (b business) info() {
	pf("Contact at %s is %s %s\n", b.name, b.contact.fName, b.contact.lName)
}
func (ex Examples) Example030() {
	rect1 := rectangle{10.0, 15.0}
	pl("Rectangle Area:", rect1.Area())
}

type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}
func (ex Examples) Example029() {
	type customer struct {
		name    string
		address string
		bal     float64
	}
	getCustInfo := func(c customer) {
		pf("%s owes us %.2f\n", c.name, c.bal)
	}
	newCustAdd := func(c *customer, address string) {
		c.address = address
	}
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 main st"
	tS.bal = 234.56
	getCustInfo(tS)
	newCustAdd(&tS, "123 South st")
	pl("Address:", tS.address)
	sS := customer{"Sally Smith", "123 Main", 0.0}
	pl("Name:", sS.name)
}
func (ex Examples) Example028() {
	pl("5 + 4=", getSumGen(5, 4))
	pl("5.6 + 4.7=", getSumGen(5.6, 4.7))
}

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}
func (ex Examples) Example027() {
	var heroes map[string]string
	heroes = make(map[string]string)
	villians := make(map[string]string)
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"
	villians["Lex Luther"] = "Lex Luther"
	superPets := map[int]string{1: "Krypto", 2: "Bat Hound"}
	pf("Batman is %v\n", heroes["Batman"])
	pl("Chip:", superPets[3])
	_, ok := superPets[3]
	pl("Is there a 3rd pet:", ok)
	for k, v := range heroes {
		pf("%s is %s\n", k, v)
	}
	delete(heroes, "The Flash")
}
func (ex Examples) Example026() {
	cmd := exec.Command("bash", "-c", "(cd app; go run main.go 1)")
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	pl(string(stdout))
}
func (ex Examples) Example025() {
	pl(os.Args)
	args := os.Args[2:]
	var iArgs = []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}
	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}
	pl("Max Value:", max)
}
func (ex Examples) Example024() {
	_, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		pl("File Doesn't Exist")
	} else {
		f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}
}
func (ex Examples) Example023() {
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Closes file when function out of scope
	defer f.Close()
	iPrimeArr := []int{2, 3, 5, 7, 1}
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime:", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
}
func (ex Examples) Example022() {
	getAverage := func(nums ...float64) float64 {
		var sum float64 = 0.0
		var numSize float64 = float64(len(nums))
		for _, val := range nums {
			sum += val
		}
		return (sum / numSize)
	}
	iSlice := []float64{11, 13, 17}
	pf("Average: %.3f\n", getAverage(iSlice...))
}
func (ex Examples) Example021() {
	dblArrVals := func(arr *[4]int) {
		for x := 0; x < 4; x++ {
			arr[x] *= 2
		}
	}
	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl(pArr)
}
func (ex Examples) Example020() {
	getSum := func(nums ...int) int {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		return sum
	}
	pl(getSum(1, 2, 3, 4, 5))
	getArrSum := func(arr []int) int {
		sum := 0
		for _, val := range arr {
			sum += val
		}
		return sum
	}
	vArr := []int{1, 2, 3, 4}
	pl("Array Sum:", getArrSum(vArr))
	changeVal := func(f3 int) int {
		f3 += 1
		return f3
	}
	f3 := 5
	pl("f3 before func:", f3)
	changeVal(f3)
	pl("f3 after func:", f3)
	changeVal2 := func(myPtr *int) {
		*myPtr = 12
	}
	f3 = 10
	pl("f3 before func:", f3)
	changeVal2(&f3)
	pl("f3 after func:", f3)
	f4 := 10
	var f4Ptr *int = &f4
	pl("f4 Address:", f4Ptr)
	pl("f4 Valuie:", *f4Ptr)
	*f4Ptr = 11
	pl("f4 Valuie:", *f4Ptr)
	pl("f4 before func:", f4)
	changeVal2(&f4)
	pl("f4 after func:", f4)
}
func (ex Examples) Example019() {
	sayHello := func() {
		pl("Hello")
	}
	sayHello()
	getSum := func(x int, y int) int {
		return x + y
	}
	pl(getSum(5, 4))
	getTwo := func(x int) (int, int) {
		return x + 1, x + 2
	}
	pl(getTwo(5))
	getQuotient := func(x float64, y float64) (ans float64, err error) {
		if y == 0 {
			return 0, fmt.Errorf("You can't divide by zero")
		} else {
			return x / y, nil
		}
	}
	pl(getQuotient(5, 0))
	pl(getQuotient(5, 4))
}
func (ex Examples) Example018() {
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"
	pl("Slice Size:", len(sl1))
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}
	for _, x := range sl1 {
		pl(x)
	}
	sArr := [5]int{1, 2, 3, 4, 5}
	sl3 := sArr[0:2]
	pl("1st 3:", sArr[:3])
	pl("Last 3:", sArr[2:])
	sArr[0] = 10
	pl("sl3:", sl3)
	sl3[0] = 1
	pl("sArr:", sArr)
	sl3 = append(sl3, 12)
	pl("sl3:", sl3)
	pl("sArr:", sArr)
	sl4 := make([]string, 6)
	pl("sl4:", sl4)
	pl("sl4[0]:", sl4[0])
}
func (ex Examples) Example017() {
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		pf("Rune Array: %d\n", v)
	}
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("I'm a string:", bStr)
}
func (ex Examples) Example016() {
	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		pl(num)
	}
	var arr1 [5]int
	arr1[0] = 1
	arr2 := [5]int{1, 2, 3, 4, 5}
	pl("index 0:", arr2[0])
	pl("Arr Length:", len(arr2))
	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}
	for i, v := range arr2 {
		pf("%d: %d\n", i, v)
	}
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}
}
func (ex Examples) Example015() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	for true {
		fmt.Print("Guess a number between 0 and 50: ")
		pl("Random Number is:", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = tws(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			pl("Pick a Lower Value")
		} else if iGuess < randNum {
			pl("Pick a Higer Value")
		} else {
			pl("You Guessed it")
			break
		}
	}
}
func (ex Examples) Example014() {
	for x := 1; x <= 5; x++ {
		pl(x)
	}
	for x := 5; x >= 1; x-- {
		pl(x)
	}
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}
}
func (ex Examples) Example013() {
	pf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1, 1)
	pf("%9f\n", 3.14)
	pf("%.2f\n", 3.141592)
	pf("%9.f\n", 3.141592)
	sp1 := fmt.Sprintf("%9.f", 3.141592)
	pl(sp1)
}
func (ex Examples) Example012() {
	pl("Abs(-10) =", math.Abs(-10))
	pl("Pow(4, 2) =", math.Pow(4, 2))
	pl("Sqrt(16) =", math.Sqrt(16))
	pl("Cbrt(8) =", math.Cbrt(8))
	pl("Ceil(4.4) =", math.Ceil(4.4))
	pl("Floor(4.4) =", math.Floor(4.4))
	pl("Round(4.4) =", math.Round(4.4))
	pl("Log2(8) =", math.Log2(8))
	pl("Log10(100) =", math.Log10(100))
	// Get the log of e to the power of 2
	pl("Log(7.389) =", math.Log(7.389))
	pl("Max(5, 4) =", math.Max(5, 4))
	pl("Min(5, 4) =", math.Min(5, 4))
	// Convert degrees to radians
	r90 := 90 * math.Pi / 180
	// Convert radians to degrees
	d90 := r90 * (180 / math.Pi)
	pf("r90: %.2f\nd90: %.2f\n", r90, d90)
	pl("Sin(90) =", math.Sin(r90))
}
func (ex Examples) Example011() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random:", randNum)
}
func (ex Examples) Example010() {
	pl("5 + 4 =", 5+4)
	pl("5 - 4 =", 5-4)
	pl("5 * 4 =", 5*4)
	pl("5 / 4 =", 5/4)
	pl("5 % 4 =", 5%4)
	mInt := 1
	mInt += 1
	mInt++
	pl("mInt:", mInt)
	pl("Folat precision =", 0.111111111111111+0.111111111111111)
}
func (ex Examples) Example009() {
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())
}
func (ex Examples) Example008() {
	rStr := "abcdefg"
	pl("Rune Count", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		pf("%d) %#U: %c\n", i, runeVal, runeVal)
	}
}
func (ex Examples) Example007() {
	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)
	pl("Length:", len(sV2))
	pl("Contains Another:", strings.Contains(sV2, "Another"))
	pl("o index:", strings.Index(sV2, "o"))
	pl("Replace:", strings.Replace(sV2, "o", "0", -2))
	sV3 := "\nSome Words\n"
	sV3 = strings.TrimSpace(sV3)
	pl(sV3)
	pl("Split:", strings.Split("a-b-c-d", "-"))
	pl("Lower:", strings.ToLower(sV2))
	pl("Upper:", strings.ToUpper(sV2))
	pl("Prefix:", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix:", strings.HasSuffix("tacocat", "cat"))
}
func (ex Examples) Example006() {
	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Important Birthday")
	} else {
		pl("Not an Important Birthday")
	}
	pl("!true =", !true)
}
func (ex Examples) Example005() {
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)
	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))
	cV5 := 50000000
	cV6 := strconv.Itoa(cV5)
	pl(cV6)
	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}
	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)
}
func (ex Examples) Example004() {
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf('♿'))
}
func (ex Examples) Example003() {
	var vName string = "Derek"
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"
	v4 := 2.4
	pf("vName: %v\nv1: %v\nv2: %v\nv3: %v\nv4: %v", vName, v1, v2, v3, v4)
}
func (ex Examples) Example002() {
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
}
func (ex Examples) Example001() {
	pl("Hello Go")
}
