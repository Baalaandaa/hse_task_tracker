package http

import (
	"fmt"
	"gateway/config"
	"gateway/internal"
	"gateway/internal/delivery/http/user"
	"gateway/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	loggerMDW "github.com/gofiber/fiber/v2/middleware/logger"
	recoverMDW "github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	fiber *fiber.App
	cfg   *config.Config
}

func NewHttpServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Init(a *internal.App) {
	s.fiber = fiber.New(fiber.Config{
		Immutable: true,
		AppName:   "task_service",
	})

	s.fiber.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	s.fiber.Use(recoverMDW.New())
	s.fiber.Use(loggerMDW.New())
	s.fiber.Get("/version", func(ctx *fiber.Ctx) error {
		err := ctx.JSON(s.cfg.App.Version)
		if err != nil {
			return err
		}
		return nil
	})

	userHandler := user.NewUserHandler(a.UC["user"].(*usecase.UserUC), s.cfg.JWT.Secret)
	user.MapHandlers(s.fiber.Group("/user"), userHandler)

}

func (s *Server) Start() error {
	fmt.Printf("LISTENING %s:%s\n", s.cfg.Server.Host, s.cfg.Server.Port)
	err := s.fiber.Listen(s.cfg.Server.Host + ":" + s.cfg.Server.Port)
	return err
}
