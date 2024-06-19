package user

import (
	"context"
	"fmt"
	"gateway/internal/model/user"
	"gateway/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type UserHandler struct {
	uc        *usecase.UserUC
	jwtSecret string
}

func NewUserHandler(uc *usecase.UserUC, jwtSecret string) *UserHandler {
	return &UserHandler{uc: uc, jwtSecret: jwtSecret}
}

func (h *UserHandler) VerifyTokenMDW() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Println(ctx.Get("token"))
		token, err := jwt.Parse(ctx.Get("token"), func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				_, err := ctx.Status(fiber.StatusUnauthorized).Writef("Unauthorized")
				if err != nil {
					return nil, err
				}
			}
			return h.jwtSecret, nil
		})
		data := token.Claims.(jwt.MapClaims)
		fmt.Println(data["username"].(string))
		uid, err := uuid.Parse(data["username"].(string))
		if !token.Valid || err != nil {
			_, err = ctx.Status(fiber.StatusUnauthorized).Writef("Unauthorized")
			if err != nil {
				return err
			}
		}
		ctx.SetUserContext(context.WithValue(ctx.UserContext(), "uid", uid))
		return ctx.Next()
	}
}

func (h *UserHandler) SignIn() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body user.SignInInput
		if err := ctx.BodyParser(&body); err != nil {
			return err
		}
		resp, err := h.uc.SignIn(ctx.UserContext(), &body)
		if err != nil {
			return err
		}
		return ctx.Status(200).JSON(resp)
	}
}

func (h *UserHandler) SignUp() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body user.SignUpInput
		if err := ctx.BodyParser(&body); err != nil {
			return err
		}
		resp, err := h.uc.SignUp(ctx.UserContext(), &body)
		if err != nil {
			return err
		}
		return ctx.Status(200).JSON(resp)
	}
}

func (h *UserHandler) UpdateInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body user.UpdateInfoInput
		if err := ctx.BodyParser(&body); err != nil {
			return err
		}
		body.UserId = ctx.UserContext().Value("uid").(uuid.UUID)
		resp, err := h.uc.UpdateInfo(ctx.UserContext(), &body)
		if err != nil {
			return err
		}
		return ctx.Status(200).JSON(resp)
	}
}
