package handler

import (
	"log"
	"problem3/web-service/repository"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type eventHandler struct {
	eventRepository      repository.EventRepository
	visitStatsRepository repository.VistorStatsRepository
}

type EventHandler interface {
	GetEvents(c *fiber.Ctx) error
	CreateEvent(c *fiber.Ctx) error
	validateEventRequest(event repository.Event) []*ErrorResponse
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func NewEventHandler(eventRepository repository.EventRepository, visitStatsRepository repository.VistorStatsRepository) EventHandler {
	return &eventHandler{eventRepository: eventRepository, visitStatsRepository: visitStatsRepository}
}

func (ep *eventHandler) validateEventRequest(event repository.Event) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(event)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (ep *eventHandler) GetEvents(c *fiber.Ctx) error {
	events := ep.eventRepository.GetAll(repository.PaginationParameter{Limit: 100, Skip: 0, SortField: "createdAt", Assending: -1})
	return c.JSON(events)
}

func (ep *eventHandler) CreateEvent(c *fiber.Ctx) error {
	event := &repository.Event{}
	if err := c.BodyParser(event); err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	event.IpAddress = c.IP()

	errors := ep.validateEventRequest(*event)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}
	ep.eventRepository.SaveEvent(*event)
	go ep.visitStatsRepository.IncreaseVistorCount(*event)
	return c.JSON(event)
}
