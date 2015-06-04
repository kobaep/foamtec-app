package inspectionReport

import (
	"html/template"
	"net/http"
)

func PageInspectionReport(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/ir/inspectionReport.html")
	t.Execute(w, nil)
}
