package pkg

import (
	"http-client/testUser/pkg/model"
)

type EditUserRepositoryInterface interface {
	EditUsers(id int) (model.User, error)
}
type EditUserUC struct {
	EditUserRepository EditUserRepositoryInterface
}

func (e EditUserUC) Execute(userID int) (model.User, error) {
	userSelected, err := e.EditUserRepository.EditUsers(userID)
	if err != nil {
		return model.User{}, err
	}
	return userSelected, nil
}

func NewEditUserUC(
	editUserRepository EditUserRepositoryInterface) *EditUserUC {
	return &EditUserUC{
		EditUserRepository: editUserRepository,
	}
}
