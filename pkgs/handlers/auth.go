package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/ec965/rss-server/pkgs/models"
	"github.com/golang-jwt/jwt/v4"
)

var ErrInvalidToken = errors.New("invalid token")
var AuthCtxUserId = "userId"

type TokenClaims struct {
	Email  string `json:"email"`
	UserId int64  `json:"userId"`
	jwt.StandardClaims
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email  string `json:"email"`
	UserId int64  `json:"userId"`
	Token  string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var b LoginBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := models.SelectUserByEmail(context.TODO(), b.Email, b.Password)
	switch {
	case err == models.ErrUserEmailNotFound:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	case err == models.ErrUserPasswordInvalid:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		UserId: user.Id,
		Email:  user.Email,
	})

	tokenStr, err := token.SignedString(hmacSecret)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LoginResponse{
		Email:  user.Email,
		UserId: user.Id,
		Token:  tokenStr,
	})
}

type SignUpBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var b SignUpBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = models.InsertUser(context.TODO(), b.Email, b.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// AuthCtx is a middleware function that validates the auth token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) < 2 {
			http.Error(w, ErrInvalidToken.Error(), http.StatusBadRequest)
			return
		}

		tokenString := splitToken[1]

		token, err := jwt.ParseWithClaims(
			tokenString,
			&TokenClaims{},
			func(t *jwt.Token) (interface{}, error) {
				return hmacSecret, nil
			})
		if err != nil {
			http.Error(w, ErrInvalidToken.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(*TokenClaims)
		if !ok || !token.Valid {
			http.Error(w, ErrInvalidToken.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), AuthCtxUserId, claims.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
