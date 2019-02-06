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

func (c *UserController) Resources(w http.ResponseWriter, r *http.Request) {
	switch m := r.Method; m {
	case http.MethodGet:
		params := mux.Vars(r)
		if len(params) == 0 {
			c.Users(w, r)
		} else {
			c.User(w, r)
		}
	case http.MethodPost:
		c.Create(w, r)
	case http.MethodPatch:
		c.Update(w, r)
	case http.MethodDelete:
		c.Delete(w, r)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func modifyResponse(user models.User) map[string]interface{} {
	u := make(map[string]interface{})
	u["id"] = user.ID
	u["name"] = user.Name
	u["email"] = user.Email
	u["created_at"] = user.CreatedAt
	u["updated_at"] = user.UpdatedAt
	u["deteled_at"] = user.DeletedAt

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
	now := time.Now()

	u := models.User{
		Name:      r.FormValue("name"),
		Email:     r.FormValue("email"),
		UpdatedAt: now,
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
