package repositories

import (
	"testing"
	"time"

	"github.com/ecojuntak/gorb/models"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
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

	query := "select (.+) from users where (.+)"

	mock.ExpectQuery(query).WillReturnRows(rows)
	ur := NewUserRepo(db)

	num := int(5)
	user, err := ur.User(num)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, u, user)
}

func TestGetByAllUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "updated_at", "created_at"}).
		AddRow(1, "eco", "eco@example.com", time.Now(), time.Now())

	query := "select (.+) from users"

	mock.ExpectQuery(query).WillReturnRows(rows)
	ur := NewUserRepo(db)

	users, err := ur.Users()
	assert.NoError(t, err)
	assert.NotNil(t, users)
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
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "insert users set name=\\?, email=\\?, created_at=\\?, updated_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Name, u.Email, u.CreatedAt, u.UpdatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	ur := NewUserRepo(db)

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
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "update users set name=\\?, email=\\?, updated_at=\\? where id=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Name, u.Email, u.UpdatedAt, u.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	ur := NewUserRepo(db)

	user, err := ur.Update(1, u)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := "delete from users where id=\\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	ur := NewUserRepo(db)

	id := int(1)
	status, err := ur.Delete(id)
	assert.NoError(t, err)
	assert.True(t, status)
}
