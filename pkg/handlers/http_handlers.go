package handlers

import (
	"cars/pkg/datatypes"
	"cars/pkg/fetch"
	"cars/pkg/tools"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"text/template"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	//check the url and method
	if r.URL.Path != "/" {
		tools.HandleError(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		tools.HandleError(w, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	//create template
	tmpl, err := template.ParseFiles("./tmpl/index.html")
	if err != nil {
		tools.HandleError(w, http.StatusInternalServerError)
		fmt.Println("Error parsing the html:", err)
		return
	}
	//make concurrent API requests
	var models []datatypes.Model
	var manufacturers []datatypes.Manufacturer
	var categories []datatypes.Category

	var wg sync.WaitGroup
	errCh := make(chan int, 3)

	wg.Add(3)

	go fetch.ApiRequest("models", &models, &wg, errCh)
	go fetch.ApiRequest("manufacturers", &manufacturers, &wg, errCh)
	go fetch.ApiRequest("categories", &categories, &wg, errCh)

	wg.Wait()

	close(errCh)
	if tools.ApiErrorFound(errCh) {
		tools.HandleError(w, <-errCh)
		return
	}

	//populate the page data
	var pageData datatypes.Data
	pageData.Path = "gallery"
	pageData.Categories = categories
	pageData.Manufacturers = manufacturers
	pageData.Models = models

	//execute the template
	err = tmpl.Execute(w, pageData)
	if err != nil {
		fmt.Println("Error executing the template:", err)
		tools.HandleError(w, http.StatusInternalServerError)
		return
	}
}

func HandleCar(w http.ResponseWriter, r *http.Request) {
	//check the url and method
	if r.URL.Path != "/car" {
		tools.HandleError(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		tools.HandleError(w, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	//parse the query from the URL and extract car id from it
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		tools.HandleError(w, http.StatusBadRequest)
		return
	}

	if len(query["id"]) == 0 {
		tools.HandleError(w, http.StatusBadRequest)
		return
	}
	modelID := query["id"][0]

	//create template
	tmpl, err := template.ParseFiles("./tmpl/car.html")
	if err != nil {
		fmt.Println("Error parsing the html:", err)
		tools.HandleError(w, http.StatusInternalServerError)
		return
	}
	//make concurrent API requests
	var model datatypes.Model
	var manufacturers []datatypes.Manufacturer
	var categories []datatypes.Category

	var wg sync.WaitGroup
	errCh := make(chan int, 3)
	wg.Add(3)
	go fetch.ApiRequest("models/"+modelID, &model, &wg, errCh)
	go fetch.ApiRequest("manufacturers", &manufacturers, &wg, errCh)
	go fetch.ApiRequest("categories", &categories, &wg, errCh)

	wg.Wait()

	close(errCh)
	if tools.ApiErrorFound(errCh) {
		tools.HandleError(w, <-errCh)
		return
	}

	if model.ID != 0 {
		model.Manufacturer = manufacturers[model.ManufacturerID-1].Name
		model.Country = manufacturers[model.ManufacturerID-1].Country
		model.FoundingYear = manufacturers[model.ManufacturerID-1].FoundingYear
		model.Category = categories[model.CategoryID-1].Name
	}
	//handle cookies to track visited pages
	tools.HandleCookies(modelID, w, r)

	//execute the template
	err = tmpl.Execute(w, model)
	if err != nil {
		fmt.Println("Error executing the template:", err)
		tools.HandleError(w, http.StatusInternalServerError)
		return
	}

}

func HandleFilter(w http.ResponseWriter, r *http.Request) {
	//check the url and method
	if r.URL.Path != "/filter" {
		tools.HandleError(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		tools.HandleError(w, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	//parse the query from the URL
	query, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil {
		tools.HandleError(w, http.StatusBadRequest)
		return
	}

	//create template
	tmpl, err := template.ParseFiles("./tmpl/index.html")
	if err != nil {
		fmt.Println("Error parsing the html:", err)
		tools.HandleError(w, http.StatusInternalServerError)
		return
	}
	//make API requests

	var models []datatypes.Model
	var manufacturers []datatypes.Manufacturer
	var categories []datatypes.Category

	var wg sync.WaitGroup
	errCh := make(chan int, 3)

	wg.Add(3)

	go fetch.ApiRequest("models", &models, &wg, errCh)
	go fetch.ApiRequest("manufacturers", &manufacturers, &wg, errCh)
	go fetch.ApiRequest("categories", &categories, &wg, errCh)

	wg.Wait()
	close(errCh)
	if tools.ApiErrorFound(errCh) {
		tools.HandleError(w, <-errCh)
		return
	}
	//sort the models based on filter parameters and populate the page data
	var pageData datatypes.Data
	pageData.Path = "search"
	pageData.Categories = categories
	pageData.Manufacturers = manufacturers
	pageData.Models = tools.Filter(query, models)

	//execute the template
	err = tmpl.Execute(w, pageData)
	if err != nil {
		fmt.Println("Error executing the template:", err)
		tools.HandleError(w, http.StatusInternalServerError)
		return
	}

}

func HandleCompare(w http.ResponseWriter, r *http.Request) {

	//check the url and method
	if r.URL.Path != "/compare" {
		tools.HandleError(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		tools.HandleError(w, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	//parse the query from the URL
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		tools.HandleError(w, http.StatusBadRequest)
		return
	}
	//create template
	tmpl, err := template.ParseFiles("./tmpl/index.html")
	if err != nil {
		fmt.Println("Error parsing the html:", err)
		tools.HandleError(w, http.StatusInternalServerError)
		return
	}
	//make API requests
	var models []datatypes.Model
	var manufacturers []datatypes.Manufacturer
	var categories []datatypes.Category

	var wg sync.WaitGroup
	errCh := make(chan int, 2)

	wg.Add(2)

	go fetch.ApiRequest("manufacturers", &manufacturers, &wg, errCh)
	go fetch.ApiRequest("categories", &categories, &wg, errCh)

	wg.Wait()
	close(errCh)
	if tools.ApiErrorFound(errCh) {
		tools.HandleError(w, <-errCh)
		return
	}

	var pageData datatypes.Data
	pageData.Path = "compare"
	pageData.Categories = categories
	pageData.Manufacturers = manufacturers
	pageData.Models = fetch.GetModelsByIDs(query["IDs"], models, manufacturers, categories, w)

	//execute the template
	err = tmpl.Execute(w, pageData)
	if err != nil {
		fmt.Println("Error executing the template:", err)
		tools.HandleError(w, http.StatusInternalServerError)
		return
	}

}

func HandleLastViewed(w http.ResponseWriter, r *http.Request) {

	//check the url and method
	if r.URL.Path != "/last" {
		tools.HandleError(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		tools.HandleError(w, http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	//get last visited cars from cookie
	IDs := strings.Split(tools.GetCookieData(r), ",")
	//create template
	tmpl, err := template.ParseFiles("./tmpl/index.html")
	if err != nil {
		fmt.Println("Error parsing the html:", err)
		tools.HandleError(w, http.StatusInternalServerError)
		return
	}
	//make API requests
	var models []datatypes.Model
	var manufacturers []datatypes.Manufacturer
	var categories []datatypes.Category

	var wg sync.WaitGroup
	errCh := make(chan int, 2)

	wg.Add(2)

	go fetch.ApiRequest("manufacturers", &manufacturers, &wg, errCh)
	go fetch.ApiRequest("categories", &categories, &wg, errCh)

	wg.Wait()
	close(errCh)
	if tools.ApiErrorFound(errCh) {
		tools.HandleError(w, <-errCh)
		return
	}

	var pageData datatypes.Data
	pageData.Path = "last_viewed"
	pageData.Models = fetch.GetModelsByIDs(IDs, models, manufacturers, categories, w)
	pageData.Manufacturers = manufacturers
	pageData.Categories = categories

	//execute the template
	err = tmpl.Execute(w, pageData)
	if err != nil {
		fmt.Println("Error executing the template:", err)
		tools.HandleError(w, http.StatusInternalServerError)
		return
	}
}
