package store

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"http-client/testUser/pkg/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeStoreUserUC struct {
	users []model.User
}

func (f *fakeStoreUserUC) Execute(user *model.User) error {
	f.users = append(f.users, *user)
	return nil
}

func TestStore_Execute(t *testing.T) {
	handler := NewStoreHandler(&fakeStoreUserUC{})

	user := model.NewUser("User 1", "Lastname 1", "1990-01-01", "Address 1", "M", 12)
	jsonData, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/store", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler.Execute(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, 1, len(fakeStoreUserUC{}.users))
	assert.Equal(t, user, fakeStoreUserUC{}.users[0])
}
