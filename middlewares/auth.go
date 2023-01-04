package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// setup middleware
var (
	AuthClient *auth.Client
)

func InitFirebase() {
	// Use a service account
	opt := option.WithCredentialsFile("./keys/key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	AuthClient, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	//log
	fmt.Println("Firebase initialized")
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		if header == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Unauthorized"}`))
			return
		}

		idToken := header[7:]

		res, err := AuthClient.VerifyIDToken(context.Background(), idToken)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message": "Unauthorized"}`))
			return
		}

		fmt.Println(res.UID)

		next.ServeHTTP(w, r)

	})
}
