package admin

import (
	"net/http"
	"strconv"

	model "github.com/JieeiroSst/LapTRWeb/models/admin"
)

func ShowListcourse(w http.ResponseWriter, r *http.Request) {
	res := model.ShowListCourse()
	tmpl.ExecuteTemplate(w, "List", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ShowSigleCourse(w http.ResponseWriter, r *http.Request) {
	nId, _ := strconv.Atoi(r.URL.Query().Get("CourseId"))
	res := model.ShowsingleCourse(nId)
	tmpl.ExecuteTemplate(w, "Show", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func EditCourse(w http.ResponseWriter, r *http.Request) {
	nId, _ := strconv.Atoi(r.URL.Query().Get("CourseId"))
	res := model.ShowsingleCourse(nId)
	tmpl.ExecuteTemplate(w, "Edit", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func InsertCourse(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nId, _ := strconv.Atoi(r.FormValue("CourseId"))
		course := model.Course{
			CourseId:   nId,
			CourseName: r.FormValue("CourseName"),
			Department: r.FormValue("Department"),
		}
		model.InsertCourse(course)
	}
	http.Redirect(w, r, "/admin/course", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nId, _ := strconv.Atoi(r.FormValue("CourseId"))
		course := model.Course{
			CourseId:   nId,
			CourseName: r.FormValue("CourseName"),
			Department: r.FormValue("Department"),
		}
		model.UpdateCourse(course)
	}
	http.Redirect(w, r, "/admin/course", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	nId, _ := strconv.Atoi(r.URL.Query().Get("CourseId"))
	model.DeleteCourse(nId)
	http.Redirect(w, r, "/admin/course", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
