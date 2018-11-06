package types

import (
	"time"

	"github.com/pisensor/pkg/models"
)

type DB interface {
	GetTempsBetweenTimes(serial string, start time.Time, end time.Time) ([]models.TempReading, error)
}
