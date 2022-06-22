package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindByUserID(userID int) ([]Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book

	err := r.db.Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func (r *repository) FindByUserID(userID int) ([]Book, error) {
	var books []Book

	err := r.db.Where("user_id = ?", userID).Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}