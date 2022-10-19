package simple

// Provider Set
type FooBarService struct {
	*FooService
	*BarService
}

func NewFooBar(fooService *FooService, barService *BarService) *FooBarService {
	return &FooBarService{FooService: fooService, BarService: barService}
}
