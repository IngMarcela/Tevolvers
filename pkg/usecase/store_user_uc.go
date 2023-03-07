package pkg

import (
	"http-client/testUser/pkg/model"
)

type CreateUserRepositoryInterface interface {
	StoreUsers(user *model.User) error
}
type CreateUserUC struct {
	CreateUserRepository CreateUserRepositoryInterface
}

func (c CreateUserUC) Execute(user *model.User) error {
	err := c.CreateUserRepository.StoreUsers(user)
	if err != nil {
		return err
	}
	return nil
}

func NewStoreUserUC(
	createUserRepository CreateUserRepositoryInterface) *CreateUserUC {
	return &CreateUserUC{
		CreateUserRepository: createUserRepository,
	}
}
