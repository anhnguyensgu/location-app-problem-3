package handler

import (
	"problem3/web-service/repository"

	"github.com/gofiber/fiber/v2"
)

type statHandler struct {
	visitStatsRepository repository.VistorStatsRepository
	eventRepository      repository.EventRepository
}

type StatHandler interface {
	GetStats(c *fiber.Ctx) error
}

type StatResponse struct {
	MostVisit   []repository.VisitorStats
	LatestVisit []repository.Event
}

func NewStatHandler(eventRepository repository.EventRepository, visitStatsRepository repository.VistorStatsRepository) StatHandler {
	return &statHandler{visitStatsRepository: visitStatsRepository, eventRepository: eventRepository}
}

func (sh *statHandler) GetStats(c *fiber.Ctx) error {
	stats := sh.visitStatsRepository.GetVisitStats(repository.PaginationParameter{Limit: 100, Skip: 0, SortField: "visit", Assending: -1})
	latestVisit := sh.eventRepository.GetAll(repository.PaginationParameter{Limit: 100, Skip: 0, SortField: "createdAt", Assending: -1})
	return c.JSON(StatResponse{MostVisit: stats, LatestVisit: latestVisit})
}
