package edit

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"http-client/testUser/infra/storage/dtos"
	"http-client/testUser/infra/storage/mapper"
	"http-client/testUser/pkg/model"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type fakeEditUserUC struct{}

func (f *fakeEditUserUC) Execute(userID int) (model.User, error) {
	userID = 1
	user := dtos.UserDTO{
		ID:       userID,
		Name:     "User " + strconv.Itoa(userID),
		LastName: "Lastname " + strconv.Itoa(userID),
		Datetime: "1990-01-01",
		Address:  "Address " + strconv.Itoa(userID),
		Gender:   "F",
		Age:      12,
	}
	usersMapper := mapper.Mapper{}.ToDomain(user)
	return *usersMapper, nil
}

type mockEditUserUC struct{}

func (m mockEditUserUC) Execute(userID int) (model.User, error) {
	return model.User{}, errors.New("error getting user")
}

func TestEdit_Execute(t *testing.T) {
	editHandler := NewEditHandler(&fakeEditUserUC{})
	req := httptest.NewRequest("GET", "/edit/1", nil)

	recorder := httptest.NewRecorder()

	editHandler.Execute(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var user User
	err := json.NewDecoder(recorder.Body).Decode(&user)
	if err != nil {
		t.Fatalf("Error decoding response body: %s", err)
	}

	assert.Equal(t, 1, user.Id)
	assert.Equal(t, "User 1", user.Name)
	assert.Equal(t, "Lastname 1", user.Lastname)
	assert.Equal(t, "1990-01-01", user.Datetime)
	assert.Equal(t, "Address 1", user.Address)
	assert.Equal(t, "F", user.Gender)
	assert.Equal(t, 12, user.Age)
}

func TestList_Execute_Error(t *testing.T) {
	mock := mockEditUserUC{}
	handler := NewEditHandler(mock)

	req, err := http.NewRequest("GET", "/edit/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.Execute(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
