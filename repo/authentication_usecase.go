package repo

import "go_gorm/model"

type AuthCheck interface {
	AuthLogin(checkLogin *model.Customer) error
}
