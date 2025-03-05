package auth

import (
	"api/cmd/internal/helpers"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type JWTConfig struct {
	Secret         string
	OauthConf      *oauth2.Config
	ExpirationTime time.Duration
}

func NewJWTConfig() *JWTConfig {
	secret := os.Getenv("JWT_SECRET")
	hoursToExpire := os.Getenv("JWT_EXPIRATION_TIME_IN_HOURS")

	hoursToExpireInt, err := strconv.Atoi(hoursToExpire)
	if err != nil {
		log.Fatal("failed to convert JWT_EXPIRATION_TIME_IN_HOURS to int")
	}

	return &JWTConfig{
		Secret: secret,
		OauthConf: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
			Scopes:       []string{"email", "profile"},
			Endpoint:     google.Endpoint,
		},
		ExpirationTime: time.Hour * time.Duration(hoursToExpireInt),
	}
}

type CreateTokenParams struct {
	ID       int32
	Username string
	Email    string
}

func (j *JWTConfig) CreateToken(userParams CreateTokenParams) (string, error) {
	claims := j.createClaims(userParams)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return j.signToken(token)
}

func (j *JWTConfig) ValidateToken(tokenString string) error {
	token, err := j.parseToken(tokenString)
	if err != nil {
		return err
	}
	return j.checkExpiration(token)
}

func (j *JWTConfig) RefreshToken(tokenString string) (string, error) {
	userParams := CreateTokenParams{}
	token, err := j.parseToken(tokenString)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to extract claims from token")
	}

	userParams.Username = claims["username"].(string)
	userParams.Email = claims["email"].(string)

	return j.CreateToken(userParams)
}

func (j *JWTConfig) createClaims(userParams CreateTokenParams) jwt.MapClaims {
	return jwt.MapClaims{
		"id":       userParams.ID,
		"username": userParams.Username,
		"email":    userParams.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
		"nbf":      time.Now().Unix(),
	}
}

func (j *JWTConfig) signToken(token *jwt.Token) (string, error) {
	tokenString, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return tokenString, nil
}

func (j *JWTConfig) checkExpiration(token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("failed to extract claims from token")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return fmt.Errorf("exp claim is not present in token")
	}

	if time.Unix(int64(exp), 0).Before(time.Now()) {
		return fmt.Errorf("token has expired")
	}

	return nil
}

func (j *JWTConfig) parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func (j *JWTConfig) ExtractClaims(r *http.Request) (jwt.MapClaims, error) {
	cookie, err := helpers.GetCookie(r, "jwt")
	if err != nil {
		log.Println("failed to get cookie from request", err)
		return nil, err
	}

	token, err := j.parseToken(cookie.Value)
	if err != nil {
		log.Println("failed to parse token from cookie", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to extract claims from token")
	}

	return claims, nil
}

func (j *JWTConfig) ExtractUserID(r *http.Request) (int32, error) {
	claims, err := j.ExtractClaims(r)
	if err != nil {
		log.Println("failed to extract claims from request", err)
		return 0, err
	}

	id, ok := claims["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("failed to extract id from claims cookie")
	}

	return int32(id), nil
}

func (j *JWTConfig) ExtractIDFromToken(token string) (int32, error) {
	claims, err := j.ExtractClaimsFromToken(token)
	if err != nil {
		return 0, err
	}

	id, ok := claims["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("failed to extract id from claims")
	}

	return int32(id), nil
}

func (j *JWTConfig) ExtractClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := j.parseToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to extract claims from token")
	}

	return claims, nil
}
