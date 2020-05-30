package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ericsotack/roommate-api/internal/config"
	"github.com/ericsotack/roommate-api/pkg/db"
	"github.com/ericsotack/roommate-api/pkg/router"
	log "github.com/sirupsen/logrus"
)

const configFile = "roommate-api.config.toml"

func main() {
	log.Info(fmt.Sprintf("Reading %s for config...", configFile))
	conf, err := config.New(configFile)
	if err != nil {
		log.WithError(err).Fatal("Failed to read in config file. Shutting down...")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbData := os.Getenv("DB_DATA")

	log.Info("Connecting to database...")
	sess, err := db.InitDB(dbUser, dbPass, dbHost, dbData)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to database. Shutting down...")
	}
	defer sess.Close()

	if conf.NEW_DB {
		if err = db.VerifyDB(sess, conf.Data.LISTS); err != nil {
			log.WithError(err).Fatal("Error matching database schema. Shutting down...")
		}
	}

	log.Info("Starting route handler...")
	routeHandler := router.NewRouter(sess)

	log.Info(fmt.Sprintf("Starting HTTP server on port %s...", conf.PORT))
	err = http.ListenAndServe(fmt.Sprintf(":%s", conf.PORT), routeHandler)
	log.WithError(err).Error("Error serving requests.")
}