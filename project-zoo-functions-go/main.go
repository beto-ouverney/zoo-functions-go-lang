package main

import (
	"fmt"
	"os"
	"project-zoo-functions-go/question1"
	"project-zoo-functions-go/question10"
	"project-zoo-functions-go/question2"
	"project-zoo-functions-go/question3"
	"project-zoo-functions-go/question4"
	"project-zoo-functions-go/question5"
	"project-zoo-functions-go/question6"
	"project-zoo-functions-go/question7"
	"project-zoo-functions-go/question8"
	"project-zoo-functions-go/question9"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("No question argument")
	} else {
		switch args[1] {
		case "q1":
			fmt.Printf("%+v\n", question1.GetSpeciesByIds(args[2:]))
		case "q2":
			intVar, err := strconv.Atoi(args[3])
			if err != nil {
				fmt.Print(err)
			}
			fmt.Printf("%+v", question2.GetAnimalsOlderThan(args[2], intVar))
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
		case "q6":
			fmt.Printf("%+v", question6.CalculateEntry())
		case "q7":
			if len(args) == 5 {
				options := question7.Options{}
				options.IncludeNames, _ = strconv.ParseBool(args[2])
				options.Sex = args[3]
				options.Sorted, _ = strconv.ParseBool(args[4])
				fmt.Printf("%+v", question7.GetAnimalMap(options))
			} else if len(args) == 4 {
				options := question7.Options{}
				options.IncludeNames, _ = strconv.ParseBool(args[2])
				if args[3] == "true" || args[3] == "false" {
					options.Sorted, _ = strconv.ParseBool(args[3])
				} else {
					options.Sex = args[3]
				}
				fmt.Printf("%+v", question7.GetAnimalMap(options))
			} else if len(args) == 3 {
				options := question7.Options{}
				options.IncludeNames, _ = strconv.ParseBool(args[2])
				fmt.Printf("%+v", question7.GetAnimalMap(options))
			} else if len(args) == 2 {
				fmt.Printf("%v", question7.GetAnimalMap(question7.Options{}))
			} else {
				fmt.Println("No question argument")
			}
		case "q8":
			if len(args) > 2 {
				fmt.Printf("%v", question8.GetSchedule(args[2]))
			} else {
				fmt.Printf("%v", question8.GetSchedule(""))
			}
		case "q9":
			if len(args) > 2 {
				fmt.Printf("%v", question9.GetOldestFromFirstSpecies(args[2]))
			} else {
				fmt.Printf("%v", question8.GetSchedule(""))
			}
		case "q10":
			if len(args) > 2 {
				fmt.Printf("%v", question10.GetEmployeesCoverage(args[2]))
			} else {
				fmt.Printf("%v", question10.GetEmployeesCoverage(""))
			}
		}
	}
}
