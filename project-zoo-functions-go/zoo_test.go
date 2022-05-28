package main

import (
	"project-zoo-functions-go/getdata"
	"project-zoo-functions-go/question1"
	"project-zoo-functions-go/question2"
	"project-zoo-functions-go/question3"
	"project-zoo-functions-go/question4"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpeciesByIDs(t *testing.T) {
	t.Log("Testing Question 1!")
	type test struct {
		data   []string
		answer []getdata.Specie
	}
	tests := []test{
		{
			data:   []string{""},
			answer: []getdata.Specie{},
		},
		{
			data: []string{
				"0938aa23-f153-4937-9f88-4858b24d6bce",
			},
			answer: []getdata.Specie{
				{
					ID:           "0938aa23-f153-4937-9f88-4858b24d6bce",
					Name:         "lions",
					Popularity:   4,
					Location:     "NE",
					Availability: []string{"Tuesday", "Thursday", "Saturday", "Sunday"},
					Residents: []getdata.Resident{
						{Name: "Zena", Sex: "female", Age: 12},
						{Name: "Maxwell", Sex: "male", Age: 15},
						{Name: "Faustino", Sex: "male", Age: 7},
						{Name: "Dee", Sex: "female", Age: 14},
					},
				},
			},
		}, {
			data: []string{"0938aa23-f153-4937-9f88-4858b24d6bce",
				"e8481c1d-42ea-4610-8e11-1752cfc05a46"},
			answer: []getdata.Specie{
				{
					ID:           "0938aa23-f153-4937-9f88-4858b24d6bce",
					Name:         "lions",
					Popularity:   4,
					Location:     "NE",
					Availability: []string{"Tuesday", "Thursday", "Saturday", "Sunday"},
					Residents: []getdata.Resident{
						{Name: "Zena", Sex: "female", Age: 12},
						{Name: "Maxwell", Sex: "male", Age: 15},
						{Name: "Faustino", Sex: "male", Age: 7},
						{Name: "Dee", Sex: "female", Age: 14},
					},
				},
				{
					ID:           "e8481c1d-42ea-4610-8e11-1752cfc05a46",
					Name:         "tigers",
					Popularity:   5,
					Location:     "NW",
					Availability: []string{"Wednesday", "Friday", "Saturday", "Tuesday"},
					Residents: []getdata.Resident{
						{Name: "Shu", Sex: "female", Age: 19},
						{Name: "Esther", Sex: "female", Age: 17},
					},
				},
			},
		},
	}

	for _, v := range tests {
		for i := 0; i < len(v.data); i++ {
			actual := question1.GetSpeciesByIds(v.data)
			assert.Equal(t, actual, v.answer)
		}
	}
}

func TestGetAnimalsOlderThan(t *testing.T) {
	t.Log("Testing Question 2!")
	actual := question2.GetAnimalsOlderThan("lions", 10)
	assert.Equal(t, actual, true)

	actual = question2.GetAnimalsOlderThan("penguins", 10)
	assert.Equal(t, actual, true)
}

func TestGetEmployeeByName(t *testing.T) {
	t.Log("Testing Question 3!")
	lionID := "0938aa23-f153-4937-9f88-4858b24d6bce"
	elephantsID := "bb2a76d8-5fe3-4d03-84b7-dba9cfc048b5"
	bearsID := "baa6e93a-f295-44e7-8f70-2bcdc6f6948d"
	snakesID := "78460a91-f4da-4dea-a469-86fd2b8ccc84"

	actual := question3.GetEmployeeByName("")
	assert.Equal(t, actual, getdata.Employee{})

	test := []struct {
		data   string
		answer getdata.Employee
	}{
		{
			data: "Emery",
			answer: getdata.Employee{
				ID:        "b0dc644a-5335-489b-8a2c-4e086c7819a2",
				FirstName: "Emery",
				LastName:  "Elser",
				Managers:  []string{"9e7d4524-363c-416a-8759-8aa7e50c0992"},
				ResponsibleFor: []string{
					lionID,
					bearsID,
					elephantsID,
				},
			},
		},
		{
			data: "Wishart",
			answer: getdata.Employee{
				ID:        "56d43ba3-a5a7-40f6-8dd7-cbb05082383f",
				FirstName: "Wilburn",
				LastName:  "Wishart",
				Managers: []string{
					"0e7b460e-acf4-4e17-bcb3-ee472265db83",
					"fdb2543b-5662-46a7-badc-93d960fdc0a8",
				},
				ResponsibleFor: []string{
					snakesID,
					elephantsID,
				},
			},
		},
	}

	for _, v := range test {
		actual := question3.GetEmployeeByName(v.data)
		t.Log(actual, v.answer)
		assert.Equal(t, actual, v.answer)
	}
}

func TestGetRelatedEmployees(t *testing.T) {
	t.Log("Testing Question 4!")
	const stephanieID = "9e7d4524-363c-416a-8759-8aa7e50c0992"

	actual := question4.GetRelatedEmployees("4b40a139-d4dc-4f09-822d-ec25e819a5ad")
	assert.Equal(t, actual, "The id entered is not that of a collaborating person manager!")

	actual = question4.GetRelatedEmployees(stephanieID)
	assert.Equal(t, actual, []string{"Burl Bethea", "Ola Orloff", "Emery Elser"})
}
