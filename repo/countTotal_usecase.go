package repo

import "go_gorm/model"

type CountTotal interface {
	CountColumn(cust *model.Customer, column string) int64
}
