package service

type Service interface {
	GetOwnName() string
}

type ServiceV1 struct{}

func (s *ServiceV1) GetOwnName() string {
	return "test-rpc"
}
