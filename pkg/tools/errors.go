package tools

import (
	"cars/pkg/datatypes"
	"fmt"
	"html/template"
	"net/http"
)

func HandleError(w http.ResponseWriter, status int) {
	var ErrMsg datatypes.ErrMsg

	ErrMsg.Code = status
	w.WriteHeader(status)
	switch status {
	case http.StatusNotFound:
		ErrMsg.Message = "Page not found"
	case http.StatusMethodNotAllowed:
		ErrMsg.Message = "Method not supported"
	case http.StatusBadRequest:
		ErrMsg.Message = "Bad Request"
	case http.StatusInternalServerError:
		ErrMsg.Message = "Internal server error"
	}

	tmpl, err := template.ParseFiles("./tmpl/error.html")
	if err != nil {
		fmt.Println("Error parsing the html:", err)
		http.Error(w, "Real internal server error this time", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, ErrMsg)
	if err != nil {
		fmt.Println("Error executing the template:", err)
		http.Error(w, "Real internal server error this time", http.StatusInternalServerError)
		return
	}
}

func ApiErrorFound(errCh <-chan int) bool {
	for errCode := range errCh {
		if errCode != 0 {
			return true
		}
	}
	return false
}
