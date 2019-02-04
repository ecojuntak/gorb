package middlewares

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
)

func CorsMiddleware() []handlers.CORSOption {
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	var cors []handlers.CORSOption
	cors = append(cors, allowedHeaders)
	cors = append(cors, allowedMethods)
	cors = append(cors, allowedOrigins)

	return cors
}

func LoggerMidldlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Infoln(r.Method, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
