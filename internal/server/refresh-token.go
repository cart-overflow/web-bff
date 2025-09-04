package server

import (
	"encoding/json"
	"net/http"

	"github.com/cart-overflow/user-api/pkg/pb"
	"github.com/cart-overflow/web-bff/internal/core"
	"github.com/cart-overflow/web-bff/internal/dto"
	"github.com/cart-overflow/web-bff/internal/rpc"
	"google.golang.org/grpc/codes"
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
		WriteAppError(w, core.NewError(
			"Invalid JSON",
			"The request body contains invalid JSON format. Please check your input and try again.",
			core.JsonDecodingReason,
			http.StatusBadRequest,
		))
		return
	}

	tokens, err := h.uc.TokenRefresh(
		r.Context(),
		&pb.TokenRefreshRequest{RefreshToken: body.RefreshToken},
	)
	if err != nil {
		WriteAppError(w, rpc.MapRpcError(err, customizeRefreshTokenRpcErr))
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
		WriteDefaultInternal(w, core.JsonEncodingReason)
		return
	}
}

func customizeRefreshTokenRpcErr(info *rpc.ErrorInfo, err *core.AppError) {
	if info.Code == codes.Unauthenticated {
		err.Title = "Invalid Refresh Token"
		err.Message = "The provided refresh token is not valid. Please log in again to obtain new tokens"
	}
}
