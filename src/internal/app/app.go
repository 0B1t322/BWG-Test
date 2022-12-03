package app

import (
	"context"
	"net/http"

	_ "github.com/0B1t322/BWG-Test/docs"
	balanceadt "github.com/0B1t322/BWG-Test/internal/adapters/balance"
	transactionsadt "github.com/0B1t322/BWG-Test/internal/adapters/transactions"
	"github.com/0B1t322/BWG-Test/internal/config"
	transactionshttp "github.com/0B1t322/BWG-Test/internal/controllers/http/transactions"
	userhttp "github.com/0B1t322/BWG-Test/internal/controllers/http/users"
	balancesrv "github.com/0B1t322/BWG-Test/internal/domain/balance/service"
	transactionsrv "github.com/0B1t322/BWG-Test/internal/domain/transaction/service"
	usersrv "github.com/0B1t322/BWG-Test/internal/domain/user/service"
	dal "github.com/0B1t322/BWG-Test/internal/infrastructure/dal/gorm"
	balancerepo "github.com/0B1t322/BWG-Test/internal/infrastructure/dal/gorm/balance"
	transactionrepo "github.com/0B1t322/BWG-Test/internal/infrastructure/dal/gorm/transactions"
	userrepo "github.com/0B1t322/BWG-Test/internal/infrastructure/dal/gorm/user"
	"github.com/0B1t322/BWG-Test/internal/services/transactionqueue"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	cfg config.Config
}

func NewApp(cfg config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Run() error {
	db, err := dal.ConnectToPostgreSQL(a.cfg.DB.PostgreSQLDSN)
	if err != nil {
		return err
	}
	if err := dal.CreateModels(db); err != nil {
		return err
	}

	userRepo := userrepo.NewPGUserRepository(db)
	balanceRepo := balancerepo.NewPGBalanceRepository(db)
	transactionRepo := transactionrepo.NewPGTransactionsRepository(db)

	userService := usersrv.NewUserService(userRepo)
	balanceService := balancesrv.NewBalanceService(
		balanceRepo,
		balanceadt.NewUserGetter(userService),
	)

	transactionService := transactionsrv.NewTransactionService(
		transactionRepo,
		transactionsadt.NewBalanceService(balanceService),
	)

	transactionQueue := transactionqueue.NewTransactionQueue(
		transactionService,
	)

	// On start we need to load all transactions from DB and put them into queue
	transactions, err := transactionRepo.GetAllTransactionsThatNotExecuted(context.Background())
	if err != nil {
		return err
	}

	transactionQueue.AddTransactions(transactions...)
	transactionQueue.ExecuteTransactions()

	eng := gin.New()

	root := eng.Group("/api")
	{
		root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		// init redirect to swagger
		root.GET(
			"/swagger",
			func(c *gin.Context) {
				c.Redirect(http.StatusMovedPermanently, "/api/swagger/index.html")
			},
		)

		userhttp.NewUserController(userService, balanceService).Build(root)
		transactionshttp.NewTransactionsController(transactionService, transactionQueue).Build(root)

	}

	logrus.Info("Start http server on port: ", a.cfg.App.Port)
	return eng.Run(":" + a.cfg.App.Port)
}
