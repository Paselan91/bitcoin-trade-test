package repository

import "app/src/domain/models"

type EventRepository interface {
	Create(task *models.SignalEvent) (*models.SignalEvent, error)
}
