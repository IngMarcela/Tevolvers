package pkg

type DeleteUserRepositoryInterface interface {
	DeleteUsers(userID int) error
}
type DeleteUserUC struct {
	DeleteUserRepository DeleteUserRepositoryInterface
}

func (d DeleteUserUC) Execute(userID int) error {
	err := d.DeleteUserRepository.DeleteUsers(userID)
	if err != nil {
		return err
	}
	return nil
}

func NewDeleteUserUC(
	deleteUserRepository DeleteUserRepositoryInterface) *DeleteUserUC {
	return &DeleteUserUC{
		DeleteUserRepository: deleteUserRepository,
	}
}
