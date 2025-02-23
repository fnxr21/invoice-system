package cmd

import (
	"os"

	"github.com/fnxr21/invoice-system/internal/database"
	"github.com/fnxr21/invoice-system/internal/router"
	"github.com/fnxr21/invoice-system/pkg/logger"
	"github.com/fnxr21/invoice-system/pkg/mysql"
	"github.com/fnxr21/invoice-system/pkg/validate"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServer() {
	dotEnv()

	e := echo.New()
	e.Validator = validate.New()
	e.HTTPErrorHandler = validate.CustomHTTPErrorHandler
	e.Use(logger.SetupLogger())
	mysql.DataBaseinit()
	database.RunMigration()

	//route
	router.RouterInit(e.Group("/api/v1"))

	//cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	//port
	PORT := os.Getenv("APP_PORT")

	// default port 500
	if PORT == "" {
		PORT = "5000"
	}

	e.Logger.Fatal(e.Start(":" + PORT))
}

func dotEnv() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
}
