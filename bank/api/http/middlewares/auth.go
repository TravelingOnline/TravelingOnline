package middlewares

import (
	"errors"
	"strings"

	"github.com/onlineTraveling/bank/api/http/handlers"
	"github.com/onlineTraveling/bank/pkg/ports/clients/clients"
	"github.com/onlineTraveling/bank/pkg/valuecontext"

	"github.com/gofiber/fiber/v2"
)

func Auth(GRPCAuthClient clients.IAuthClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		if authorization == "" {
			return handlers.SendError(c, errors.New("authorization header missing"), fiber.StatusUnauthorized)
		}

		// Split the Authorization header value
		parts := strings.Split(authorization, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return handlers.SendError(c, errors.New("invalid authorization token format"), fiber.StatusUnauthorized)
		}

		//pureToken := parts[1]
		pureToken := parts[1]
		user, err := GRPCAuthClient.GetUserByToken(pureToken)
		if err != nil {
			return handlers.SendError(c, err, fiber.StatusUnauthorized)
		}

		c.Locals(valuecontext.UserClaimKey, user)

		return c.Next()
	}
}
