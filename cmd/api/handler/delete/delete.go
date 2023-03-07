package delete

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DeleteUserUC interface {
	Execute(userID int) error
}
type Delete struct {
	deleteUserUC DeleteUserUC
}

func (d Delete) Execute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]
	id, _ := strconv.Atoi(userID)
	err := d.deleteUserUC.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "El usuario con id %s ha sido eliminado correctamente", userID)
}

func NewDeleteHandler(
	deleteUserUC DeleteUserUC,
) *Delete {
	return &Delete{
		deleteUserUC: deleteUserUC,
	}
}
