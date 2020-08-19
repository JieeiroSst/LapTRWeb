package admin

import (
	"net/http"
	"strconv"

	model "github.com/JieeiroSst/LapTRWeb/models/admin"
)

func ShowListBook(w http.ResponseWriter, r *http.Request) {
	res := model.ShowListBook()
	tmpl.ExecuteTemplate(w, "Book-List", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ShowSingleBook(w http.ResponseWriter, r *http.Request) {
	nId, _ := strconv.Atoi(r.URL.Query().Get("ISBN"))
	res := model.ShowSingleBook(nId)
	tmpl.ExecuteTemplate(w, "Book-Show", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func EditBook(w http.ResponseWriter, r *http.Request) {
	nId, _ := strconv.Atoi(r.URL.Query().Get("ISBN"))
	res := model.ShowSingleBook(nId)
	tmpl.ExecuteTemplate(w, "Book-Edit", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	nId, _ := strconv.Atoi(r.URL.Query().Get("ISBN"))
	model.DeleteBook(nId)
	http.Redirect(w, r, "/admin/book", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Book-New", nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func InsertBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("ISBN"))
		book := model.Book{
			ISBN:      id,
			Title:     r.FormValue("Title"),
			Language:  r.FormValue("Language"),
			MRP:       r.FormValue("MRP"),
			Publisher: r.FormValue("Publisher"),
			PubDate:   r.FormValue("PubDate"),
		}
		model.InsertBook(book)
	}
	http.Redirect(w, r, "/admin/book", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("ISBN"))
		book := model.Book{
			ISBN:      id,
			Title:     r.FormValue("Title"),
			Language:  r.FormValue("Language"),
			MRP:       r.FormValue("MRP"),
			Publisher: r.FormValue("Publisher"),
			PubDate:   r.FormValue("PubDate"),
		}
		model.UpdateBook(book)
	}
	http.Redirect(w, r, "/admin/book", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
