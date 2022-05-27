package main

import (
	"project-zoo-functions-go/getdata"
	"project-zoo-functions-go/question1"
	"project-zoo-functions-go/question2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpeciesByIds(t *testing.T) {

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
	actual := question2.GetAnimalsOlderThan("lions", 10)
	assert.Equal(t, actual, true)

	actual = question2.GetAnimalsOlderThan("penguins", 10)
	assert.Equal(t, actual, true)
}
