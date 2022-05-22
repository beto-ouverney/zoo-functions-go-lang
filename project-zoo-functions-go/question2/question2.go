package question2

import (
	"fmt"
	"project-zoo-functions-go/getData"
	"strconv"
)

func GetAnimalsOlderThan(animal string, ageParam string) bool {
	age, err := strconv.Atoi(ageParam)
	if err != nil {
		fmt.Print(err)
	}
	result := getData.ReturnStruct()
	haveOlder := false

	for _, v := range result.Species {
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
