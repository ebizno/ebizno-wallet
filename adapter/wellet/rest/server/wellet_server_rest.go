package server

import (
	"fmt"
	"github.com/Paulo-Lopes-Estevao/e-bizno/adapter/wellet/factory"
	"github.com/Paulo-Lopes-Estevao/e-bizno/adapter/wellet/infrastructure/database/postgres"
	"github.com/Paulo-Lopes-Estevao/e-bizno/adapter/wellet/rest/controllers"
	"github.com/Paulo-Lopes-Estevao/e-bizno/adapter/wellet/rest/routes"
	"github.com/Paulo-Lopes-Estevao/e-bizno/usecase/system"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

const PORT = 2004

func WelletServerRest() {
	server := echo.New()
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	dbWellet := postgres.ConnectedInPostgres()
	welletFactory := factory.NewWelletFactoryRepository(dbWellet).FactoryRepositoryWellet()
	welletUsecaseSystem := system.NewWelletUsecaseSystem(welletFactory)
	welletController := controllers.NewWelletController(welletUsecaseSystem)
	routes.WelletRouter(server, welletController)
	log.Printf("Server Rest Wellet running... on [::]:%d", PORT)
	if err := server.Start(fmt.Sprintf(":%d", PORT)); err != nil {
		log.Println("Not Running Server A...", err.Error())
	}

}
