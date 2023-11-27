package webapp

import (
	"monban/database"
	"monban/webapp/handler"
	monbanMiddleware "monban/webapp/middleware"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupWebApp(db *database.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Get("/", handler.GetHomeHandler)
	r.Get("/apikey", monbanMiddleware.Transaction(db, handler.GetAPIKeysHandler))
	return r
}
