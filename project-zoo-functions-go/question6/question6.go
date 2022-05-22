package question6

import "project-zoo-functions-go/getData"

type Entrant struct {
	Name string `json:"name"`
	Age  int
}

type EntrantsType struct {
	child  int
	adult  int
	senior int
}

func countEntrants() EntrantsType {
	result := EntrantsType{0, 0, 0}
	data := getData.GetEntrants()
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

func CalculateEntry() float64 {

	data := getData.GetZooData()
	result := 0.0
	entrantsTotal := countEntrants()
	result += float64(entrantsTotal.child) * data.Prices.Child
	result += float64(entrantsTotal.adult) * data.Prices.Adult
	result += float64(entrantsTotal.senior) * data.Prices.Senior

	return result

}
