package http

import (
	"agency/api/pb"
	"agency/api/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateAgency(svcGetter ServiceGetter[*service.AgencyService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req pb.AgencyCreateRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		if err := svc.CreateAgency(c.UserContext(), &req); err != nil {
			return err
		}

		return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "created"})

	}
}

func DeleteAgency(svcGetter ServiceGetter[*service.AgencyService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return err
		}

		err = svc.DeleteAgency(c.UserContext(), uint(id))

		if err != nil {
			return err
		}

		return nil
	}
}
