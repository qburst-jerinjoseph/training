package app

import "github.com/go-chi/chi"

func (s *server) InitRouter() chi.Router {
	r := chi.NewRouter()
	r.Mount("/training/v1", s.initRouterV1())
	return r
}
func (s *server) initRouterV1() chi.Router {
	r := chi.NewRouter()
	r.Use(s.middlewareLog)
	r.Get("/alteration-methods", s.handleAlterationMethodList)
	r.Get("/alteration-methods/{id}", s.handleAlterationMethodGet)

	return r
}
