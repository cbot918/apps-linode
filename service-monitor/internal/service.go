package internal

type Service struct {
	R *Repository
}

func NewService() *Service {
	return &Service{
		R: NewRepository(),
	}
}

func (s *Service) SignupService(user *SignupRequest) error {
	data := user
	// deal with jwt in the future

	return s.R.SignupRepository(data)
}

func (s *Service) InsertService() error {
	return s.R.InsertRepository()
}
