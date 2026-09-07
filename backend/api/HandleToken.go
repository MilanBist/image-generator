package api

import (
	"encoding/json"
	"net/http"
	"github.com/image-generator/internal/models"
	"github.com/image-generator/utils"
)

type refreshToken struct{
	RefreshToken string		`json:"refreshToken"`
}

func(srv Server) HandleNewAccessToken(w http.ResponseWriter, r *http.Request){
	// get a new access token based on certain string
	var token refreshToken
	json.NewDecoder(r.Body).Decode(&token)

	if token.RefreshToken == ""{
		response := models.Response{
			Success: false,
			Message: "Can't get the token.",
			Data: "Attach refreshToken as Bearer <token_string>",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&response)
	}

	// else first validate the token string
	userId, email, err := utils.ValidateToken(token.RefreshToken, "refresh")
	if err != nil{
		response := models.Response{
			Success: false,
			Message: "Wrong token.",
			Data: "Attach correct token as Bearer <token_string>",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&response)
	}

	// respond with the new access token
	_, accessToken, err := utils.GenerateTokens(int64(userId), email, "access")
	if err != nil{
		response := models.Response{
			Success: false,
			Message: "Error in generating new access token.",
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&response)
	}

	response := models.Response{
		Success: true,
		Message: "Access token generated",
		Data: map[string]string{
			"accessToken": accessToken,
		},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&response)

}