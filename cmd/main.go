package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/14mdzk/dev_finance/internal/app/controller"
	"github.com/14mdzk/dev_finance/internal/app/repository"
	"github.com/14mdzk/dev_finance/internal/app/service"
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
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.Use(middleware.LogginMiddleware(), middleware.RecoveryMiddleware())

	transactionTypeRepository := repository.NewTransactionTypeRepository(DBConn)
	transactionTypeService := service.NewTransactionTypeService(transactionTypeRepository)
	transactionTypeController := controller.NewTransactionTypeController(transactionTypeService)

	router.GET("/transaction_types", transactionTypeController.BrowseTransactionType)
	router.POST("/transaction_types", transactionTypeController.CreateTransactionType)
	router.GET("/transaction_types/:id", transactionTypeController.FindTransactionType)
	router.PATCH("/transaction_types/:id", transactionTypeController.UpdateTransactionType)
	router.DELETE("/transaction_types/:id", transactionTypeController.DeleteTransactionType)

	transactionCategoryRepository := repository.NewTransactionCategoryRepository(DBConn)
	transactionCategoryService := service.NewTransactionCategoryService(transactionCategoryRepository)
	transactionCategoryController := controller.NewTransactionCategoryController(transactionCategoryService)

	router.GET("/transaction_categories", transactionCategoryController.BrowseTransactionCategory)
	router.POST("/transaction_categories", transactionCategoryController.CreateTransactionCategory)
	router.GET("/transaction_categories/:id", transactionCategoryController.FindTransactionCategory)
	router.PATCH("/transaction_categories/:id", transactionCategoryController.UpdateTransactionCategory)
	router.DELETE("/transaction_categories/:id", transactionCategoryController.DeleteTransactionCategory)

	currencyRepository := repository.NewCurrencyRepository(DBConn)
	currencyService := service.NewCurrencyService(currencyRepository)
	currencyController := controller.NewCurrencyController(currencyService)

	router.GET("/currencies", currencyController.BrowseCurrency)
	router.POST("/currencies", currencyController.CreateCurrency)
	router.GET("/currencies/:id", currencyController.FindCurrency)
	router.PATCH("/currencies/:id", currencyController.UpdateCurrency)
	router.DELETE("/currencies/:id", currencyController.DeleteCurrency)

	router.Run(fmt.Sprintf(":%s", cfg.ServerPort))
}
