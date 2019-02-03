package repositories

import (
	"database/sql"
	"log"

	"github.com/ecojuntak/gorb/models"
)

type userRepo struct {
	DB *sql.DB
}

type UserRepository interface {
	User(id int) (models.User, error)
	Users() ([]models.User, error)
	Create(u models.User) (models.User, error)
	Update(id int, u models.User) (models.User, error)
	Delete(id int) (bool, error)
}

func NewUserRepo(DB *sql.DB) UserRepository {
	return &userRepo{DB}
}

func (r *userRepo) User(id int) (models.User, error) {
	var u models.User

	err := r.DB.QueryRow("select * from users where id=?", id).
		Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)

	return u, err
}

func (r *userRepo) Users() ([]models.User, error) {
	var users []models.User

	rows, err := r.DB.Query(`select * from users`)
	if err != nil {
		log.Fatalln(err)
	}

	defer rows.Close()

	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)

		if err != nil {
			log.Fatalln(err)
		}

		users = append(users, u)
	}

	return users, nil
}

func (r *userRepo) Create(u models.User) (models.User, error) {
	stmt, err := r.DB.Prepare("insert users set name=?, email=?, created_at=?, updated_at=?")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := stmt.Exec(u.Name, u.Email, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		log.Fatalln(err)
	}

	id, err := res.LastInsertId()
	idInt := int(id)

	user, _ := r.User(idInt)

	return user, err
}

func (r *userRepo) Update(i int, u models.User) (models.User, error) {
	stmt, err := r.DB.Prepare("update users set name=?, email=?, updated_at=? where id=?")

	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(u.Name, u.Email, u.UpdatedAt, i)
	if err != nil {
		log.Fatalln(err)
	}

	user, _ := r.User(i)

	return user, err
}

func (r *userRepo) Delete(id int) (bool, error) {
	stmt, err := r.DB.Prepare("delete from users where id=?")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatalln(err)
	}

	return true, nil
}
