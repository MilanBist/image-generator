package utils

import (
	"errors"
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


func generateClaims(email, typeof string, timing int, userId int64) Claims{
	claims := Claims{
		Email: email,
		Type: typeof,
		RegisteredClaims: jwt.RegisteredClaims{
				Subject: strconv.FormatInt(userId, 10),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(timing))),
				IssuedAt: jwt.NewNumericDate(time.Now()),
			},
	}
	return claims
}

func GenerateTokens(userId int64, email, requirement string)(string,string,error){
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	accessSecretKey := []byte(os.Getenv("SECRET_KEY_ACCESS"))
	refreshSecretKey := []byte(os.Getenv("SECRET_KEY_REFRESH"))


	// return on the basis of the rquirement.
	switch requirement{
	case "access":
		accessTokenClaims := generateClaims(email, "access", 30, userId)
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
		accessTokenString, err := accessToken.SignedString(accessSecretKey)
		if err != nil{
			fmt.Println("[GENERATING TOKENS]: Error in generating the access claims")
			return "", "", errors.New("Error in generating the claims for access token.")
		}
		return "", accessTokenString, nil
	
	case "refresh":
		refreshTokenClaims := generateClaims(email, "refresh", 60*24*7, userId)
		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
		refreshTokenString, err := refreshToken.SignedString(refreshSecretKey)

		if err != nil {
		return "", "", err
		}
		return refreshTokenString, "", nil


	case "both":
		accessTokenClaims := generateClaims(email, "access", 30, userId)
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
		accessTokenString, err := accessToken.SignedString(accessSecretKey)
		if err != nil{
			fmt.Println("[GENERATING TOKENS]: Error in generating the access claims")
			return "", "", errors.New("Error in generating the claims for access token.")
		}

		refreshTokenClaims := generateClaims(email, "refresh", 60*24*7, userId)
		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
		refreshTokenString, err := refreshToken.SignedString(refreshSecretKey)

		if err != nil {
		return "", "", err
		}
		return refreshTokenString, accessTokenString, nil
	}
	return "","", errors.New("Wrong requirement call.")
}
