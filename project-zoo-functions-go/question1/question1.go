package question1

import "project-zoo-functions-go/getData"

type specie struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Popularity   int      `json:"popularity"`
	Location     string   `json:"location"`
	Availability []string `json:"availability"`
	Residents    []struct {
		Name string `json:"name"`
		Sex  string `json:"sex"`
		Age  int    `json:"age"`
	} `json:"residents"`
}

func ResultQuestion1(args []string) []specie {

	result := getData.GetZooData()

	var ids []string
	ids = append(ids, args...)

	var results = make([]specie, 0, len(result.Species))

	for _, vId := range ids {
		for _, vS := range result.Species {
			if vId == vS.ID {
				results = append(results, vS)
			}
		}
	}
	return results
}
