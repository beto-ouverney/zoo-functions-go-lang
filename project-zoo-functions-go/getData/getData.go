package getData

import (
	"encoding/json"
	"fmt"
	"os"
)

type species struct {
	Species []struct {
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
	} `json:"species"`
	Employees []struct {
		ID             string   `json:"id"`
		FirstName      string   `json:"firstName"`
		LastName       string   `json:"lastName"`
		Managers       []string `json:"managers"`
		ResponsibleFor []string `json:"responsibleFor"`
	} `json:"employees"`
	Hours struct {
		Tuesday struct {
			Open  int `json:"open"`
			Close int `json:"close"`
		} `json:"Tuesday"`
		Wednesday struct {
			Open  int `json:"open"`
			Close int `json:"close"`
		} `json:"Wednesday"`
		Thursday struct {
			Open  int `json:"open"`
			Close int `json:"close"`
		} `json:"Thursday"`
		Friday struct {
			Open  int `json:"open"`
			Close int `json:"close"`
		} `json:"Friday"`
		Saturday struct {
			Open  int `json:"open"`
			Close int `json:"close"`
		} `json:"Saturday"`
		Sunday struct {
			Open  int `json:"open"`
			Close int `json:"close"`
		} `json:"Sunday"`
		Monday struct {
			Open  int `json:"open"`
			Close int `json:"close"`
		} `json:"Monday"`
	} `json:"hours"`
	Prices struct {
		Adult  float64 `json:"adult"`
		Senior float64 `json:"senior"`
		Child  float64 `json:"child"`
	} `json:"prices"`
}

func ReturnStruct() species {
	jsonFile, err := os.ReadFile("./data/output.json")
	if err != nil {
		fmt.Print(err)
	}
	var jsonData species
	err = json.Unmarshal(jsonFile, &jsonData)
	if err != nil {
		fmt.Print(err)
	}
	return jsonData
}
