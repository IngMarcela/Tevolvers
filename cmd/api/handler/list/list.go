package list

import (
	"encoding/json"
	"http-client/testUser/pkg/model"
	"net/http"
)

type ListUserUC interface {
	Execute() ([]model.User, error)
}
type list struct {
	listUserUC ListUserUC
}

func (l list) Execute(w http.ResponseWriter, r *http.Request) {
	usersList, err := l.listUserUC.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var arrayUser []User
	for _, user := range usersList {
		usersMapper := Mapper{}.ToDomainResponse(user)
		arrayUser = append(arrayUser, usersMapper)
	}

	jsonData, err := json.Marshal(arrayUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func NewListHandler(
	listUserUC ListUserUC,
) *list {
	return &list{
		listUserUC: listUserUC,
	}
}
