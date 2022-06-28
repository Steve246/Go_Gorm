package model

import (
	"time"

	"gorm.io/gorm"
)

// type DeletedAt gorm.DeletedAt

type BaseModel struct {
	CreatedAt time.Time
	UpdateAt  time.Time
	// DeletedAt sql.NullTime
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
