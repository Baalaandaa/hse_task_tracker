package internal

import (
	"fmt"
	"task_service/config"
	"task_service/internal/repository"
	"task_service/internal/usecase"
	"task_service/pkg/storage"
)

type App struct {
	UC  map[string]interface{}
	cfg *config.Config
}

func NewApp(cfg *config.Config) *App {
	return &App{
		UC:  make(map[string]interface{}),
		cfg: cfg,
	}
}

func (a *App) Init() error {
	db, err := storage.Connect(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		a.cfg.Postgres.Host,
		a.cfg.Postgres.Port,
		a.cfg.Postgres.User,
		a.cfg.Postgres.Password,
		a.cfg.Postgres.DBName,
		a.cfg.Postgres.SSLMode))
	if err != nil {
		return err
	}

	userRepo := repository.NewUserRepository(db)

	userUC := usecase.NewUserUC(userRepo, a.cfg.JWT.Secret)
	a.UC["user"] = userUC
	return nil
}
