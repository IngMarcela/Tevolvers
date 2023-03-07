package dtos

import (
	"http-client/testUser/pkg/model"
)

type UserDTO struct {
	ID       int    `gorm:"column:id;PRIMARY_KEY"`
	Name     string `gorm:"column:name;not null"`
	LastName string `gorm:"column:lastname;not null"`
	Datetime string `gorm:"column:datetime;not null"`
	Address  string `gorm:"column:address;not null"`
	Gender   string `gorm:"column:gender;not null"`
	Age      int    `gorm:"column:age;not null"`
}

func NewUserDTO(user *model.User) UserDTO {
	return UserDTO{
		ID:       user.Id(),
		Name:     user.Name(),
		LastName: user.LastName(),
		Datetime: user.Datetime(),
		Address:  user.Address(),
		Gender:   user.Gender(),
		Age:      user.Age(),
	}
}

func (UserDTO) TableName() string {
	return "users"
}
