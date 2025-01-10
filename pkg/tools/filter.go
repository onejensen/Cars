package tools

import (
	"cars/pkg/datatypes"
	"net/url"
	"strconv"
)

func Filter(query url.Values, models []datatypes.Model) []datatypes.Model {
	filteredResult := []datatypes.Model{}
	//map[categoryID:[4] manufacturerID:[7] maxYear:[2023] minYear:[2023]]
	for _, model := range models {
		if !matchManufacturer(model, query["manufacturerID"]) {
			continue
		}
		if !matchCategory(model, query["categoryID"]) {
			continue
		}
		if !matchYear(model, query["minYear"], query["maxYear"]) {
			continue
		}
		filteredResult = append(filteredResult, model)
	}

	return filteredResult
}

func matchManufacturer(model datatypes.Model, ManufacturerID []string) bool {
	//["7"]
	if ManufacturerID[0] == "" {
		return true
	}
	if ManufacturerID[0] == strconv.Itoa(model.ManufacturerID) {
		return true
	}
	return false
}

func matchCategory(model datatypes.Model, CategoryID []string) bool {
	if CategoryID[0] == "" {
		return true
	}
	if CategoryID[0] == strconv.Itoa(model.CategoryID) {
		return true
	}
	return false
}

func matchYear(model datatypes.Model, MinYearQuery, MaxYearQuery []string) bool {
	if MinYearQuery[0] == "" && MaxYearQuery[0] == "" {
		return true
	}
	minYear, err := strconv.Atoi(MinYearQuery[0])
	if err != nil || minYear < 0 {
		minYear = 0
	}
	maxYear, err := strconv.Atoi(MaxYearQuery[0])
	if err != nil || maxYear > 5000 {
		maxYear = 5000
	}

	return model.Year >= minYear && model.Year <= maxYear
}
