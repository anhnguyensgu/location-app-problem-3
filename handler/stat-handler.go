package handler

import (
	"problem3/web-service/repository"

	"github.com/gofiber/fiber/v2"
)

type statHandler struct {
	visitStatsRepository repository.VistorStatsRepository
}

type StatHandler interface {
	GetStats(c *fiber.Ctx) error
}

func NewStatHandler(visitStatsRepository repository.VistorStatsRepository) StatHandler {
	return &statHandler{visitStatsRepository: visitStatsRepository}
}

func (sh *statHandler) GetStats(c *fiber.Ctx) error {
	stats := sh.visitStatsRepository.GetVisitStats()
	return c.JSON(stats)
}
