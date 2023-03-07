package pkg

import (
	"http-client/testUser/pkg/model"
)

type UpdateUserRepositoryInterface interface {
	UpdateUsers(user *model.User, userID int) error
}
type UpdateUserUC struct {
	UpdateUserRepository UpdateUserRepositoryInterface
}

func (u UpdateUserUC) Execute(user *model.User, userID int) error {
	err := u.UpdateUserRepository.UpdateUsers(user, userID)
	if err != nil {
		return err
	}
	return nil
}

func NewUpdateUserUC(
	updateUserRepository UpdateUserRepositoryInterface) *UpdateUserUC {
	return &UpdateUserUC{
		UpdateUserRepository: updateUserRepository,
	}
}
