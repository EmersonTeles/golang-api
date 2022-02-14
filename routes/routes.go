package routes

import (
	"net/http"

	"github.com/EmersonTeles/golang-api/controllers"
	"github.com/EmersonTeles/golang-api/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalityById).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreateNewPersonality).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.UpdatePersonality).Methods("Put")
	http.ListenAndServe(":8000", r)
}
