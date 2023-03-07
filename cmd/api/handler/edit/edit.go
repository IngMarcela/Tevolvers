package edit

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"http-client/testUser/pkg/model"
	"net/http"
	"strconv"
)

type EditUserUC interface {
	Execute(userID int) (model.User, error)
}
type Edit struct {
	editUserUC EditUserUC
}

func (e Edit) Execute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	id, _ := strconv.Atoi(userID)
	userSelected, err := e.editUserUC.Execute(id)
	fmt.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	usersMapper := Mapper{}.ToDomainResponse(userSelected)

	jsonData, err := json.Marshal(usersMapper)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func NewEditHandler(
	editUserUC EditUserUC,
) *Edit {
	return &Edit{
		editUserUC: editUserUC,
	}
}
