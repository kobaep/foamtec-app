package main

import (
	"inspectionReport"
	"net/http"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/ir", makeHandler(inspectionReport.PageInspectionReport))
	http.ListenAndServe(":8080", nil)
}
