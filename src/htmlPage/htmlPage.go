package htmlPage

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"html/template"
	"net/http"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func PageInspectionReportPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inspectionReport page")
	userName := getUserName(r)
	if userName != "" {
		t, _ := template.ParseFiles("templates/ir/inspectionReport.html")
	} else {
		t, _ := template.ParseFiles("templates/login.html")
	}
	t.Execute(w, r)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/login.html")
	t.Execute(w, nil)
}
