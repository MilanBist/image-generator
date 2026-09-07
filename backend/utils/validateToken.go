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


func validateJwt(tokenString string, key []byte) (*jwt.Token, error) {
    jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
        if token.Method != jwt.SigningMethodHS256 {
            return nil, errors.New("unexpected signing method")
        }
        return key, nil
    })

    if err != nil {
		fmt.Println(err)
        fmt.Println("[VALIDATION ERROR]:", err)
        return nil, err
    }

    if !jwtToken.Valid {
        return nil, errors.New("invalid token")
    }

	fmt.Println("VALIDATION PLACE: Jwt token: ", jwtToken)
    return jwtToken, nil
}

func ValidateToken(token, tokenType string) (int,string,error){
	// get the secret key
	err := godotenv.Load()
	if err != nil{
		fmt.Println("[VALIDATING ERROR]: Error in loading the env file")
		return -1,"", err
	}

	// Now get the secret key
	var secretKey []byte
	switch tokenType{
	case "access":
		secretKey = []byte(os.Getenv("SECRET_KEY_ACCESS"))

	case "refresh":
		secretKey = []byte(os.Getenv("SECRET_KEY_REFRESH"))
	}	

	// now validate the jwt token and get the claims
	jwtToken, err := validateJwt(token, secretKey)
	if err != nil{
		return -1,"", err
	}

	// get all of the claims
	claims := jwtToken.Claims.(jwt.MapClaims)
	sub, ok := claims["sub"].(string)
	if !ok {
    	return 0,"", errors.New("invalid subject")
	}

	email, ok := claims["email"].(string)
	if !ok{
		fmt.Println("[VALIDATION]: Not correct email format.")
		return -1, "", errors.New("Invalid email format.")
	}

	userId, err := strconv.ParseInt(sub, 10, 64)
	expFloat, ok := claims["exp"].(float64)

	if !ok {
		fmt.Println("[UTILS: Token Validation] Not correct exp format")
		return -1, "", errors.New("Invalid exp format.")
	}

	if time.Now().Unix() > int64(expFloat) {
		fmt.Println("[UTILS: Token Validation] Token expired")
		return -1, "", errors.New("Token expired.")
	}

	return int(userId), email, nil
}