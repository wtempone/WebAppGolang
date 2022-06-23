package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store    *session.Store
	AUTH_KEY string = "autenticado"
	USER_ID  string = "user_id"
)

func Config() {
	router := fiber.New()
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 5,
	})

	router.Use(NewMiddleware(), cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))

	router.Post("/auth/registro", Registro)
	router.Post("/auth/login", Login)
	router.Post("/auth/logout", Logout)
	router.Get("/auth/check", VerificaSessao)
	router.Get("/usuario", ObtemUsuario)

	router.Listen(":5000")
}
