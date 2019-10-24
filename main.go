package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"./controllers"
	m "./models"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {
	m.InitTables()

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},                                               // All origins
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"}, // Allowing only get, just an example
	})

	rand.Seed(time.Now().UnixNano())
	router.HandleFunc("/api/traits", controllers.CreateClassTrait).Methods("POST")
	router.HandleFunc("/api/traits", controllers.GetClassTraits).Methods("GET")

	router.HandleFunc("/api/class", controllers.CreateCharacterClass).Methods("POST")
	router.HandleFunc("/api/class", controllers.GetCharacterClasses).Methods("GET")

	router.HandleFunc("/api/attribute", controllers.CreateCharacterAttribute).Methods("POST")
	router.HandleFunc("/api/attribute", controllers.GetCharacterAttributes).Methods("GET")

	router.HandleFunc("/api/skills", controllers.CreateCharacterSkill).Methods("POST")
	router.HandleFunc("/api/skills", controllers.GetCharacterSkills).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("client/public").HTTPBox()))

	port := "9000"
	if port == "" {
		port = "8000"
	}

	// err2 := open.Run("http://localhost:9000")
	// if err2 != nil {
	// 	log.Println(err2)
	// }

	err := http.ListenAndServe(":"+port, c.Handler(router))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("server running")
	}

}
