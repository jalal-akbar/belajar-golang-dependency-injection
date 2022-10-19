package simple

// Provider Set
type Foo struct{}

func NewFoo() *Foo {
	return &Foo{}
}

type FooService struct {
	*Foo
}

func NewFooService(foo *Foo) *FooService {
	return &FooService{Foo: foo}
}
