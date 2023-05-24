package main

import (
	"time"

	"github.com/14mdzk/dev_finance/internal/app/controller"
	"github.com/14mdzk/dev_finance/internal/app/repository"
	"github.com/14mdzk/dev_finance/internal/app/service"
	"github.com/14mdzk/dev_finance/internal/app/service/session"
	"github.com/14mdzk/dev_finance/internal/pkg/config"
	"github.com/14mdzk/dev_finance/internal/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	cfg    config.Config
	db     *sqlx.DB
	router *gin.Engine
}

func NewServer(cfg config.Config, db *sqlx.DB) (*Server, error) {
	server := &Server{
		cfg: cfg,
		db:  db,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
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
	router.Use(middleware.Logging(), middleware.Recovery())

	userRepository := repository.NewUserRepository(DBConn)
	registerService := service.NewRegisterService(userRepository)
	registerController := controller.NewRegisterController(registerService)

	router.POST("/register", registerController.RegisterUser)

	authRepository := repository.NewAuthRepository(DBConn)
	tokenService := service.NewTokenService(cfg.AccessTokenKey, cfg.RefreshTokenKey, cfg.AccessTokenDuration, cfg.RefreshTokenDuration)
	sessionService := session.NewSessionService(userRepository, authRepository, tokenService)
	sessionController := controller.NewSessionController(sessionService, tokenService)

	router.POST("/auth/login", sessionController.Login)
	router.GET("/auth/refresh", sessionController.RefreshToken)

	router.Use(middleware.Session(tokenService))

	router.POST("/auth/logout", sessionController.Logout)

	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router.GET("/users", middleware.Authorize("users", "read", enforcer), userController.BrowseUser)
	router.PATCH("/users/:id/change_password", middleware.Authorize("users", "change_password", enforcer), userController.ChangePasswordUser)
	router.DELETE("/users/:id", middleware.Authorize("users", "delete", enforcer), userController.DeleteUser)

	transactionTypeRepository := repository.NewTransactionTypeRepository(DBConn)
	transactionTypeService := service.NewTransactionTypeService(transactionTypeRepository)
	transactionTypeController := controller.NewTransactionTypeController(transactionTypeService)

	router.GET("/transaction_types", middleware.Authorize("transaction_types", "read", enforcer), transactionTypeController.BrowseTransactionType)
	router.POST("/transaction_types", middleware.Authorize("transaction_types", "create", enforcer), transactionTypeController.CreateTransactionType)
	router.GET("/transaction_types/:id", middleware.Authorize("transaction_types", "read", enforcer), transactionTypeController.FindTransactionType)
	router.PATCH("/transaction_types/:id", middleware.Authorize("transaction_types", "update", enforcer), transactionTypeController.UpdateTransactionType)
	router.DELETE("/transaction_types/:id", middleware.Authorize("transaction_types", "delete", enforcer), transactionTypeController.DeleteTransactionType)

	transactionCategoryRepository := repository.NewTransactionCategoryRepository(DBConn)
	transactionCategoryService := service.NewTransactionCategoryService(transactionCategoryRepository)
	transactionCategoryController := controller.NewTransactionCategoryController(transactionCategoryService)

	router.GET("/transaction_categories", middleware.Authorize("transaction_categories", "read", enforcer), transactionCategoryController.BrowseTransactionCategory)
	router.POST("/transaction_categories", middleware.Authorize("transaction_categories", "create", enforcer), transactionCategoryController.CreateTransactionCategory)
	router.GET("/transaction_categories/:id", middleware.Authorize("transaction_categories", "read", enforcer), transactionCategoryController.FindTransactionCategory)
	router.PATCH("/transaction_categories/:id", middleware.Authorize("transaction_categories", "update", enforcer), transactionCategoryController.UpdateTransactionCategory)
	router.DELETE("/transaction_categories/:id", middleware.Authorize("transaction_categories", "delete", enforcer), transactionCategoryController.DeleteTransactionCategory)

	currencyRepository := repository.NewCurrencyRepository(DBConn)
	currencyService := service.NewCurrencyService(currencyRepository)
	currencyController := controller.NewCurrencyController(currencyService)

	router.GET("/currencies", middleware.Authorize("currencies", "read", enforcer), currencyController.BrowseCurrency)
	router.POST("/currencies", middleware.Authorize("currencies", "create", enforcer), currencyController.CreateCurrency)
	router.GET("/currencies/:id", middleware.Authorize("currencies", "read", enforcer), currencyController.FindCurrency)
	router.PATCH("/currencies/:id", middleware.Authorize("currencies", "update", enforcer), currencyController.UpdateCurrency)
	router.DELETE("/currencies/:id", middleware.Authorize("currencies", "delete", enforcer), currencyController.DeleteCurrency)

	transactionRepository := repository.NewTransactionRepository(DBConn)
	transactionService := service.NewTransactionService(
		transactionRepository,
		transactionCategoryRepository,
		transactionTypeRepository,
		userRepository,
	)
	transactionController := controller.NewTransactionController(transactionService)

	router.GET("/transactions", middleware.Authorize("transactions", "read", enforcer), transactionController.BrowseTransaction)
	router.POST("/transactions", middleware.Authorize("transactions", "create", enforcer), transactionController.CreateTransaction)
	router.PATCH("/transactions/:id", middleware.Authorize("transactions", "update", enforcer), transactionController.UpdateTransaction)
	router.DELETE("/transactions/:id", middleware.Authorize("transactions", "delete", enforcer), transactionController.DeleteTransaction)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
