package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ecojuntak/gorb/controllers"
	"github.com/ecojuntak/gorb/models"
	"github.com/ecojuntak/gorb/repositories/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUsers(t *testing.T) {
	uu := make([]models.User, 1)
	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Users").Return(uu)

	uc := controllers.NewUserController(mockUserRepo)

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uc.Resources)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestUser(t *testing.T) {
	u := models.User{}

	id := 0

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("User", id).Return(u)

	uc := controllers.NewUserController(mockUserRepo)

	req, err := http.NewRequest("GET", "/users/"+strconv.Itoa(id), nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", uc.Resources)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestUserShouldFailed(t *testing.T) {
	u := models.User{}

	id := 0

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("User", id).Return(u)

	uc := controllers.NewUserController(mockUserRepo)

	req, err := http.NewRequest("GET", "/users/"+strconv.Itoa(id), nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{name}", uc.Resources)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestCreate(t *testing.T) {
	now := time.Now()

	u := models.User{
		Name:      "eco",
		Email:     "eco@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	tempUser := u
	tempUser.ID = 1

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Create", mock.AnythingOfType("models.User")).Return(u)

	uc := controllers.NewUserController(mockUserRepo)

	payload, err := json.Marshal(tempUser)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/users", strings.NewReader(string(payload)))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users", uc.Resources)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestUpdate(t *testing.T) {
	now := time.Now()

	u := models.User{
		ID:        1,
		Name:      "eco",
		Email:     "eco@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	id := int(u.ID)

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Update", id, mock.AnythingOfType("models.User")).Return(u)

	uc := controllers.NewUserController(mockUserRepo)

	payload, err := json.Marshal(u)
	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/users/"+strconv.Itoa(id), strings.NewReader(string(payload)))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", uc.Resources)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestUpdateShouldFailed(t *testing.T) {
	now := time.Now()

	u := models.User{
		ID:        1,
		Name:      "eco",
		Email:     "eco@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	id := int(u.ID)

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Update", id, mock.AnythingOfType("models.User")).Return(u)

	uc := controllers.NewUserController(mockUserRepo)

	payload, err := json.Marshal(u)
	assert.NoError(t, err)

	req, err := http.NewRequest("PATCH", "/users/"+strconv.Itoa(id), strings.NewReader(string(payload)))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{name}", uc.Resources)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestDetele(t *testing.T) {
	id := 1

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Delete", id).Return(true)

	uc := controllers.NewUserController(mockUserRepo)

	req, err := http.NewRequest("DELETE", "/users/"+strconv.Itoa(id), nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", uc.Resources)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestDeteleShouldFailed(t *testing.T) {
	id := 1

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Delete", id).Return(true)

	uc := controllers.NewUserController(mockUserRepo)

	req, err := http.NewRequest("DELETE", "/users/"+strconv.Itoa(id), nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{name}", uc.Resources)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestMethodNotAllows(t *testing.T) {
	id := 1

	msg := map[string]string{"error": "Method not allowed"}

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("Delete", id).Return(msg)

	uc := controllers.NewUserController(mockUserRepo)

	req, err := http.NewRequest("OPTION", "/users/"+strconv.Itoa(id), nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", uc.Resources)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Controller returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}
