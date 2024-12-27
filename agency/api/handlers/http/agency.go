package http

import (
	"agency/api/service"

	"github.com/gofiber/fiber/v2"
)

func CreateAgency(svcGetter ServiceGetter[*service.AgencyService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		if err := svc.CreateAgency(c.UserContext()); err != nil {
			return err
		}

		return nil

	}
}
