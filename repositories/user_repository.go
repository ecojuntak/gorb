package repositories

import (
	"github.com/ecojuntak/gorb/models"
	"github.com/jinzhu/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

type UserRepository interface {
	User(id int) models.User
	Users() []models.User
	Create(u models.User) models.User
	Update(id int, u models.User) models.User
	Delete(id int) bool
}

func NewUserRepo(DB *gorm.DB) UserRepository {
	return &userRepo{DB}
}

func (r *userRepo) User(id int) models.User {
	var u models.User

	r.DB.Find(&u, id)

	return u
}

func (r *userRepo) Users() []models.User {
	var uu []models.User

	r.DB.Find(&uu)

	return uu
}

func (r *userRepo) Create(u models.User) models.User {
	r.DB.Create(&u)

	return u
}

func (r *userRepo) Update(id int, u models.User) models.User {
	var user models.User
	r.DB.First(&user, id)

	r.DB.Model(&user).Select([]string{"name", "updated_at"}).
		Updates(map[string]interface{}{"name": u.Name, "updated_at": u.UpdatedAt})

	return user
}

func (r *userRepo) Delete(id int) bool {
	var user models.User
	r.DB.First(&user, id)

	r.DB.Delete(&user)

	return true
}
