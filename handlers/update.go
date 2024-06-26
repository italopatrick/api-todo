package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/italopatrick/api-todo/models"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	} else {
		rows, err := models.Update(int64(id), todo)
		if err != nil {
			log.Printf("Erro ao atualizar registro: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if rows > 1 {
			log.Printf("Error: foram atualizados: %v", rows)
		}

		resp := map[string]any{
			"Error":   false,
			"Message": "Dados atualizados com sucesso!",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
