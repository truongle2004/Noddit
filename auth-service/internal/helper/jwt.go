package helper

import (
	requestDto "auth-service/internal/dto/request"
	"auth-service/internal/dto/response"
	"auth-service/internal/environment"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/truongle2004/service-context/core"
)

// TODO: remove hardcoded path

var (
	accessTokenSecret  *ecdsa.PrivateKey
	refreshTokenSecret *ecdsa.PrivateKey
)

var (
	accessTokenExpire  int
	refreshTokenExpire int
)

const (
	DefaultExpireToken = 24
)

func InitJwtHelper() {
	accessExpireStr := environment.AccessTokenDuration
	refreshExpireStr := environment.RefreshTokenDuration

	var err error

	accessTokenExpire, err = strconv.Atoi(accessExpireStr)
	if err != nil {
		log.Println("the access token expire is not set, using default value")
		accessTokenExpire = DefaultExpireToken
	}

	refreshTokenExpire, err = strconv.Atoi(refreshExpireStr)
	if err != nil {
		log.Println("the refresh token token expire is not set, using default value")
		refreshTokenExpire = DefaultExpireToken * 7
	}

	if accessTokenSecret, err = loadPrivateKey(environment.PrivateKeyPath); err != nil {
		log.Printf("failed to load access token private key: %v", err)
		return
	}

	if refreshTokenSecret, err = loadPrivateKey(environment.PrivateKeyPath); err != nil {
		log.Printf("failed to load refresh token private key: %v", err)
		return
	}
}

func loadPrivateKey(path string) (*ecdsa.PrivateKey, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open private key file %s: %w", path, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("failed to close private key file", err)
		}
	}(file)

	keyData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read private key file %s: %w", path, err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block – check file format and content")
	}

	if block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("unexpected PEM block type: %s", block.Type)
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ECDSA private key from %s: %w", path, err)
	}

	return privateKey, nil
}

func generateAccessToken(userId string, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    userId,
		"email": email,
		"exp":   time.Now().Add(time.Duration(accessTokenExpire) * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(accessTokenSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateRefreshToken(userId string, email string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Duration(refreshTokenExpire) * time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(refreshTokenSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func GenerateAccessTokenAndRefreshToken(userId string, email string) (string, string, error) {
	accessToken, err := generateAccessToken(userId, email)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateRefreshToken(userId, email)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func getUserDataFromToken(token string, publicKey *ecdsa.PublicKey) (string, string, error) {

	claims, err := core.ValidateToken(token, publicKey)
	if err != nil {
	}

	userId := (*claims)["userId"].(string)
	email := (*claims)["email"].(string)

	return userId, email, nil
}

func RefreshToken(c *gin.Context) {
	platform := c.GetHeader("X-Client-Platform")
	if platform != "mobile" && platform != "web" {
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.
			WithError("Invalid platform").
			WithReason("Refresh token failed"))
		return
	}

	publicKey, err := core.LoadPublicKey(environment.PublicKeyPath)
	if err != nil {
		ResponseServerError(c, "failed to load public key", err)
		return
	}

	var token string
	switch platform {
	case "mobile":
		var refreshTokenDto requestDto.RefreshTokenDto
		if err := c.ShouldBindJSON(&refreshTokenDto); err != nil {
			ResponseServerError(c, "failed to bind json", err)
			return
		}
		token = refreshTokenDto.RefreshToken

	case "web":
		token, err = c.Cookie("refreshToken")
		if err != nil {
			c.JSON(http.StatusUnauthorized, core.ErrUnauthorized.
				WithError("You need to login").
				WithReason("Invalid credential provided").
				WithDetail("error", err.Error()))
			return
		}
	default:
		c.JSON(http.StatusBadRequest, core.ErrBadRequest.WithError("Invalid platform"))
	}

	userId, email, err := getUserDataFromToken(token, publicKey)
	if err != nil {
		ResponseServerError(c, "failed to get user data from token", err)
		return
	}

	accessToken, refreshToken, err := GenerateAccessTokenAndRefreshToken(userId, email)
	if err != nil {
		ResponseServerError(c, "failed to generate access and refresh token", err)
		return
	}

	if platform == "mobile" {
		c.JSON(http.StatusOK, response.TokenResponseDto{
			Token: response.TokenDto{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
				TokenType:    "Bearer",
			},
		})
		return
	}

	c.SetCookie("refreshToken", refreshToken, refreshTokenExpire, "/", "localhost", true, true)
	c.JSON(http.StatusOK, response.TokenResponseDto{
		Token: response.TokenDto{
			AccessToken: accessToken,
			TokenType:   "Bearer",
		},
	})
}
