package repo

import "gorm.io/gorm"

type CountTotal interface {
	CountColumn(result interface{}, groupBy string) (*gorm.DB, error)
}
