package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router struct embeds the Gorilla Mux router
type Router struct {
	*mux.Router
}

func (r *Router) apiRoutes() {
	apiRouter := r.PathPrefix("/api").Subrouter()
	v1Router := apiRouter.PathPrefix("/v1").Subrouter()
	v1Router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}

func NewRouter() *Router {
	r := &Router{mux.NewRouter()}
	r.apiRoutes()
	return r
}
