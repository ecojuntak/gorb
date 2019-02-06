package repositories

import (
	"github.com/ecojuntak/gorb/models"
	"github.com/jinzhu/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

type UserRepository interface {
	User(id int) (models.User, error)
	Users() ([]models.User, error)
	Create(u models.User) (models.User, error)
	Update(id int, u models.User) (models.User, error)
	Delete(id int) (bool, error)
}

func NewUserRepo(DB *gorm.DB) UserRepository {
	return &userRepo{DB}
}

func (r *userRepo) User(id int) (models.User, error) {
	var u models.User

	r.DB.Find(&u, id)

	return u, nil
}

func (r *userRepo) Users() ([]models.User, error) {
	var users []models.User

	r.DB.Find(&users)

	return users, nil
}

func (r *userRepo) Create(u models.User) (models.User, error) {
	r.DB.Create(&u)

	return u, nil
}

func (r *userRepo) Update(id int, u models.User) (models.User, error) {
	var user models.User
	r.DB.First(&user, id)

	r.DB.Model(&user).Select([]string{"name", "updated_at"}).
		Updates(map[string]interface{}{"name": u.Name, "updated_at": u.UpdatedAt})

	return user, nil
}

func (r *userRepo) Delete(id int) (bool, error) {
	var user models.User
	r.DB.First(&user, id)

	r.DB.Delete(&user)

	return true, nil
}
