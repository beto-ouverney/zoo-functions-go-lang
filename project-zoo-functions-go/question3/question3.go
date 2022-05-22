package question3

import (
	"project-zoo-functions-go/getData"
	"strings"
)

type Employee struct {
	ID             string   `json:"id"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Managers       []string `json:"managers"`
	ResponsibleFor []string `json:"responsibleFor"`
}

func GetEmployeeByName(args []string) []Employee {

	data := getData.GetZooData()
	var result = make([]Employee, 0, len(data.Employees))
	if len(args) == 1 {
		for _, v := range data.Employees {
			name := v.FirstName + v.LastName
			if strings.Contains(strings.ToLower(name), strings.ToLower(args[0])) {
				result = append(result, v)
			}
		}
	}
	return result
}
