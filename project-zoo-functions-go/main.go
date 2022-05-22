package main

import (
	"fmt"
	"os"
	"project-zoo-functions-go/question1"
	"project-zoo-functions-go/question2"
	"project-zoo-functions-go/question3"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("No question argument")
	} else {

		switch args[1] {
		case "q1":
			fmt.Printf("%+v\n", question1.ResultQuestion1(args[2:]))
		case "q2":
			fmt.Printf("%+v", question2.GetAnimalsOlderThan(args[2], args[3]))
		case "q3":
			fmt.Printf("%+v", question3.GetEmployeeByName(args[2:]))
		}

	}
}
