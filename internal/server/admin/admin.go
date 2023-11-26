package admin

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"syscall"
	"time"

	"github.com/amalshaji/localport/internal/server/admin/handler"
	"github.com/amalshaji/localport/internal/server/admin/service"
	"github.com/amalshaji/localport/internal/server/config"
	"github.com/amalshaji/localport/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django/v3"
)

type AdminServer struct {
	app    *fiber.App
	config *config.AdminConfig
	log    *slog.Logger
}

func New(config *config.Config, service *service.Service) *AdminServer {
	engine := django.New("./internal/server/admin/templates", ".html")
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 engine,
	})

	app.Use(logger.New())
	app.Use(recover.New())

	if !config.Admin.UseVite {
		app.Static("/", "./internal/server/admin/web/dist")
	}

	app.Static("/static", "./internal/server/admin/static", fiber.Static{
		Compress: true,
	})

	// middleware to handle authentication
	// throw api error for api requests
	// redirect to login for dashboard requests
	// finally, set user in locals

	clientPages := []string{"/connections", "/overview", "/settings", "/users", "/my-account"}
	authExcludedEndpoints := []string{"/config/validate", "/api/setting/signup"}

	// handle auth
	app.Use(func(c *fiber.Ctx) error {
		if slices.Contains(authExcludedEndpoints, c.Path()) {
			return c.Next()
		}

		token := c.Cookies("localport-session")
		user, err := service.GetUserBySession(token)

		if err != nil {
			if slices.Contains(clientPages, c.Path()) {
				return c.Redirect("/")
			} else if c.Path() == "/" {
				return c.Next()
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
			}

		} else {
			if c.Path() == "/" {
				return c.Redirect("/overview")
			}
		}
		// set user in locals
		c.Locals("user", user)
		return c.Next()
	})

	handler := handler.New(config, service)
	handler.RegisterUserRoutes(app)
	handler.RegisterConnectionRoutes(app)
	handler.RegisterGithubAuthRoutes(app)
	handler.RegisterSettingsRoutes(app)
	handler.RegisterInviteRoutes(app)
	handler.RegisterClientConfigRoutes(app)

	// server index templates for all routes
	// should be explicit?
	app.Use("*", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"UseVite":  config.Admin.UseVite,
			"ViteTags": getViteTags(),
		})
	})

	return &AdminServer{
		app:    app,
		config: &config.Admin,
		log:    utils.GetLogger(),
	}
}

func (s *AdminServer) Start() {
	s.log.Info("starting admin server", "port", s.config.ListenAddress())

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.app.Listen(s.config.ListenAddress()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Error("failed to start admin server", "error", err)
			done <- nil
		}
	}()

	<-done
	s.log.Info("stopping admin server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.app.ShutdownWithContext(ctx); err != nil {
		s.log.Error("failed to stop proxy server", "error", err)
	}
}
