package quilltypes

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string
	Notes    []Notes
	AuthID   int
	Auth     Auth
}

type Auth struct {
	gorm.Model
	Username string
	Password string
}
