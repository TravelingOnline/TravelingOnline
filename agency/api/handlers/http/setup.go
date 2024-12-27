package http

import (
	"agency/app"
	"agency/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	router := fiber.New()

	api := router.Group("/api/", setUserContext)

	registerAgencyAPI(appContainer, api)

	return router.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func registerAgencyAPI(appContainer app.App, router fiber.Router) {

	agencySvcGetter := agencyServiceGetter(appContainer)

	router.Post("/agency", setTransaction(appContainer.DB()), CreateAgency(agencySvcGetter))
	router.Get("/agency", setTransaction(appContainer.DB())) // TODO: FINISH SERVICE
	router.Patch("/agency/:id", setTransaction(appContainer.DB()))
	router.Delete("/agency/:id", setTransaction(appContainer.DB()))

}
