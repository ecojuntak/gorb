package middlewares

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func LoggerMidldlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Infoln(r.Method, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
