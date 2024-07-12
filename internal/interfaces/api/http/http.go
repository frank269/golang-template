package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mpt-tiendc/golang-template/internal/config"
)

type HttpServer interface {
	Start() error
}

type HttpServerImpl struct {
	config *config.AppConfig
	mux    *chi.Mux
}

func NewHttpServer(config *config.AppConfig) *HttpServerImpl {
	return &HttpServerImpl{
		config: config,
		mux:    chi.NewRouter(),
	}
}

func (h *HttpServerImpl) Start() error {
	err := h.useRoutes()
	if err != nil {
		return err
	}
	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", h.config.Host, h.config.Port),
		Handler:           h.mux,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	return server.ListenAndServe()
}
