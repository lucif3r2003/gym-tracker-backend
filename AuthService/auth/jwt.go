package auth

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type JWTManager struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	once       sync.Once
	initErr    error
}

func NewJWTManager() *JWTManager {
	return &JWTManager{}
}

// lazy load keys
func (j *JWTManager) initKeys() error {
	j.once.Do(func() {
		// load .env
		if err := godotenv.Load(".env"); err != nil {
			j.initErr = fmt.Errorf("cannot load .env: %w", err)
			return
		}

		privPath := os.Getenv("JWT_PRIVATE_KEY_PATH")
		if privPath == "" {
			j.initErr = fmt.Errorf("JWT_PRIVATE_KEY_PATH not set")
			return
		}
		privBytes, err := os.ReadFile(privPath)
		if err != nil {
			j.initErr = fmt.Errorf("cannot read private key: %w", err)
			return
		}
		privKey, err := jwt.ParseECPrivateKeyFromPEM(privBytes)
		if err != nil {
			j.initErr = fmt.Errorf("cannot parse private key: %w", err)
			return
		}
		j.privateKey = privKey

		pubPath := os.Getenv("JWT_PUBLIC_KEY_PATH")
		if pubPath == "" {
			j.initErr = fmt.Errorf("JWT_PUBLIC_KEY_PATH not set")
			return
		}
		pubBytes, err := os.ReadFile(pubPath)
		if err != nil {
			j.initErr = fmt.Errorf("cannot read public key: %w", err)
			return
		}
		pubKey, err := jwt.ParseECPublicKeyFromPEM(pubBytes)
		if err != nil {
			j.initErr = fmt.Errorf("cannot parse public key: %w", err)
			return
		}
		j.publicKey = pubKey
	})
	return j.initErr
}

func (j *JWTManager) GenerateAccessToken(userId string) (string, error) {
	if err := j.initKeys(); err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(j.privateKey)
}

func (j *JWTManager) GenerateRefreshToken(userId string) (string, error) {
	if err := j.initKeys(); err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(j.privateKey)
}

func (j *JWTManager) VerifyToken(tokenString string) (*jwt.Token, error) {
	if err := j.initKeys(); err != nil {
		return nil, err
	}

	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return j.publicKey, nil
	})
}
