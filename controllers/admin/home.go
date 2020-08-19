package admin

import (
	"log"
	"net/http"
)

func HomeAdmin(w http.ResponseWriter, r *http.Request) {
	log.Fatal(tmpl.ExecuteTemplate(w, "Home", nil))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
