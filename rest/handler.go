package rest

import (
	"encoding/json"
	"fmt"
	"github.com/tahcohcat/combat-engine/entities"
	"log"
	"net/http"
)

func Get_Index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hello world</h1>`)
}

func Game_Handler_POST(w http.ResponseWriter, r *http.Request) {

	log.Print(r.Body)

	decoder := json.NewDecoder(r.Body)

	var game entities.Game

	err := decoder.Decode(&game)
	if err != nil {
		panic(err)
	}

	log.Println(game)

	w.Header().Set("Content-Type", "application/json")

	responseBody := make(map[string]string)
	responseBody["message"] = "Game created successfully"
	responseBody["id"] = string(game.Id)
	json.NewEncoder(w).Encode(responseBody)
	w.WriteHeader(http.StatusOK)
}

func Game_Handler_GET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hello world</h1>`)

}

