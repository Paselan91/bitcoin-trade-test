package usecase

import (
	"app/src/domain/models"
	"app/src/domain/repository"
)

type EventUsecase interface {
	Create(title, content string) (*models.SignalEvent, error)
}

type eventUsecase struct {
	eventRepo repository.EventRepository
}

func NewEventUsecase(eventRepo repository.EventRepository) EventUsecase {
	return &eventUsecase{eventRepo: eventRepo}
}
