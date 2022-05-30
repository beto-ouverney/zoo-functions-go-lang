package question9

import (
	"project-zoo-functions-go/getdata"
)

//Response represents the response os the function GetOldestFromFierstSpecies
type Response struct {
	Name string
	Sex  string
	Age  int
}

//GetOldestFromFirstSpecies searches for information on the oldest animal of the first species managed by the person collaborating with the parameter.
func GetOldestFromFirstSpecies(id string) Response {
	data := getdata.GetZooData()
	firstSpecie := ""
	result := Response{}
	for _, v := range data.Employees {
		if v.ID == id {
			firstSpecie = v.ResponsibleFor[0]
			break
		}
	}
	for _, v := range data.Species {
		if v.ID == firstSpecie {
			residents := v.Residents
			for i := 1; i < len(residents); i++ {
				if residents[0].Age < residents[i].Age {
					residents[0] = residents[i]
				}
			}
			return Response{residents[0].Name, residents[0].Sex, residents[0].Age}
		}
	}
	return result
}
