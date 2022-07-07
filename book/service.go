package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Store(bookRequestStore BookRequestStore) (Book, error)
	Update(ID int, bookRequestUpdate BookRequestUpdate) (Book, error)
	Destroy(ID int) (Book, error)
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

func (s *service) Update(ID int, bookRequestUpdate BookRequestUpdate) (Book, error) {

	book, _ := s.repository.FindByID(ID)

	price, _ := bookRequestUpdate.Price.Int64()
	rating, _ := bookRequestUpdate.Rating.Int64()

	book.Title = bookRequestUpdate.Title
	book.Description = bookRequestUpdate.Description
	book.Price = int(price)
	book.Rating = int(rating)

	updateBook, err := s.repository.Update(book)

	return updateBook, err

}

func (s *service) Destroy(ID int) (Book, error) {
	book, _ := s.repository.FindByID(ID)

	destroyBook, err := s.repository.Destroy(book)

	return destroyBook, err
}
