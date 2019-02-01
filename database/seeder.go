package database

import (
	"database/sql"
	"log"

	"github.com/ecojuntak/gorb/repositories"

	"github.com/bxcodec/faker"
	"github.com/ecojuntak/gorb/models"
)

func RunSeeder(db *sql.DB) (err error) {
	var user = models.User{}

	userRepo := repositories.NewUserRepo(db)

	for i := 0; i < 10; i++ {
		err = faker.FakeData(&user)
		if err != nil {
			log.Println(err)
		}

		_, err := userRepo.Create(user)

		if err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}
