package http

import (
	"agency/api/pb"
	"agency/api/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateTour(svcGetter ServiceGetter[*service.TourService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req pb.TourCreateRequest

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		if err := svc.CreateTour(c.UserContext(), &req); err != nil {
			return err
		}
		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"message": "created",
		})
	}
}
