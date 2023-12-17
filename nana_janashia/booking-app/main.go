package main

import (
	"booking-app/arrays"
	"booking-app/concurrents"
	"booking-app/maps"
	"booking-app/structs"
	"fmt"
)

func main() {
	var choice string
outer:
	for {
		fmt.Printf("Please enter the example number to run.\n1)Arrays 2)Maps 3)Structs 4)Concurrents q)Quit: ")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			arrays.Main()
		case "2":
			maps.Main()
		case "3":
			structs.Main()
		case "4":
			concurrents.Main()
		default:
			if choice == "q" {
				fmt.Println("Bye!")
			} else {
				fmt.Println("Incorrect Entry! Bye!")
			}
			break outer
		}
	}
}
