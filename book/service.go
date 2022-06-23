package book

type Service interface {
	GetBooks(userID int) ([]Book, error)
	GetBookByID(input GetBookDetailInput) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetBooks(userID int) ([]Book, error) {
	if userID != 0 {
		books, err := s.repository.FindByUserID(userID)
		if err != nil {
			return books, err
		}

		return books, nil
	}

	books, err := s.repository.FindAll()
	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *service) GetBookByID(input GetBookDetailInput) (Book, error) {
	book, err := s.repository.FindByID(input.ID)
	if err != nil {
		return book, err
	}

	return book, nil
}