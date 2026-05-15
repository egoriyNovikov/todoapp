package core_auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"time"

	core_config "github.com/egoriynovikov/todoapp/internal/core/config"
	core_error "github.com/egoriynovikov/todoapp/internal/core/error"
	"github.com/egoriynovikov/todoapp/internal/feathers/users"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user users.User) (string, error) {
	if user.Email == "" || user.Password == "" {
		return "", errors.New("email and password are required")
	}
	claims := jwt.RegisteredClaims{
		Subject:   user.ID,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(core_config.NewConfig().JWTSecret))
}

func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			core_error.WriteError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(core_config.NewConfig().JWTSecret), nil
		})
		if err != nil || !parsedToken.Valid {
			core_error.WriteError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func VerifyPassword(password string, hashedPassword string) bool {
	return sha256.Sum256([]byte(password)) == sha256.Sum256([]byte(hashedPassword))
}
