package repo

import (
	"fmt"

	"gorm.io/gorm"
)

type Repo struct {
	List *ListRepo
	Todo *TodoRepo
}

func NewRepo(db *gorm.DB) *Repo {
	err := db.AutoMigrate(
		&ListModel{},
	)
	if err != nil {
		panic(fmt.Sprintf("can't prepare database: %s", err))
	}

	return &Repo{
		NewListRepo(db),
		NewTodoRepo(db),
	}
}
