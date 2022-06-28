package main

import (
	"database/sql"
	"time"
)

// type DeletedAt gorm.DeletedAt

type BaseModel struct {
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt sql.NullTime
}
