package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

type SimpleService struct {
	*SimpleRepository
	Action string
}

func NewSimpleService(simpleRepository *SimpleRepository) (*SimpleService, error) {
	if simpleRepository.Error {
		return nil, errors.New("Failed to create service")
	} else {
		return &SimpleService{SimpleRepository: simpleRepository, Action: "Add"}, nil
	}
}
