package model

import (
	"time"

	"gorm.io/gorm"
)

type RAM struct {
	gorm.Model
	MemoryUsage float32
	TimeStamp   time.Time
}
