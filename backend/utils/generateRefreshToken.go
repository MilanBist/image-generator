package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateRefreshToken(userId int64, email string)(string, error){
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	secretKey := []byte(os.Getenv("SECRET_KEY_REFRESH"))

	// define the claims
	claims := jwt.MapClaims{
		"sub":  strconv.FormatInt(userId, 10),
		"email": email,
		"type": "refresh",
		"exp":  time.Now().Add(time.Hour * 24 *7).Unix(), // expires in 10 mins
		"iat":  time.Now().Unix(),
	}

	// this signing method means like same secret key performs finding and generating the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// create the string with the secret key
	refreshTokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}
	return refreshTokenString, nil
}