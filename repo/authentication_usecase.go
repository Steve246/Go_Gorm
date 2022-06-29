package repo

import "go_gorm/model"

type AuthCheck interface {
	AuthLogin(name string, password string) (model.UserCredential, error)
}
