package main

import (
	"log"

	"github.com/efrainmg90/twittor/bd"
	"github.com/efrainmg90/twittor/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	}
	handlers.Manejadores()
}
