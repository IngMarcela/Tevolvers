package update

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"http-client/testUser/pkg/model"
	"net/http"
	"strconv"
)

type UpdateUserUC interface {
	Execute(user *model.User, userID int) error
}
type Update struct {
	updateUserUC UpdateUserUC
}

func (u Update) Execute(w http.ResponseWriter, r *http.Request) {
	var request User
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := mapper{}.ToDomain(request)
	vars := mux.Vars(r)
	userID := vars["id"]
	id, _ := strconv.Atoi(userID)

	err = u.updateUserUC.Execute(user, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Usuario con id %s actualizado correctamente", userID)
}

func NewUpdateHandler(
	updateUserUC UpdateUserUC,
) *Update {
	return &Update{
		updateUserUC: updateUserUC,
	}
}
