package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/image-generator/database"
	"github.com/image-generator/internal/models"
	"github.com/image-generator/utils"
)

func (srv *Server) HandleRegister(w http.ResponseWriter, r *http.Request){
	var registerData models.Register
	json.NewDecoder(r.Body).Decode(&registerData)


	// now validate all of the given data and add to the database
	err := utils.ValidateRegisterCredentials(registerData)


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

	
	// if the user already exists
	checkUserExistence, id := database.UserExistence(registerData.Email, srv.Db)
	if checkUserExistence == true{
		response := models.Response{
			Success: false,
			Message: "User already exists",
		}
		w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(response)
		return
	}

	// add the user to the database


	// provide the user the access-token and refresh-token
	accessToken,refreshToken, err := utils.GenerateTokens(int64(id), registerData.Email, "both")


	if err != nil{
		fmt.Println("[REGISTER HANDLER] ", err)
		response := models.Response{
			Success: false,
			Message: "Server error.",
		}
		w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// add the users to the database
	


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