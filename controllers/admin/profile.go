package admin

import (
	"net/http"
	"strconv"

	models "github.com/JieeiroSst/LapTRWeb/models/admin"
)

func ShowListProfile(w http.ResponseWriter, r *http.Request) {
	res := models.ShowListProfile()
	tmplProfile.ExecuteTemplate(w, "List", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ShowSingleProfile(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(r.URL.Query().Get("UserId"))
	res := models.ShowSingleProfile(Id)
	tmplProfile.ExecuteTemplate(w, "Show", res)
}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(r.URL.Query().Get("UserId"))
	res := models.ShowSingleProfile(Id)
	tmplProfile.ExecuteTemplate(w, "Edit", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(r.URL.Query().Get("UserId"))
	models.DeleteProfile(Id)
	http.Redirect(w, r, "/admin/profile", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func UpdateProdile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Id, _ := strconv.Atoi(r.FormValue("UserId"))
		p := models.Profile{
			UserID:   Id,
			UserName: r.FormValue("UserName"),
			Address:  r.FormValue("Address"),
			Password: r.FormValue("Password"),
		}
		models.UpdateProfile(p)
	}
	http.Redirect(w, r, "/admin/profile", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
