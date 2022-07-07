package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Store(bookRequestStore BookRequestStore) (Book, error)
	Update(ID int, bookRequestUpdate BookRequestUpdate) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Store(bookRequestStore BookRequestStore) (Book, error) {
	price, _ := bookRequestStore.Price.Int64()
	rating, _ := bookRequestStore.Rating.Int64()

	book := Book{
		Title:       bookRequestStore.Title,
		Description: bookRequestStore.Description,
		Price:       int(price),
		Rating:      int(rating),
	}

	newBook, err := s.repository.Store(book)

	return newBook, err

}

func (s *service) Update(ID int, BookRequestUpdate BookRequestUpdate) (Book, error) {

	book, _ := s.repository.FindByID(ID)

	price, _ := BookRequestUpdate.Price.Int64()
	rating, _ := BookRequestUpdate.Rating.Int64()

	book.Title = BookRequestUpdate.Title
	book.Description = BookRequestUpdate.Description
	book.Price = int(price)
	book.Rating = int(rating)

	updateBook, err := s.repository.Update(book)

	return updateBook, err
}
