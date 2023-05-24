package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	sqlxadapter "github.com/memwey/casbin-sqlx-adapter"
	log "github.com/sirupsen/logrus"

	"github.com/14mdzk/dev_finance/internal/pkg/config"
	"github.com/14mdzk/dev_finance/internal/pkg/db"
)

var (
	cfg      config.Config
	DBConn   *sqlx.DB
	enforcer *casbin.Enforcer
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

	opts := &sqlxadapter.AdapterOptions{
		DB:        db,
		TableName: "casbin_rules",
	}

	e, err := casbin.NewEnforcer("config/casbin_model.conf", sqlxadapter.NewAdapterFromOptions(opts))
	if err != nil {
		log.Error(fmt.Errorf("%w", err))
		log.Panic("cannot init casbin")
	}

	fmt.Println("Casbin started")
	enforcer = e

	serverLog, err := log.ParseLevel(cfg.ServerLog)
	if err != nil {
		serverLog = log.InfoLevel
	}

	log.SetLevel(serverLog)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	server, err := NewServer(cfg, DBConn)
	if err != nil {
		log.Panic("cannot initialize server")
	}

	serverPort := fmt.Sprintf(":%s", cfg.ServerPort)
	err = server.Start(serverPort)
	if err != nil {
		log.Error(fmt.Errorf("error when initializing server: %w", err))
	}
}
