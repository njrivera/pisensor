package types

import (
	"github.com/pisensor/pkg/models"
)

type DB interface {
	InsertReading(models.TempReading) error
}
