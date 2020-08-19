package admin

import (
	"net/http"
	"strconv"

	"github.com/JieeiroSst/LapTRWeb/models/admin"
)

func IndexAuthor(w http.ResponseWriter, r *http.Request) {
	res := admin.ShowListAuthor()
	tmpl.ExecuteTemplate(w, "List", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func IndexSignleauthor(w http.ResponseWriter, r *http.Request) {
	nId, _ := strconv.Atoi(r.URL.Query().Get("AuthId"))
	res := admin.ShowSingleAuthor(nId)
	tmpl.ExecuteTemplate(w, "Show", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func NewAuthor(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func EditAuthor(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(r.URL.Query().Get("AuthId"))
	res := admin.ShowSingleAuthor(Id)
	tmpl.ExecuteTemplate(w, "Edit", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteAuth(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(r.URL.Query().Get("AuthId"))
	admin.DeleteAuth(Id)
	http.Redirect(w, r, "/admin/auth", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CreateAuth(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func InsertAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("AuthId"))
		author := admin.Author{
			AuthId:     id,
			Name:       r.FormValue("Name"),
			Affliation: r.FormValue("Affliation"),
			Email:      r.FormValue("Email"),
		}
		admin.InsertAuthor(author)
	}
	http.Redirect(w, r, "/admin/auth", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func UpdateAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("AuthId"))
		author := admin.Author{
			AuthId:     id,
			Name:       r.FormValue("Name"),
			Affliation: r.FormValue("Affliation"),
			Email:      r.FormValue("Email"),
		}
		admin.UpdateAuthor(author)
	}
	http.Redirect(w, r, "/admin/auth", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
