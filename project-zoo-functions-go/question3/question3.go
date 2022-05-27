package question3

import (
	"project-zoo-functions-go/getdata"
	"strings"
)

//GetEmployeeByName is responsible for searching for collaborating people by their first or last name
func GetEmployeeByName(args []string) []getdata.Employee {

	data := getdata.GetZooData()
	var result = make([]getdata.Employee, 0, len(data.Employees))
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
