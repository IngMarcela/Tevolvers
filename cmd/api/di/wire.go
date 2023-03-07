package di

import (
	"github.com/gorilla/mux"
	delete2 "http-client/testUser/cmd/api/handler/delete"
	"http-client/testUser/cmd/api/handler/edit"
	"http-client/testUser/cmd/api/handler/list"
	"http-client/testUser/cmd/api/handler/store"
	"http-client/testUser/cmd/api/handler/update"
	"http-client/testUser/infra/storage/daos"
	"http-client/testUser/infra/storage/repository"
	"http-client/testUser/pkg/usecase"
)

func Initialize() (*mux.Router, error) {
	userDaos := daos.NewUserDAO(connectionBDProvider())

	employeesRepository := repository.NewEmployeesRepository(userDaos)

	listUserUC := pkg.NewListUserUC(employeesRepository)
	storeUserUC := pkg.NewStoreUserUC(employeesRepository)
	editUserUC := pkg.NewEditUserUC(employeesRepository)
	updateUserUC := pkg.NewUpdateUserUC(employeesRepository)
	deleteUserUC := pkg.NewDeleteUserUC(employeesRepository)

	handlerList := list.NewListHandler(listUserUC)
	handlerStore := store.NewStoreHandler(storeUserUC)
	handlerEdit := edit.NewEditHandler(editUserUC)
	handlerUpdate := update.NewUpdateHandler(updateUserUC)
	handlerDelete := delete2.NewDeleteHandler(deleteUserUC)

	router := mux.NewRouter()
	router.HandleFunc("/", handlerList.Execute).Methods("GET")
	router.HandleFunc("/store", handlerStore.Execute).Methods("POST")
	router.HandleFunc("/edit/{id}", handlerEdit.Execute).Methods("GET")
	router.HandleFunc("/update/{id}", handlerUpdate.Execute).Methods("PUT")
	router.HandleFunc("/delete/{id}", handlerDelete.Execute).Methods("POST")

	return router, nil
}
