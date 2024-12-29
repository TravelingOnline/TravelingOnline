package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/onlineTraveling/auth/api/service"
	"github.com/onlineTraveling/auth/pkg/adapters/storage/types"
	"github.com/onlineTraveling/auth/pkg/logger"
	"github.com/onlineTraveling/auth/protobufs"
)

func SignUp(svcGetter ServiceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req protobufs.UserSignUpFirstRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.SignUp(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, service.ErrUserCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}
func SignIn(svcGetter ServiceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req protobufs.UserSignInRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.SignIn(c.UserContext(), &req)
		if err != nil {

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}
func SignUpCodeVerification(svcGetter ServiceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req protobufs.UserSignUpSecondRequest
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.SignUpCodeVerification(c.UserContext(), &req)
		if err != nil {
			if errors.Is(err, service.ErrUserCreationValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func GetUserByID(svcGetter ServiceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		resp, err := svc.GetByID(c.UserContext(), uint(id))
		if err != nil {
			if errors.Is(err, service.ErrUserNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func Update(svcGetter ServiceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		var req types.User
		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}
		err := svc.Update(c.UserContext(), &req)
		if err != nil {
			logger.Error("error in update user", nil)
			return err
		}
		return nil
	}
}

func DeleteByID(svcGetter ServiceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		err = svc.DeleteByID(c.UserContext(), id)
		if err != nil {
			logger.Error("error in delete user", nil)
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		logger.Info("deleted user successfully", nil)
		return nil
	}
}
