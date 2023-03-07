package repository

import (
	"http-client/testUser/infra/storage/daos"
	"http-client/testUser/pkg/model"
)

type EmployeesRepository struct {
	connexion *daos.UserDao
}

func (e EmployeesRepository) StoreUsers(user *model.User) error {
	err := e.connexion.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (e EmployeesRepository) DeleteUsers(userID int) error {
	err := e.connexion.DeleteUsers(userID)
	if err != nil {
		return err
	}
	return nil
}

func (e EmployeesRepository) UpdateUsers(user *model.User, userID int) error {
	err := e.connexion.UpdateUsers(user, userID)
	if err != nil {
		return err
	}
	return nil
}

func (e EmployeesRepository) EditUsers(id int) (model.User, error) {
	userSelected, err := e.connexion.EditUsers(id)
	if err != nil {
		return model.User{}, err
	}
	return userSelected, nil
}

func (e EmployeesRepository) GetUsers() ([]model.User, error) {
	users, err := e.connexion.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewEmployeesRepository(connexion *daos.UserDao) *EmployeesRepository {
	return &EmployeesRepository{
		connexion: connexion,
	}
}
