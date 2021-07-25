package middlew

import (
	"net/http"

	"github.com/efrainmg90/twittor/bd"
)

func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(rw, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
