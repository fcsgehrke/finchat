package crypt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Crypter struct {
	JWTExpirationTime time.Duration
	JWTSecretKey      string
}

func NewCrypter(jwtSecret string, jwtExpTime time.Duration) (*Crypter, error) {
	return &Crypter{
		JWTSecretKey:      jwtSecret,
		JWTExpirationTime: jwtExpTime,
	}, nil
}

func (c *Crypter) EncryptPassword(password string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPwd), nil
}

func (c *Crypter) ValidatePassword(password string, savedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(savedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

func (c *Crypter) GenerateToken(user string, id int) (string, error) {
	claims := jwt.MapClaims{
		"user": user,
		"id":   id,
		"exp":  time.Now().Add(c.JWTExpirationTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(c.JWTSecretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}
