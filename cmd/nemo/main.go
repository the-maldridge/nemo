package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/jinzhu/gorm"

	"github.com/the-maldridge/nemo/pkg/http"
	"github.com/the-maldridge/nemo/pkg/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	appLogger := hclog.New(&hclog.LoggerOptions{
		Name:  "nemo",
		Level: hclog.LevelFromString("DEBUG"),
	})
	hclog.SetDefault(appLogger)

	appLogger.Info("Nemo is starting up")

	db, err := gorm.Open("sqlite3", "nemo.db")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Series{})
	db.AutoMigrate(&models.Event{})
	db.AutoMigrate(&models.Question{})
	db.AutoMigrate(&models.Vote{})
	db.AutoMigrate(&models.Comment{})

	s, err := http.New(db)
	if err != nil {
		appLogger.Error("Error loading webserver", "error", err)
		return
	}
	s.Serve()

	defer db.Close()
}
