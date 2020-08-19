package admin

import (
	"net/http"
)

func HomeAdmin(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Admin-Home", nil)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
