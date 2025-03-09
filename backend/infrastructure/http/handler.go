package http

import (
	"todo-planner/internal/schedular"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	scheduleService schedular.IService
}

type IHandler interface {
	GetSchedules(c *fiber.Ctx) error
}

func NewHandler(scheduleService schedular.IService) IHandler {
	return &Handler{scheduleService: scheduleService}
}

func (h *Handler) GetSchedules(c *fiber.Ctx) error {
	return nil
}
