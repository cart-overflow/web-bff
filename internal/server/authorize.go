package server

import (
	"net/http"
	"time"

	"github.com/cart-overflow/user-api/pkg/pb"
)

type Authorize struct {
	port string
	uc   pb.UserServiceClient
}

func NewAuthorize(port string, uc pb.UserServiceClient) *Authorize {
	return &Authorize{port, uc}
}

func (h *Authorize) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	provider := pb.OAuthProvider_google
	// TODO: use real host
	redirectUrl := "http://localhost:" + h.port + "/oauth-callback"

	resp, err := h.uc.GetOAuthUrl(r.Context(), &pb.GetOAuthUrlRequest{
		Provider:    provider,
		RedirectUrl: redirectUrl,
	})
	if err != nil {
		// TODO: log error
		WriteDefaultInternalWeb(w)
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
