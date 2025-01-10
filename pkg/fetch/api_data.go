package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

//ApiRequest gets the data from API based on provided endpoint, and unmarshal it to the struct
func ApiRequest(endpoint string, target any, wg *sync.WaitGroup, errCh chan int) {
	defer wg.Done()

	apiURL := "http://localhost:3000/api/"
	
	

	response, err := http.Get(apiURL + endpoint)
	
	if err != nil {
		fmt.Printf("API:ERROR while GETting the response: %s\n", err.Error())
		if strings.Contains(err.Error(), "connect: connection refused"){
			errCh <- http.StatusInternalServerError
		} else {
			errCh <- http.StatusBadRequest
		}		
		return 
	}
	
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("API:ERROR while READing the response: %s\n", err.Error())
		errCh <- http.StatusBadRequest
		return
	}

	err = json.Unmarshal(responseData, &target)
	if err != nil {
		fmt.Printf("API:ERROR while UNMARSHALing the response: %s\n", err.Error())
		errCh <- http.StatusBadRequest
		return 
	}


}

