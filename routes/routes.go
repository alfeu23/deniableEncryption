package routes

import (
	"net/http"

	"deniableEncryption/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/vote", handlers.VoteHandler).Methods("POST")

	router.HandleFunc("/receipt/{id}", handlers.ReceiptHandler).Methods("GET")
	router.HandleFunc("/verify/{id}", handlers.VerifyReceiptHandler).Methods("GET")

	router.HandleFunc("/election/results", handlers.ElectionResultsHandler).Methods("GET")

	router.HandleFunc("/admin/flush", handlers.FlushCacheHandler).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return router
}
