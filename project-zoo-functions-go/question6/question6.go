package question6

import "project-zoo-functions-go/getdata"

//EntrantsType represents zoo type costumers
type EntrantsType struct {
	child  int
	adult  int
	senior int
}

func countEntrants() EntrantsType {
	result := EntrantsType{0, 0, 0}
	data := getdata.GetEntrants()
	for _, v := range data.Entrants {
		if v.Age < 18 {
			result.child++
		} else if v.Age >= 18 && v.Age < 50 {
			result.adult++
		} else {
			result.senior++
		}
	}
	return result
}

//CalculateEntry It should receive the visitors array and return an object with the count according to the following sort criteria: Persons under the age of 18 are classified as children; People aged 18 or over and under 50 are classified as adults; People aged 50 or over are classified as older (senior) people.
func CalculateEntry() float64 {

	data := getdata.GetZooData()
	result := 0.0
	entrantsTotal := countEntrants()
	result += float64(entrantsTotal.child) * data.Prices.Child
	result += float64(entrantsTotal.adult) * data.Prices.Adult
	result += float64(entrantsTotal.senior) * data.Prices.Senior

	return result

}
