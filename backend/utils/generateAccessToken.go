package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Claims struct {
	Email string `json:"email"`
	Type  string `json:"type"`
	jwt.RegisteredClaims
}

func GenerateTokens(userId int64, email, requirement string)(string,string,error){
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	accessSecretKey := []byte(os.Getenv("SECRET_KEY_ACCESS"))
	refreshSecretKey := []byte(os.Getenv("SECRET_KEY_REFRESH"))

	refreshTokenClaims := Claims{
		Email: email,
		Type: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: strconv.FormatInt(userId, 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	accessTokenClaims := Claims{
		Email: email,
		Type: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: strconv.FormatInt(userId, 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30* time.Minute)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	// this signing method means like same secret key performs finding and generating the token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	// create the string with the secret key
	accessTokenString, err := accessToken.SignedString(accessSecretKey)
	refreshTokenString, err2 := refreshToken.SignedString(refreshSecretKey)

	if err != nil {
		return "", "", err
	} else if err2 != nil{
		return "", "", err2
	}

	if requirement == "access"{
		return accessTokenString, "", nil
	} else if requirement == "both"{
		return accessTokenString, refreshTokenString, nil
	}
	
	// just get the refresh token
	return "", refreshTokenString, nil
}
