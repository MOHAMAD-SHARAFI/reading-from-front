package cmd

import (
	"log"
	"project/internal/config"
	"project/internal/handlers"
	"project/internal/models"
	"project/internal/repository"
	"project/internal/usecase"
	"project/pkg/database"
)

func Run(conf *config.Config) (err error) {
	db, err := database.Connect(conf.DatabasePath, conf.Debug)
	if err != nil {
		return err
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	if err = database.Migrate(db, []any{models.User{}}); err != nil {
		return err
	}

	repo := repository.New(db.DB())

	uc := usecase.New(repo)

	return handlers.Init(conf, uc)
}
