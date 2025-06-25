package services

type HelloService struct {
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (service *HelloService) SayHello() string {
	return "Good morning"
}
