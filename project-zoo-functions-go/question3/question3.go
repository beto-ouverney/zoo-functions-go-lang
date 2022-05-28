package question3

import (
	"project-zoo-functions-go/getdata"
	"strings"
)

//GetEmployeeByName is responsible for searching for collaborating people by their first or last name
func GetEmployeeByName(args string) getdata.Employee {
	data := getdata.GetZooData()
	result := getdata.Employee{}
	if args == "" {
		return result
	}
	for _, v := range data.Employees {
		name := v.FirstName + v.LastName
		if strings.Contains(strings.ToLower(name), strings.ToLower(args)) {
			result = v
		}
	}

	return result
}
