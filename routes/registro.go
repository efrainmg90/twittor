package routes

import (
	"encoding/json"
	"net/http"

	"github.com/efrainmg90/twittor/bd"
	"github.com/efrainmg90/twittor/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in the body data!"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "The email is required!", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "The password should be least 6 characters!", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "The user already exists", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil || status == false {
		http.Error(w, "Error inserting the user"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
