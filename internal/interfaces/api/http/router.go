package http

import (
	"github.com/go-chi/chi"
	"github.com/mpt-tiendc/golang-template/internal/infrastructure/handlers"
	"github.com/mpt-tiendc/golang-template/internal/infrastructure/webserver"
)

func (s *HttpServerImpl) useRoutes() error {
	userHandler, err := handlers.NewUserHandler(s.config)
	if err != nil {
		return err
	}
	s.mux.Route("/api/v1", func(r chi.Router) {
		r.Use(webserver.Cors())
		r.Post("/user", userHandler.CreateUser())
		r.Get("/user/{id}", userHandler.GetUserById())
		r.Put("/user/{id}", userHandler.UpdateUser())
		r.Delete("/user/{id}", userHandler.DeleteUser())
	})
	return nil
}
