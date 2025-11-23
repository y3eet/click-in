package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/y3eet/click-in/internal/config"
)

func NewAuth(cfg *config.Config) {

	store := sessions.NewCookieStore([]byte(cfg.SecretKey))
	store.MaxAge(86400 * 30)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = cfg.IsProd
	store.Options.SameSite = http.SameSiteLaxMode
	if cfg.IsProd {
		// Provider callbacks are cross-site POSTs; SameSite=None keeps the goth session cookie attached.
		store.Options.SameSite = http.SameSiteNoneMode
	}

	gothic.Store = store
	goth.UseProviders(
		google.New(cfg.GoogleClientID, cfg.GoogleClientSecret, cfg.BaseUrl+"/auth/google/callback", "email",
			"profile",
			"https://www.googleapis.com/auth/userinfo.profile"),
	)
}
