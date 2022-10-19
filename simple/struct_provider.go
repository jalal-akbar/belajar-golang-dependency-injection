package simple

type A struct {
}

func NewA() *A {
	return &A{}
}

type B struct {
}

func NewB() *B {
	return &B{}
}

// struct provider
type AdanB struct {
	*A
	*B
}
