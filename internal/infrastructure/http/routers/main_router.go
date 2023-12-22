package routers

import "github.com/gorilla/mux"

// MainRouter returns a single router with all routes merged
func MainRouter(routers ...*mux.Router) *mux.Router {
	mr := mux.NewRouter()

	for _, r := range routers {
		mr.PathPrefix("/").Handler(r)
	}

	return mr
}
