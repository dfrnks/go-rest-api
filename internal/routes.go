package internal

import (
	"github.com/dfrnks/go-rest-api/internal/People"
	"github.com/gorilla/mux"
	"net/http"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)

	router.HandleFunc("/contato", People.GetAll).Methods("GET")
	router.HandleFunc("/contato", People.Create).Methods("POST")
	router.HandleFunc("/contato/{id}", People.Get).Methods("GET")
	router.HandleFunc("/contato/{id}", People.Delete).Methods("DELETE")

	return router
}
