package delivery

import (
	"context"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
	"net/http"
	"strings"
)

type key string

const userIDCtx key = "userID"

func (h *Handler) userIdentity(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		authHeader := request.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(writer, models.ErrInvalidAuthorizationHeader.Error(), http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")

		if len(authHeaderParts) != 2 {
			http.Error(writer, models.ErrInvalidAuthorizationHeader.Error(), http.StatusUnauthorized)
			return
		}

		if !strings.EqualFold(authHeaderParts[0], "Bearer") {
			http.Error(writer, models.ErrInvalidAuthorizationHeader.Error(), http.StatusUnauthorized)
			return
		}

		if authHeaderParts[1] == "" {
			http.Error(writer, models.ErrInvalidAuthorizationHeader.Error(), http.StatusUnauthorized)
			return
		}

		userID, err := h.tokenManager.Parse(authHeaderParts[1])
		if err != nil {
			http.Error(writer, models.ErrInvalidAuthorizationHeader.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(request.Context(), userIDCtx, userID)

		next.ServeHTTP(writer, request.WithContext(ctx))
	}
}
