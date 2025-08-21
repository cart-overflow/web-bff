package handlers

import (
	"html/template"
	"net/http"

	"github.com/cart-overflow/user/api"
)

type OAuthCallback struct {
	userClient api.UserServiceClient
}

func NewOAuthCallback() *OAuthCallback {
	return &OAuthCallback{}
}

// TODO: Return html errors instead of json
func (h *OAuthCallback) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	redirectState := r.URL.Query().Get("state")
	state, err := r.Cookie("state")
	if err != nil {
		WriteWebError(w, "State parameter is required", http.StatusBadRequest)
		return
	}

	res, err := h.userClient.ExchangeCode(r.Context(), &api.ExchangeCodeRequest{
		Code:          code,
		RedirectState: redirectState,
		ClientState:   state.Value,
	})
	if err != nil {
		WriteInternalServerWebError(w)
		return
	}

	tokens := map[string]string{
		"access_token":  string(res.AccessToken),
		"refresh_token": string(res.RefreshToken),
	}

	// TODO: Set safe target origin ('*')
	tmplRaw := "<!doctype html><html><script>window.opener?.postMessage({{.}}, '*')</script></html>"
	tmpl := template.Must(template.New("post-tokens").Parse(tmplRaw))

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, tokens)
}
