package handlers

import (
	"encoding/json"
	"finvestapi/internal/db"
	"finvestapi/internal/models"
	"finvestapi/internal/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var defaultKey = []byte("you_can_change_secret_key_in_env")

func GetSecretKey() []byte {
	envKey := os.Getenv("SECRET_KEY")
	if envKey == "" {
		return defaultKey
	}
	return []byte(envKey)
}

type contextKey string

const UserClaimsKey contextKey = "userClaims"

func generateJWT(login string) (string, error) {
	claims := jwt.MapClaims{
		"Username": login,
		"Since":    time.Now().Unix(),
		"Until":    time.Now().Add(time.Hour * 24).Unix(), // токен живёт 24 часа
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(GetSecretKey()))
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	findUser, err := db.FindUserByLogin(loginRequest.Login)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	checkPassword := utils.GetHashPassword(loginRequest.Password)
	if findUser.Password != checkPassword {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	jwtToken, err := generateJWT(loginRequest.Login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    jwtToken,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   900, // 15 минут
	})
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("access_token")
		if err != nil {
			//http.Error(w, "unauthorized", http.StatusUnauthorized)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		tokenStr := cookie.Value

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return GetSecretKey(), nil
		})

		if err != nil || !token.Valid {
			//http.Error(w, "invalid token", http.StatusUnauthorized)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func HandlerLogout(w http.ResponseWriter, r *http.Request) {
	log.Println("HandlerLogout")
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Path:     "/",
	})
	w.WriteHeader(http.StatusOK)
}
