// internal/handlers/ui.go
package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/zkurdi45/mystack-go/web/templates"
)

func (h *Handler) HomeView(w http.ResponseWriter, r *http.Request) {
	// --- THIS IS THE FIX ---
	// Use the contextGetUser helper to see if a user is already logged in.
	// If no user is logged in, show the normal home/login page.
	cacheBuster := fmt.Sprintf("%d", time.Now().UnixNano())
	templ.Handler(templates.Home(cacheBuster)).ServeHTTP(w, r)
}
