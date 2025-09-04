package rpc

import (
	"net/http"

	"github.com/cart-overflow/web-bff/internal/core"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorInfo struct {
	Code   codes.Code
	Reason string
}

type ErrorCustomizer = func(info *ErrorInfo, err *core.AppError)

func MapRpcError(
	err error,
	customizers ...ErrorCustomizer,
) *core.AppError {
	st, ok := status.FromError(err)
	if !ok {
		return core.DefaultInternal(core.UnknownReason)
	}

	var errInfo *epb.ErrorInfo
	for _, d := range st.Details() {
		switch detail := d.(type) {
		case epb.ErrorInfo:
			errInfo = &detail
		}
	}

	rpcCode := st.Code()
	reason := getReason(errInfo)
	appErr := &core.AppError{
		Title:    defaultRpcErrorTitle(rpcCode),
		Message:  defaultRpcErrorMessage(rpcCode),
		HttpCode: defaultHttpCode(rpcCode),
		Reason:   reason,
	}

	for _, c := range customizers {
		c(&ErrorInfo{Code: rpcCode, Reason: reason}, appErr)
	}

	return appErr
}

func getReason(info *epb.ErrorInfo) string {
	if info != nil {
		return info.Reason
	} else {
		return core.UnknownReason
	}
}

func defaultRpcErrorTitle(c codes.Code) string {
	switch c {
	case codes.InvalidArgument:
		return "Please Check Your Input"
	case codes.Unauthenticated:
		return "Please Sign In"
	case codes.Unavailable:
		return "Service Temporarily Down"
	default:
		return core.InternalErrorTitle
	}
}

func defaultRpcErrorMessage(c codes.Code) string {
	switch c {
	case codes.InvalidArgument:
		return "Some information appears to be incorrect. Please review and try again"
	case codes.Unauthenticated:
		return "You need to be signed in to access this feature"
	case codes.Unavailable:
		return "Our service is temporarily unavailable. We're working to restore it quickly"
	default:
		return core.InternalErrorMsg
	}
}

func defaultHttpCode(c codes.Code) int {
	switch c {
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.Unavailable:
		return http.StatusBadGateway
	default:
		return http.StatusInternalServerError
	}
}
