package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	"github.com/ecojuntak/gorb/controllers"
	"github.com/ecojuntak/gorb/models"
	"github.com/ecojuntak/gorb/repositories/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUsers(t *testing.T) {
	var uu []models.User
	err := faker.FakeData(&uu)
	assert.NoError(t, err)

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Users").Return(uu, nil)

	uc := controllers.NewUserController(mockUserRepo)

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uc.Users)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestUser(t *testing.T) {
	var u models.User
	err := faker.FakeData(&u)
	assert.NoError(t, err)

	id := int(u.ID)

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("User", id).Return(u, nil)

	uc := controllers.NewUserController(mockUserRepo)

	req, err := http.NewRequest("GET", "/user/"+strconv.Itoa(id), nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uc.User)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestCreate(t *testing.T) {
	u := models.User{
		Name:      "eco",
		Email:     "eco@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tempUser := u
	tempUser.ID = 1

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Create", mock.AnythingOfType("models.User")).Return(u, nil)

	uc := controllers.NewUserController(mockUserRepo)

	payload, err := json.Marshal(tempUser)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/users", strings.NewReader(string(payload)))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uc.Create)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestUpdate(t *testing.T) {
	var u models.User
	err := faker.FakeData(&u)
	assert.NoError(t, err)

	id := int(u.ID)

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Update", id, mock.AnythingOfType("models.User")).Return(u, nil)

	uc := controllers.NewUserController(mockUserRepo)

	payload, err := json.Marshal(u)
	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/users/"+strconv.Itoa(id), strings.NewReader(string(payload)))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", uc.Update)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestDetele(t *testing.T) {
	var u models.User
	err := faker.FakeData(&u)
	assert.NoError(t, err)

	id := int(u.ID)

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Delete", id).Return(true, nil)

	uc := controllers.NewUserController(mockUserRepo)

	req, err := http.NewRequest("DELETE", "/users/"+strconv.Itoa(id), nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", uc.Delete)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}
