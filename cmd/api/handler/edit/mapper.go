package edit

import (
	"http-client/testUser/pkg/model"
)

type Mapper struct {
}

func (m Mapper) ToDomainResponse(user model.User) User {
	return User{
		user.Id(), user.Name(), user.LastName(), user.Datetime(), user.Address(), user.Gender(), user.Age(),
	}
}
