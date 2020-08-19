package admin

import (
	"html/template"
	"net/http"
)

var tmplAuth = template.Must(template.ParseGlob("views/admin/author/*"))
var tmpBook = template.Must(template.ParseGlob("views/admin/book/*"))
var tmpBuySell = template.Must(template.ParseGlob("views/admin/buy-sell/*"))
var tmplCourse = template.Must(template.ParseGlob("views/admin/course/*"))
var tpmlLibrary = template.Must(template.ParseGlob("views/admin/library/*"))
var tmplProfile = template.Must(template.ParseGlob("views/admin/profile/*"))
var tmplSeller = template.Must(template.ParseGlob("views/admin/seller/*"))

var tmpl = template.Must(template.ParseFiles("views/admin/Home.tmpl"))

func HomeAdmin(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
