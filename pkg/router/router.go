package router

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"upper.io/db.v3/lib/sqlbuilder"
)

var rs = [](func(sqlbuilder.Database, *mux.Router)){
	routes.
}

func NewRouter(sess sqlbuilder.Database) http.Handler {
	router := mux.NewRouter()
	for _, r := range rs {
		r(sess, router)
	}

	// Add middleware for CORS policy
	// TODO update this
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{},
		AllowedMethods: []string{},
	})

	return middleware.Handler(router)
}