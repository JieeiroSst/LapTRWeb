package admin

import (
	"net/http"
	"strconv"

	model "github.com/JieeiroSst/LapTRWeb/models/admin"
)

func ShowListBuySell(w http.ResponseWriter, r *http.Request) {
	res := model.ShowListBuySell()
	tmpl.ExecuteTemplate(w, "list", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ShowSingleSell(w http.ResponseWriter, r *http.Request) {
	nId, _ := strconv.Atoi(r.URL.Query().Get("SellerId"))
	res := model.ShowSigleBuySell(nId)
	tmpl.ExecuteTemplate(w, "Show", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func EditBuySell(w http.ResponseWriter, r *http.Request) {
	nId, _ := strconv.Atoi(r.URL.Query().Get("SellerId"))
	res := model.ShowSigleBuySell(nId)
	tmpl.ExecuteTemplate(w, "Edit", res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CreateBuysell(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func InsertBuySell(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nId, _ := strconv.Atoi(r.FormValue("SellerId"))
		buyid, _ := strconv.Atoi(r.FormValue("BuyerId"))
		stype, _ := strconv.Atoi(r.FormValue("Stype"))
		isbn, _ := strconv.Atoi(r.FormValue("ISBN"))
		tid, _ := strconv.Atoi(r.FormValue("Tid"))
		buy := model.BuySellRecord{
			SellerId: nId,
			BuyerId:  buyid,
			Stype:    stype,
			ISBN:     isbn,
			Tid:      tid,
		}

		model.InsertBuySell(buy)
	}
	http.Redirect(w, r, "/admin/buy-sell", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func deleteBuysell(w http.ResponseWriter, r *http.Request) {
	nId, _ := strconv.Atoi(r.FormValue("SellerId"))
	model.DeleteBuySell(nId)
	http.Redirect(w, r, "/admin/buy-sell", 301)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
