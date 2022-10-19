package simple

type SayHello interface {
	Hello(name string) string
}

type SayHelloService struct {
	SayHello
}
type SayHelloImpl struct {
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello" + name
}

func NewSayHelloService(sayHello SayHello) *SayHelloService {
	return &SayHelloService{SayHello: sayHello}
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}
