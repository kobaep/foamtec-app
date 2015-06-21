package main

import (
	"htmlPage"
	"net/http"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", makeHandler(htmlPage.IndexPage))
	http.HandleFunc("/login", makeHandler(htmlPage.LoginPage))
	http.HandleFunc("/ir", makeHandler(htmlPage.PageInspectionReportPage))

	http.ListenAndServe(":8080", nil)
}
