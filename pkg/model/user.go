package model

type User struct {
	id       int
	name     string
	lastname string
	datetime string
	address  string
	gender   string
	age      int
}

func NewUser(name string, lastname string, datetime string, address string, gender string, age int) *User {
	return &User{
		name:     name,
		lastname: lastname,
		datetime: datetime,
		address:  address,
		gender:   gender,
		age:      age,
	}

}

func NewUserRecovery(id int, name string, lastname string, datetime string, address string, gender string, age int) *User {
	return &User{
		id:       id,
		name:     name,
		lastname: lastname,
		datetime: datetime,
		address:  address,
		gender:   gender,
		age:      age,
	}

}

func (u User) AccountID() string {
	return u.AccountID()
}

func (u User) Name() string {
	return u.name
}

func (u User) LastName() string {
	return u.lastname
}

func (u User) Datetime() string {
	return u.datetime
}

func (u User) Address() string {
	return u.address
}

func (u User) Gender() string {
	return u.gender
}

func (u User) Age() int {
	return u.age
}

func (u User) Id() int {
	return u.id
}
