package pkg

import (
	"http-client/testUser/pkg/model"
)

type GetUsersRepositoryInterface interface {
	GetUsers() ([]model.User, error)
}
type ListUserUC struct {
	ListUsersRepository GetUsersRepositoryInterface
}

func (l ListUserUC) Execute() ([]model.User, error) {
	users, err := l.ListUsersRepository.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
func NewListUserUC(
	listUsersRepository GetUsersRepositoryInterface) *ListUserUC {
	return &ListUserUC{
		ListUsersRepository: listUsersRepository,
	}
}
