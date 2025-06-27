package app

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/radahn42/onetime-note/internal/config"
	"github.com/radahn42/onetime-note/internal/handler"
	"github.com/radahn42/onetime-note/internal/service"
	"github.com/radahn42/onetime-note/internal/storage"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	cfg    *config.Config
	logger *zap.Logger
	server *http.Server
}

func New(cfg *config.Config) *App {
	logger, _ := zap.NewProduction()

	redisStore := storage.NewRedis(
		cfg.Redis.Addr,
		cfg.Redis.Password,
	)
	svc := service.NewNoteService(redisStore)
	h := handler.NewNoteHandler(svc)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/notes", func(r chi.Router) {
		r.Post("/", h.Create)
		r.Get("/{id}", h.Get)
	})

	srv := &http.Server{
		Addr:    cfg.App.Addr,
		Handler: r,
	}

	return &App{
		cfg:    cfg,
		logger: logger,
		server: srv,
	}
}

func (a *App) Run() error {
	go func() {
		a.logger.Info("starting server", zap.String("addr", a.cfg.App.Addr))
		err := a.server.ListenAndServe()
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				a.logger.Fatal("server error", zap.Error(err))
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	a.logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return a.server.Shutdown(ctx)
}
