package http

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/plutonio00/pay-api/internal/config"
	v1 "github.com/plutonio00/pay-api/internal/delivery/http/v1"
	"github.com/plutonio00/pay-api/internal/service"
	"github.com/swaggo/http-swagger"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(conf *config.Config) *mux.Router {
	router := mux.NewRouter()
	router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *mux.Router) {
	handlerV1 := v1.NewHandler(h.services)
	handlerV1.Init(router)
}
