package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cart-overflow/user-api/pkg/pb"
	"github.com/cart-overflow/web-bff/internal/dto"
)

type RefreshToken struct {
	uc pb.UserServiceClient
}

func NewRefreshToken(uc pb.UserServiceClient) *RefreshToken {
	return &RefreshToken{uc}
}

func (h *RefreshToken) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body dto.RefreshTokenRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		// TODO: log error
		WriteAppError(
			w,
			"Invalid JSON",
			"The request body contains invalid JSON format. Please check your input and try again.",
			"INVALID_JSON",
			http.StatusBadRequest,
		)
		return
	}

	tokens, err := h.uc.TokenRefresh(
		r.Context(),
		&pb.TokenRefreshRequest{RefreshToken: body.RefreshToken},
	)
	if err != nil {
		// TODO: log error
		// TODO: map errors from user
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := dto.RefreshTokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		// TODO: log error
		WriteInternalServerError(w)
		return
	}
}
