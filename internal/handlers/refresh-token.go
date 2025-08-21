package handlers

import "net/http"

type RefreshToken struct {
}

func NewRefreshToken() *RefreshToken {
	return &RefreshToken{}
}

func (h *RefreshToken) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
