package question2

import (
	"fmt"
	"project-zoo-functions-go/getdata"
	"strconv"
)

//GetAnimalsOlderThan based on the name of a species and a minimum age, checks if all animals of that species have the specified minimum age.
func GetAnimalsOlderThan(animal string, ageParam string) bool {
	age, err := strconv.Atoi(ageParam)
	if err != nil {
		fmt.Print(err)
	}
	data := getdata.GetZooData()
	haveOlder := false

	for _, v := range data.Species {
		if v.Name == animal {
			for _, vR := range v.Residents {
				if vR.Age >= age {
					haveOlder = true
					break
				}
			}
		}
	}
	return haveOlder
}
