package main

import (
	"net/http"
	"os"

	"github.com/ecojuntak/gorb/middlewares"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"

	"github.com/ecojuntak/gorb/database"
	"github.com/urfave/cli"

	"github.com/spf13/viper"
)

func onError(err error, failedMessage string) {
	if err != nil {
		logrus.Errorln(failedMessage)
		logrus.Errorln(err)
	}
}

func loadConfig() (err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()

	return err
}

func runServer(db *gorm.DB) {
	r := LoadRouter(db)
	corsOption := middlewares.CorsMiddleware()

	logrus.Infoln("Server run on " + getAddress())

	http.ListenAndServe(getAddress(), handlers.CORS(corsOption[0], corsOption[1], corsOption[2])(r))
}

func main() {
	err := loadConfig()
	if err != nil {
		logrus.Errorln(err)
	}

	db, err := database.InitDatabase()
	err = db.DB().Ping()

	if err != nil {
		logrus.Errorln(err)
	}

	defer db.Close()

	cliApp := cli.NewApp()
	cliApp.Name = "GORB"
	cliApp.Version = "1.0.0"

	cliApp.Commands = []cli.Command{
		{
			Name:        "migrate",
			Description: "Run database migration",
			Action: func(c *cli.Context) error {
				err = database.Migrate(db)
				onError(err, "Failed to migrate database schema")

				return err
			},
		},
		{
			Name:        "seed",
			Description: "Run database seeder",
			Action: func(c *cli.Context) error {
				err = database.RunSeeder(db)
				onError(err, "Failed to generate fake data")

				return err
			},
		},
		{
			Name:        "start",
			Description: "Start REST API Server",
			Action: func(c *cli.Context) error {
				runServer(db)
				return nil
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		logrus.Fatalln(err)
	}
}
