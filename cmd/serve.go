package cmd

import (
	"github.com/cool-service-go/config"
	"github.com/cool-service-go/database"
	"github.com/cool-service-go/handler"
	"github.com/cool-service-go/repository"
	"github.com/cool-service-go/route"
	"github.com/cool-service-go/service"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs the application and serves the APIs",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	cfg, err := config.LoadConfig(cfgFile)
	if err != nil {
		panic(err)
	}

	db, err := database.InitDB(cfg)
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(cfg)
	handler := handler.NewHandler(repo, service)

	e := echo.New()

	route.InitRoutes(e, handler)

	e.Logger.Fatal(e.Start(":5060"))
}
