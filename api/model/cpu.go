package model

import (
	"time"

	"gorm.io/gorm"
)

type CPU struct {
	gorm.Model
	CurrentPercent float32
	TimeStamp      time.Time
}
