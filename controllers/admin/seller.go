package admin

import (
	"net/http"
	"strconv"

	model "github.com/JieeiroSst/LapTRWeb/models/admin"
)

func ShowListSeller(w http.ResponseWriter, r *http.Request) {
	res := model.ShowListSell()
	tmpl.ExecuteTemplate(w, "Seller-List", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ShowSingleSeller(w http.ResponseWriter, r *http.Request) {
	res := model.ShowListSell()
	tmpl.ExecuteTemplate(w, "Seller-Show", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteSell(w http.ResponseWriter, r *http.Request) {
	Id, _ := strconv.Atoi(r.URL.Query().Get("SellerId"))
	model.DeleteSeller(Id)
	http.Redirect(w, r, "/admin/selle", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
