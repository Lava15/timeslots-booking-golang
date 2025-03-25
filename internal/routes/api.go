package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lava15/timeslots-booking-golang/internal/handlers"
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

	//Services endpoints
	serviceRouter := v1Router.PathPrefix("/services").Subrouter()
	serviceRouter.HandleFunc("", handlers.GetServices).Methods(http.MethodGet)
}

func NewRouter() *Router {
	r := &Router{mux.NewRouter()}
	r.apiRoutes()
	return r
}
