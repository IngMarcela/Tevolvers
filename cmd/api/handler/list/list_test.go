package list

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"http-client/testUser/infra/storage/dtos"
	"http-client/testUser/infra/storage/mapper"
	"http-client/testUser/pkg/model"
)

type fakeListUserUC struct{}

func (f *fakeListUserUC) Execute() ([]model.User, error) {
	usersList := []dtos.UserDTO{
		{
			ID:       1,
			Name:     "User 1",
			LastName: "Daniela",
			Datetime: "1990-01-01",
			Address:  "ariza",
			Gender:   "F",
			Age:      12,
		},
		{
			ID:       2,
			Name:     "User 2",
			LastName: "Daniela",
			Datetime: "1990-01-01",
			Address:  "ariza",
			Gender:   "F",
			Age:      12,
		},
		{
			ID:       3,
			Name:     "User 3",
			LastName: "Daniela",
			Datetime: "1990-01-01",
			Address:  "ariza",
			Gender:   "F",
			Age:      12,
		},
	}
	var arrayUser []model.User
	for _, user := range usersList {
		usersMapper := mapper.Mapper{}.ToDomain(user)
		arrayUser = append(arrayUser, *usersMapper)
	}
	return arrayUser, nil
}

type mockListUserUC struct{}

func (m mockListUserUC) Execute() ([]model.User, error) {
	return nil, errors.New("error getting users list")
}

func TestList_Execute(t *testing.T) {
	listHandler := NewListHandler(&fakeListUserUC{})
	req := httptest.NewRequest("GET", "/", nil)

	recorder := httptest.NewRecorder()

	listHandler.Execute(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var users []User
	err := json.NewDecoder(recorder.Body).Decode(&users)
	if err != nil {
		t.Fatalf("Error decoding response body: %s", err)
	}

	assert.Len(t, users, 3)
	assert.Equal(t, 1, users[0].Id)
	assert.Equal(t, "User 1", users[0].Name)
	assert.Equal(t, 2, users[1].Id)
	assert.Equal(t, "User 2", users[1].Name)
	assert.Equal(t, 3, users[2].Id)
	assert.Equal(t, "User 3", users[2].Name)
}

func TestList_Execute_Error(t *testing.T) {
	mock := mockListUserUC{}
	handler := NewListHandler(mock)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.Execute(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "error getting users list")
}
