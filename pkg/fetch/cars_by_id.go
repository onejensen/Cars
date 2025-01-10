package fetch

import (
	"cars/pkg/datatypes"
	"cars/pkg/tools"
	"net/http"
	"sync"
)
//GetModelsByIDs Concurrently gets data for all the models in IDs, puts in in the models, while getting the names for cart type and manufacturer from categories and manufacturers. 
func GetModelsByIDs(IDs []string, models []datatypes.Model, manufacturers []datatypes.Manufacturer, categories []datatypes.Category, w http.ResponseWriter) []datatypes.Model {
	var selectedModels []datatypes.Model
	var wg sync.WaitGroup

	idCount := len(IDs)
	errCh := make(chan int, idCount)

	//check if IDs are empty, and return an empty slice
	if idCount == 1 && IDs[0] == ""{
		return selectedModels
	}
	//populate the resulting slice in advance, to avoid out of bounds error
	for i := 0; i < idCount; i++ {
		selectedModels = append(selectedModels, datatypes.Model{})
	}
	//concurrently get all the cars data from API
	wg.Add(idCount)
	for i, id := range IDs {
		go ApiRequest("models/"+id, &selectedModels[i], &wg, errCh)
	}
	wg.Wait()
	close(errCh)
	if tools.ApiErrorFound(errCh) {
		tools.HandleError(w, <-errCh)
		return selectedModels
	}

	//remove a structs with Model.ID 0, and get the Manufacturer and Category name for the rest of the models
	for i := idCount - 1; i >= 0; i-- {
		if selectedModels[i].ID == 0 {
			selectedModels = append(selectedModels[:i], selectedModels[i+1:]...)
		} else {
			selectedModels[i].Category = categories[selectedModels[i].CategoryID-1].Name
			selectedModels[i].Manufacturer = manufacturers[selectedModels[i].ManufacturerID-1].Name
		}
	}
	return selectedModels
}
