package controllers

import "github.com/hdysh/dqtsearch/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	s.Router.HandleFunc("/units", middlewares.SetMiddlewareJSON(s.GetUnits)).Methods("GET")
	s.Router.HandleFunc("/unitfilter", middlewares.SetMiddlewareJSON(s.GetUnitsFiltered)).Methods("GET")
	s.Router.HandleFunc("/units/{id}", middlewares.SetMiddlewareJSON(s.GetUnit)).Methods("GET")
}
