package main

import (
	"fmt"
	"os"
	"project-zoo-functions-go/question1"
	"project-zoo-functions-go/question2"
	"project-zoo-functions-go/question3"
	"project-zoo-functions-go/question4"
	"project-zoo-functions-go/question5"
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
		case "q4":
			fmt.Printf("%+v", question4.GetRelatedEmployees(args[2]))
		case "q5":
			if len(args) > 3 {
				param := question5.ParamTypeQ5{
					Specie: args[2],
					Sex:    args[3]}
				fmt.Printf("%+v", question5.CountAnimals(param))
			} else if len(args) > 2 {
				param := question5.ParamTypeQ5{
					Specie: args[2]}
				fmt.Printf("%+v", question5.CountAnimals(param))
			} else {
				fmt.Println("No question argument")
			}
		}
	}
}
