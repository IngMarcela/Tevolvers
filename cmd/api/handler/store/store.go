package store

import (
	"encoding/json"
	"http-client/testUser/pkg/model"
	"net/http"
)

type StoreUserUC interface {
	Execute(user *model.User) error
}
type Store struct {
	storeUserUC StoreUserUC
}

func (s Store) Execute(w http.ResponseWriter, r *http.Request) {
	var request User
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := mapper{}.ToDomain(request)
	err = s.storeUserUC.Execute(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewStoreHandler(
	storeUserUC StoreUserUC,
) *Store {
	return &Store{
		storeUserUC: storeUserUC,
	}
}
