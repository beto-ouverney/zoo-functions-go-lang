package question9

import (
	"project-zoo-functions-go/getData"
	"strconv"
)

func GetOldestFromFirstSpecies(id string) []string {
	data := getData.GetZooData()
	firstSpecie := ""
	result := []string{}
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
			result = []string{residents[0].Name, residents[0].Sex, strconv.Itoa(residents[0].Age)}
			break
		}
	}
	return result
}
