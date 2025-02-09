package auth

import (
	"strconv"
	"time"

	"github.com/dtg-lucifer/go-backend/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(secret string, userId int) (string, error) {
  expiration := time.Second * time.Duration(config.Env.JWT_EXP)

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "userId": strconv.Itoa(userId),
    "expiredAt": time.Now().Add(expiration).Unix(),
  })

  tokenStr, err := token.SignedString([]byte(secret))
  if err != nil {
    return "", err
  }

  return tokenStr, nil
}
