// internal/router/router.go
package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zkurdi45/mystack-go/internal/handler"
)

// New creates a new configured chi router.
func New(h *handler.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fs := http.FileServer(http.Dir("./web/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", h.HomeView)

	return r
}
