package question2

import (
	"project-zoo-functions-go/getdata"
)

//GetAnimalsOlderThan based on the name of a species and a minimum age, checks if all animals of that species have the specified minimum age.
func GetAnimalsOlderThan(animal string, age int) bool {
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
