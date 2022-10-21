package simple

import (
	"errors"
)

// struct
type SimpleRepository struct {
	Error bool // Error Handling
}

type SimpleService struct {
	*SimpleRepository // Depend on Simple Repository
}

// Provide or Constructtor
func NewSimpleRepository(isError bool /*paramter injection*/) *SimpleRepository {
	return &SimpleRepository{Error: isError}
}

func NewSimpleService(simpleRepository *SimpleRepository) (*SimpleService, error) { //Error Handling
	if simpleRepository.Error {
		return nil, errors.New("failed NewSimpleService")
	} else {
		return &SimpleService{
			SimpleRepository: simpleRepository,
		}, nil
	}
}
