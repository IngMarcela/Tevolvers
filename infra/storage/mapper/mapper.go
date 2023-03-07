package mapper

import (
	"http-client/testUser/infra/storage/dtos"
	"http-client/testUser/pkg/model"
)

type Mapper struct {
}

func (m Mapper) ToDomain(user dtos.UserDTO) *model.User {
	return model.NewUserRecovery(user.ID, user.Name, user.LastName, user.Datetime, user.Address, user.Gender, user.Age)
}
