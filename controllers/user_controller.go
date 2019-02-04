package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ecojuntak/gorb/models"

	"github.com/gorilla/mux"

	"github.com/ecojuntak/gorb/repositories"
)

type UserController struct {
	repo repositories.UserRepository
}

func NewUserController(repo repositories.UserRepository) UserController {
	return UserController{repo}
}

func modifyResponse(user models.User) map[string]interface{} {
	u := make(map[string]interface{})
	u["id"] = user.ID
	u["name"] = user.Name
	u["email"] = user.Email
	u["created_at"] = user.CreatedAt
	u["updated_at"] = user.UpdatedAt

	return u
}

func (c *UserController) Users(w http.ResponseWriter, r *http.Request) {
	users, err := c.repo.Users()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var uu []map[string]interface{}

	for _, user := range users {
		uu = append(uu, modifyResponse(user))
	}

	respondWithJSON(w, http.StatusOK, uu)
}

func (c *UserController) User(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	user, err := c.repo.User(id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u := modifyResponse(user)

	respondWithJSON(w, http.StatusOK, u)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	u := models.User{
		Name:      r.FormValue("name"),
		Email:     r.FormValue("email"),
		CreatedAt: now,
		UpdatedAt: now,
	}

	res, err := c.repo.Create(u)
	user := modifyResponse(res)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	u := models.User{
		Name:      r.FormValue("name"),
		Email:     r.FormValue("email"),
		UpdatedAt: time.Now(),
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := c.repo.Update(id, u)
	user := modifyResponse(res)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ok, err := c.repo.Delete(id)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if ok {
		respondWithJSON(w, http.StatusOK, true)
	}
}
