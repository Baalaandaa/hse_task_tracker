package user

import "github.com/gofiber/fiber/v2"

func MapHandlers(r fiber.Router, h *UserHandler) {
	r.Post("/sign_in", h.SignIn())
	r.Post("/sign_up", h.SignUp())
	r.Put("/", h.VerifyTokenMDW(), h.UpdateInfo())
}
