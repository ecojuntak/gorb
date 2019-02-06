package repositories

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/ecojuntak/gorb/models"
	"github.com/stretchr/testify/assert"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	now := time.Now()

	u := models.User{
		ID:        1,
		Name:      "eco",
		Email:     "eco@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "updated_at", "created_at"}).
		AddRow(1, "eco", "eco@example.com", now, now)

	query := "^SELECT (.+) FROM \"users\" WHERE (.+)$"

	mock.ExpectQuery(query).WillReturnRows(rows)
	ur := NewUserRepo(gormDB)

	user, err := ur.User(1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, u, user)
}

func TestGetByAllUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "updated_at", "created_at"}).
		AddRow(1, "eco", "eco@example.com", time.Now(), time.Now()).
		AddRow(2, "eco", "eco@site.com", time.Now(), time.Now())

	query := "^SELECT (.+) FROM \"users\""

	mock.ExpectQuery(query).WillReturnRows(rows)
	ur := NewUserRepo(gormDB)

	users, err := ur.Users()
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestCreateUser(t *testing.T) {
	now := time.Now()
	u := models.User{
		Name:      "eco",
		Email:     "eco@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "^INSERT INTO \"users\" (.+) VALUES (.+)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Name, u.Email, u.CreatedAt, u.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	ur := NewUserRepo(gormDB)

	user, err := ur.Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUpdate(t *testing.T) {
	now := time.Now()
	u := models.User{
		ID:        1,
		Name:      "eco",
		Email:     "eco@example.com",
		UpdatedAt: now,
	}

	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "^UPDATE \"users\" set name=\\?, updated_at=\\? where id=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Name, u.UpdatedAt, u.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	ur := NewUserRepo(gormDB)

	user, err := ur.Update(1, u)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	fmt.Print(user)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "^DELETE FROM \"users\" WHERE (.+)$"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	ur := NewUserRepo(gormDB)

	id := int(1)
	status, err := ur.Delete(id)
	assert.NoError(t, err)
	assert.True(t, status)
}
