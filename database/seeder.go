package database

import (
	"github.com/ecojuntak/gorb/repositories"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/bxcodec/faker"
	"github.com/ecojuntak/gorb/models"
)

func RunSeeder(db *gorm.DB) (err error) {
	var user = models.User{}

	userRepo := repositories.NewUserRepo(db)

	for i := 0; i < 10; i++ {
		err = faker.FakeData(&user)

		if err != nil {
			logrus.Fatalln(err)
		}

		u := userRepo.Create(user)

		logrus.Println(u)
	}

	return nil
}
