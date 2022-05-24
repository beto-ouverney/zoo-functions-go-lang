package question3

import (
	"project-zoo-functions-go/getData"
	"strings"
)

func GetEmployeeByName(args []string) []getData.Employee {

	data := getData.GetZooData()
	var result = make([]getData.Employee, 0, len(data.Employees))
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
