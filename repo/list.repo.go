package repo

import "gorm.io/gorm"

type ListRepo struct {
	db *gorm.DB
}

func NewListRepo(db *gorm.DB) *ListRepo {
	return &ListRepo{db}
}

func (r *ListRepo) GetAll() ([]*ListModel, error) {
	lists := make([]*ListModel, 0)
	err := r.db.Find(&lists).Error
	return lists, err
}

func (r *ListRepo) Insert(input *ListModel) error {
	return r.db.Save(input).Error
}
