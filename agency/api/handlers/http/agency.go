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

func GetAgency(svcGetter ServiceGetter[*service.AgencyService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		if id, err := strconv.Atoi(c.Params("id")); err == nil {
			agency, err := svc.GetAgencyByID(c.UserContext(), uint(id))
			if err != nil {
				return err

			}
			return c.Status(http.StatusFound).JSON(fiber.Map{
				"id":       agency.ID,
				"name":     agency.Name,
				"owner_id": agency.OwnerID,
			})
		}

		if id := c.Query("owner"); id != "" {
			intID, err := strconv.Atoi(id)
			if err != nil {
				return err
			}

			agency, err := svc.GetAgencyByOwnerID(c.UserContext(), uint(intID))

			if err != nil {
				return err
			}
			return c.Status(http.StatusFound).JSON(fiber.Map{
				"id":       agency.ID,
				"name":     agency.Name,
				"owner_id": agency.OwnerID,
			})
		}

		if page, err := strconv.Atoi(c.Query("page")); err == nil {
			pagesize, err := strconv.Atoi(c.Query("size"))
			if err != nil {
				pagesize = 20
			}
			agencies, err := svc.GetAll(c.UserContext(), page, pagesize)

			if err != nil {
				return err
			}

			return c.Status(http.StatusFound).JSON(fiber.Map{
				"agencies": agencies,
			})
		}
		return c.SendStatus(http.StatusNotFound)
	}
}

// func UpdateAgency(svcGetter ServiceGetter[*service.AgencyService]) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		id , err := strconv.Atoi(c.Params("id"))
// 		if err != nil {
// 			return err 
// 		}

		
// 	}
// }
