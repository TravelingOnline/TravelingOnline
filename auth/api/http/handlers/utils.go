package handlers

import (
	"context"
)

// func userClaims(ctx *fiber.Ctx) *jwt.UserClaims {
// 	if u := ctx.Locals("user"); u != nil {
// 		userClaims, ok := u.(*jwt2.Token).Claims.(*jwt.UserClaims)
// 		if ok {
// 			return userClaims
// 		}
// 	}
// 	return nil
// }

type ServiceGetter[T any] func(context.Context) T
