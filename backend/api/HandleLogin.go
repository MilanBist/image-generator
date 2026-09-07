package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/image-generator/database"
	"github.com/image-generator/internal/models"
	"github.com/image-generator/utils"
)

type tokenResponse struct{
	RefreshToken string		`json:"access"`
	AccessToken string		`json:"refresh"`
}

func (srv *Server) HandleLogin(w http.ResponseWriter, r *http.Request){
	var LoginData models.Login
	json.NewDecoder(r.Body).Decode(&LoginData)
	fmt.Println("[LOGIN HANDLER]: Login Data", LoginData)


	// now validate all of the given data and add to the database
	err := utils.ValidateLoginCredentials(LoginData)


	if err != nil{
		response := models.Response{
			Success: false,
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// check loginCredentials from database
	
	// if the user doesn't exists
	id := database.LoginData(LoginData, srv.Db)
	if id == -1{
		response := models.Response{
			Success: false,
			Message: "User doesnot exists",
		}
		w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(404)
		json.NewEncoder(w).Encode(response)
		return
	}

	// provide the user the access-token and refresh-token
	accessToken,refreshToken, err := utils.GenerateTokens(int64(id), LoginData.Email, "both")

	fmt.Println("[LOGIN HANDLER]: Accesstoken and refresh token: ", refreshToken, accessToken)

	if err != nil{
		fmt.Println("[LOGIN HANDLER] ", err)
		response := models.Response{
			Success: false,
			Message: "Server error.",
		}
		w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.Response{
		Success: true,
		Message: "Successfully logged in.",
		Data: tokenResponse{
			AccessToken: accessToken,
			RefreshToken: refreshToken,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)

}