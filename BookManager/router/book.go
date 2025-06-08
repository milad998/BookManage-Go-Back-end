package router

import (
    "github.com/gorilla/mux"
    "BookManager/controller"
)

func Router() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/books", controller.CreateBook).Methods("POST")
    router.HandleFunc("/books", controller.GetBooks).Methods("GET")
    
    return router
}
