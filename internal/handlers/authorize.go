package handlers

import (
	"net/http"
	"time"

	"github.com/cart-overflow/user/api"
)

type Authorize struct {
	port       string
	userClient api.UserServiceClient
}

func NewAuthorize(port string, userClient api.UserServiceClient) *Authorize {
	return &Authorize{port, userClient}
}

func (h *Authorize) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	provider := api.OAuthProvider_GOOGLE
	redirectUrl := "http://localhost:" + h.port + "/oauth-callback"

	resp, err := h.userClient.GetOAuthUrl(r.Context(), &api.GetOAuthUrlRequest{
		Provider:    provider,
		RedirectUrl: redirectUrl,
	})
	if err != nil {
		WriteInternalServerWebError(w)
		return
	}

	cookie := http.Cookie{
		Name:     "state",
		Value:    resp.State,
		Path:     "/",
		Expires:  time.Now().Add(3 * time.Minute),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, resp.Url, http.StatusFound)
}
