package handlers

import (
	"html/template"
	"net/http"

	"github.com/cart-overflow/user-api/pkg/pb"
)

type OAuthCallback struct {
	uc pb.UserServiceClient
}

func NewOAuthCallback(uc pb.UserServiceClient) *OAuthCallback {
	return &OAuthCallback{uc}
}

// TODO: Return html errors instead of json
func (h *OAuthCallback) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	redirectState := r.URL.Query().Get("state")
	state, err := r.Cookie("state")
	if err != nil {
		// TODO: log error
		WriteWebError(w, "State parameter is required", http.StatusBadRequest)
		return
	}

	res, err := h.uc.ExchangeCode(r.Context(), &pb.ExchangeCodeRequest{
		Code:          code,
		RedirectState: redirectState,
		ClientState:   state.Value,
	})
	if err != nil {
		// TODO: log error
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
