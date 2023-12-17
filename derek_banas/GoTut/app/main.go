package main

import (
	"bufio"
	stuff "example/project/mypackage"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var tws = strings.TrimSpace

func main() {
	ex := Examples{}
	if len(os.Args) > 1 {
		runExample(os.Args[1], ex)
	} else {
	App:
		for {
			fmt.Print("Enter an example number from 1 to 2 or q to quit: ")
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
	val, _ := strconv.ParseInt(choice, 0, 64)
	num := fmt.Sprintf("%03d", val)
	if choice == "q" {
		fmt.Println("Bye")
		return false
	} else if val >= 1 && val <= 32 {
		fmt.Printf("Running example%v:\n", num)
		reflect.ValueOf(ex).MethodByName("Example" + num).Call(nil)
	} else {
		fmt.Println("Incorrect Entry. Bye!")
		return false
	}
	return true
}

type Examples struct{}

func (ex Examples) Example002() {
	date := stuff.Date{}
	err := date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(1974)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("1st Day: %d/%d/%d\n", date.Day(), date.Month(), date.Year())
}
func (ex Examples) Example001() {
	fmt.Println("Hello", stuff.Name)
	intArr := []int{2, 3, 5, 7, 11}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))
}
