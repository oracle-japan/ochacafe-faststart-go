package repo

import (
	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	Name       string
	Date       string
	Topics     string
	Presenters string
}
