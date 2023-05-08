package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/14mdzk/dev_finance/internal/pkg/config"
	"github.com/14mdzk/dev_finance/internal/pkg/db"
	"github.com/14mdzk/dev_finance/internal/pkg/middleware"
)

var (
	cfg    config.Config
	DBConn *sqlx.DB
)

func init() {
	configLoad, err := config.LoadConfig(".")
	if err != nil {
		log.Panic(err.Error())
	}

	cfg = configLoad

	db, err := db.ConnectDB(cfg.DBDriver, cfg.DBConnection)
	if err != nil {
		log.Panic(err.Error())
	}

	DBConn = db

	serverLog, err := log.ParseLevel(cfg.ServerLog)
	if err != nil {
		serverLog = log.InfoLevel
	}

	log.SetLevel(serverLog)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.Use(middleware.LogginMiddleware(), middleware.RecoveryMiddleware())

	router.Run(fmt.Sprintf(":%s", cfg.ServerPort))
}
