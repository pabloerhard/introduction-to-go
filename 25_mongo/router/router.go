package router

import (
	"github.com/gorilla/mux"
	"mongo/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.CreateOneMovie).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateOneMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies/{id}", controller.DeleteAllMovies).Methods("DELETE")
	return router
}
