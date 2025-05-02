package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ByChanderZap/api-basics/config"
	"github.com/ByChanderZap/api-basics/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateJWT(userID string) (string, error) {
	expiration := time.Hour * time.Duration(config.Envs.JWTExpirationInHours)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Envs.JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// I dont like this approach because i dont want to pass the whole userStore.Querier to the handler
// Also i think it is not necessary to fetch the user from the database every time
// I think i can just use the userID from the token and pass it to the context

// func WithJWTAuth(handlerFunc http.HandlerFunc, store userStore.Querier) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 	}
// }

type contextKey string

const UserKey contextKey = "userID"

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromRequest(r)

		token, err := validateToken(tokenString)
		if err != nil {
			log.Println("error validating token:", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := claims["userID"].(string)
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, userId)
		r = r.WithContext(ctx)
		handlerFunc(w, r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	utils.RespondWithError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	if tokenAuth == "" || !strings.HasPrefix(tokenAuth, "Bearer ") {
		return ""
	}
	return strings.TrimPrefix(tokenAuth, "Bearer ")
}

func validateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Envs.JWTSecret), nil
	})
}

func GetUserIdFromContext(ctx context.Context) (uuid.UUID, error) {
	idString, ok := ctx.Value(UserKey).(string)
	if !ok {
		log.Println("error getting userID from the context")
		return uuid.Nil, fmt.Errorf("unauthorized")
	}
	id, err := uuid.Parse(idString)
	if err != nil {
		log.Println("error parsing userID from the context")
		return uuid.Nil, fmt.Errorf("unauthorized")
	}
	return id, nil
}
