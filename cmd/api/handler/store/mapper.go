package store

import "http-client/testUser/pkg/model"

type mapper struct {
}

func (m mapper) ToDomain(user User) *model.User {
	return model.NewUser(user.Name, user.Lastname, user.Datetime, user.Address, user.Gender, user.Age)
}
