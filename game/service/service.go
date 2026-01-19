package service

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetHealth() (string, error) {
	return "Service is healthy", nil
}
