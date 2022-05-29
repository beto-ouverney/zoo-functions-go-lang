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
	for _, v := range data {
		if v.Age < 18 {
			result.Child++
		} else if v.Age >= 18 && v.Age < 50 {
			result.Adult++
		} else {
			result.Senior++
		}
	}
	return result
}

//CalculateEntry It should receive the visitors array and return an object with the count according to the following sort criteria: Persons under the age of 18 are classified as children; People aged 18 or over and under 50 are classified as adults; People aged 50 or over are classified as older (senior) people.
func CalculateEntry(entrants []getdata.Entrant) float64 {

	data := getdata.GetZooData()
	result := 0.0
	if len(entrants) == 0 {
		return 0
	}
	entrantsTotal := CountEntrants(entrants)
	result += float64(entrantsTotal.Child) * data.Prices.Child
	result += float64(entrantsTotal.Adult) * data.Prices.Adult
	result += float64(entrantsTotal.Senior) * data.Prices.Senior

	return result

}
