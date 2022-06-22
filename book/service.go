package book

type Service interface {
	FindBooks(userID int) ([]Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindBooks(userID int) ([]Book, error) {
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