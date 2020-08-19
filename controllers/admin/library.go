package admin

import (
	"net/http"
	"strconv"

	model "github.com/JieeiroSst/LapTRWeb/models/admin"
)

func ShowListLibrary(w http.ResponseWriter, r *http.Request) {
	res := model.ShowListLibrary()
	tpmlLibrary.ExecuteTemplate(w, "List", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ShowSingleLibrary(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(r.URL.Query().Get("Id"))
	res := model.ShowSingleLibrary(Id)
	tpmlLibrary.ExecuteTemplate(w, "Show", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func EditLibrary(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(r.URL.Query().Get("Id"))
	res := model.ShowSingleLibrary(Id)
	tpmlLibrary.ExecuteTemplate(w, "Edit", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CreateLibrary(w http.ResponseWriter, r *http.Request) {
	tpmlLibrary.ExecuteTemplate(w, "New", nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteLibrary(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(r.URL.Query().Get("Id"))
	model.DeleteLibrary(Id)
	http.Redirect(w, r, "/admin/library", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func InsertLibrary(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Id, _ := strconv.Atoi(r.FormValue("Id"))
		isbn, _ := strconv.Atoi(r.FormValue("ISBN"))
		l := model.Library{
			Id:   Id,
			ISBN: isbn,
		}
		model.InserLibrary(l)
	}
	http.Redirect(w, r, "/admin/library", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func UpdateLibrary(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Id, _ := strconv.Atoi(r.FormValue("Id"))
		isbn, _ := strconv.Atoi(r.FormValue("ISBN"))
		l := model.Library{
			Id:   Id,
			ISBN: isbn,
		}
		model.InserLibrary(l)
	}
	http.Redirect(w, r, "/admin/library", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
