package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/jpsdeveloper/sportstream-api/internal/pkg/httputil"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("panic recovered",
					"error", err,
					"stack", string(debug.Stack()),
				)
				httputil.JSONError(w, http.StatusInternalServerError,
					"INTERNAL_ERROR", "An unexpected error occurred", "")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
