package http

import (
	"fmt"
	"os"
	"time"

	"github.com/onlineTraveling/auth/app"
	"github.com/onlineTraveling/auth/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Run(appContainer app.App, config config.ServerConfig) error {
	app := fiber.New(fiber.Config{
		AppName: "Traveling Online v0.0.1",
	})

	app.Use(TraceMiddleware())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} TraceID: ${locals:traceID}\n",
		Output: os.Stdout,
	}))
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        config.RateLimitMaxAttempt,
		Expiration: time.Duration(config.RatelimitTimePeriod) * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			xForwardedFor := c.Get("x-forwarded-for")
			if xForwardedFor == "" {
				return c.IP()
			}
			return xForwardedFor
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendString("STOP SENDING TOO MUCH REQUESTS")
		},
	}))

	api := app.Group("/api/v1")

	registerAPI(appContainer, config, api)

	// certFile := "/app/server.crt"
	// keyFile := "/app/server.key"

	// return app.ListenTLS(fmt.Sprintf(":%d", config.HttpPort), certFile, keyFile)
	return app.Listen(fmt.Sprintf(":%d", config.HttpPort))
}
func registerAPI(appContainer app.App, cfg config.ServerConfig, api fiber.Router) {
	userRouter := api.Group("/user")
	notifRouter := api.Group("/notif")
	userSvcGetter := userServiceGetter(appContainer, cfg)
	notifSvcGetter := notificationServiceGetter(appContainer, cfg)
	//user
	userRouter.Post("/sign-up", SignUp(userSvcGetter))
	userRouter.Post("/sign-in", SignIn(userSvcGetter))
	userRouter.Post("/sign-up-code-verification", SignUpCodeVerification(userSvcGetter))
	userRouter.Get("/users/:id", GetUserByID(userSvcGetter))
	userRouter.Put("/user/update", Update(userSvcGetter))
	//notif
	notifRouter.Post("/send_message", SendMessage(notifSvcGetter))
	notifRouter.Get("/unread-messages/:user_id", GetUnreadMessages(notifSvcGetter))
}
