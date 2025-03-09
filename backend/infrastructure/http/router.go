package http

import (
	"fmt"
	"todo-planner/infrastructure"

	"github.com/gofiber/fiber/v2"
)

type IRouter interface {
	SetupRoutes() 
	Start() error
}

type Router struct {
	router *fiber.App
	config infrastructure.Config
	handler IHandler
}

func NewRouter(config infrastructure.Config, handler IHandler) IRouter {
	router := fiber.New()
	return &Router{
		router: router,
		config: config,
		handler: handler,
	}
}

func (r *Router) SetupRoutes() {
	r.router.Get("/schedules", r.handler.GetSchedules)
}

func (r *Router) Start() error {
	return r.router.Listen(fmt.Sprintf(":%d", r.config.HTTP.Port))
}