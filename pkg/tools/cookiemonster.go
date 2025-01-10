//COOOOOOKIIIIEEES!!!

package tools

import (
	"net/http"
	"strings"
)

// HandleCookies get the cookie value (comma separated string) containing 5 IDs of last visited cars from r,
// adds the current model ID to the beginning of the string and deletes last ID from it, then sets the new cookie value using w
func HandleCookies(id string, w http.ResponseWriter, r *http.Request) {

	//get cookie data
	LastVisited := GetCookieData(r)
	//add current id to the beginning of the cookie
	updatedCookieData := updateCookieData(id, LastVisited)
	//set new cookie
	setCookieData(w, updatedCookieData)

}

func GetCookieData(r *http.Request) string {

	cookieData, err := r.Cookie("lastVisited")
	if err != nil {
		return ""
	}
	return cookieData.Value

}

func updateCookieData(id string, lastVisited string) string {

	IDs := strings.Split(lastVisited, ",")
	for i, value := range IDs {
		if value == id {
			IDs = append(IDs[:i], IDs[i+1:]...)
			break
		}
	}
	IDs = append([]string{id}, IDs...)
	if len(IDs) > 5 {
		IDs = IDs[:5]
	}
	newData := strings.Join(IDs, ",")
	return strings.TrimRight(newData, ",")
}

func setCookieData(w http.ResponseWriter, cookieData string) {

	cookie := http.Cookie{
		Name:   "lastVisited",
		Value:  cookieData,
		MaxAge: 31536000,
	}
	http.SetCookie(w, &cookie)
}
