package main

import (
	"fmt"
	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// Charger les configuration contenu dans congig.json en utilisant viper
	LoadAppConfig()

	// Initialisation de la Base de données
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialisation du routeur
	router := mux.NewRouter().StrictSlash(true)

	// Enregistrer le routeur
	RegisterProductRoutes(router)

	// Démarrer le serveur
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}
