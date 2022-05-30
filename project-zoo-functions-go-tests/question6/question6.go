package question6

import "project-zoo-functions-go/getdata"

//EntrantsType represents zoo type costumers
type EntrantsType struct {
	Child  int
	Adult  int
	Senior int
}

//CountEntrants a counter to zoo entrants
func CountEntrants(data []getdata.Entrant) EntrantsType {
	result := EntrantsType{0, 0, 0}
	return result
}

//CalculateEntry It should receive the visitors array and return an object with the count according to the following sort criteria: Persons under the age of 18 are classified as children; People aged 18 or over and under 50 are classified as adults; People aged 50 or over are classified as older (senior) people.
func CalculateEntry(entrants []getdata.Entrant) float64 {
	return 0.0
}
