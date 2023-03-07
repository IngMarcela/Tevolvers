package daos

import (
	"gorm.io/gorm"
	"http-client/testUser/infra/storage/dtos"
	"http-client/testUser/infra/storage/mapper"
	"http-client/testUser/pkg/model"
)

type UserDao struct {
	connexion *gorm.DB
}

func (u UserDao) Save(user *model.User) error {
	userDTO := dtos.NewUserDTO(user)
	err := u.connexion.Create(&userDTO).Error
	if err != nil {
		return err
	}
	return nil
}

func (u UserDao) DeleteUsers(userID int) error {
	err := u.connexion.Delete(&model.User{}, "id = ?", userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (u UserDao) UpdateUsers(user *model.User, userID int) error {
	userDTO := dtos.NewUserDTO(user)
	err := u.connexion.Model(&model.User{}).Where("id = ?", userID).Updates(&userDTO).Error
	if err != nil {
		return err
	}
	return nil
}

func (u UserDao) EditUsers(id int) (model.User, error) {
	userSelectedDto := dtos.UserDTO{}
	err := u.connexion.Where("id = ?", id).Find(&userSelectedDto).Error
	if err != nil {
		return model.User{}, err
	}
	user := mapper.Mapper{}.ToDomain(userSelectedDto)
	return *user, nil
}

func (u UserDao) GetUsers() ([]model.User, error) {
	var usersDto []dtos.UserDTO
	err := u.connexion.Find(&usersDto).Error
	if err != nil {
		return nil, err
	}
	var arrayUser []model.User
	for _, userDto := range usersDto {
		user := mapper.Mapper{}.ToDomain(userDto)
		arrayUser = append(arrayUser, *user)
	}

	return arrayUser, nil
}

func NewUserDAO(connexion *gorm.DB) *UserDao {
	return &UserDao{
		connexion: connexion,
	}
}
