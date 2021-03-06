package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Store(book Book) (Book, error)
	Update(book Book) (Book, error)
	Destroy(book Book) (Book, error)
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

	return books, err
}

func (r *repository) FindByID(ID int) (Book, error) {
	var book Book

	err := r.db.Find(&book, ID).Error

	return book, err
}

func (r *repository) Store(book Book) (Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) Update(book Book) (Book, error) {
	// err := r.db.Where("id = ?", ID).First(&book).Error
	err := r.db.Save(&book).Error

	return book, err
}

func (r *repository) Destroy(book Book) (Book, error) {
	err := r.db.Delete(&book).Error

	return book, err
}
