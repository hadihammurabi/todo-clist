package repo

import "gorm.io/gorm"

type ListModel struct {
	*gorm.Model
	Name string
}
