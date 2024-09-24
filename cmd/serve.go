package cmd

import (
	"clean-arch/config"
	"clean-arch/infra/conn"
	"clean-arch/infra/logger"
	"clean-arch/internal/http/controllers"
	httpRoutes "clean-arch/internal/http/routes"
	httpServer "clean-arch/internal/http/server"
	"clean-arch/internal/repositories/db"
	authservice "clean-arch/internal/services/auth"
	txservice "clean-arch/internal/services/transaction"
	userservice "clean-arch/internal/services/user"
	"context"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	// base context
	baseContext := context.Background()

	logger.NewLogClient(config.App().LogLevel)
	lc := logger.Client()
	dbClient := conn.Db()
	_ = dbClient

	dbRepo := db.NewRepository(dbClient, &lc)

	txsvc := txservice.NewDBTransaction(lc, dbRepo)
	authsvc := authservice.NewAuthService(dbRepo)
	usersvc := userservice.NewUserService(dbRepo)

	_ = txsvc

	// HttpServer
	var HttpServer = httpServer.New()

	authController := controllers.NewAuthController(
		baseContext,
		authsvc,
	)

	userController := controllers.NewUserController(
		baseContext,
		usersvc,
	)

	var Routes = httpRoutes.New(
		HttpServer.Echo,
		authController,
		userController,
	)

	// Spooling
	Routes.Init()
	HttpServer.Start(lc)
}
