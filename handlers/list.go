package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/italopatrick/api-todo/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
		http.Error(w, "Erro ao obter registros", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		log.Printf("Erro ao codificar dados em JSON: %v", err)
		http.Error(w, "Erro ao codificar dados em JSON", http.StatusInternalServerError)
		return
	}
}
