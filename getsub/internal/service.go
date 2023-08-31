package internal

type Service struct {
	R *Repository
}

func NewService() *Service {
	return &Service{
		R: NewRepository(),
	}
}

func (s *Service) GetSubService(user *GetSubRequest) error {
	// data := user
	// deal with jwt in the future

	return nil
	// return s.R.GetSubRepository(data)
}

func (s *Service) InsertService() error {
	return s.R.InsertRepository()
}
