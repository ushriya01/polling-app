package handlers

import (
	"net/http"
	"poll-app/internal/models"

	"github.com/julienschmidt/httprouter"
)

func ClearDataHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := models.ClearAllData(r.Context())
	if err != nil {
		http.Error(w, "Failed to clear data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All data cleared successfully"))
}
