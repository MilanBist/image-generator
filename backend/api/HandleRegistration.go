package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/image-generator/internal/models"
	"github.com/image-generator/utils"
)

type RegisterStore interface{
	CheckPastInitialization(email string) (bool, int)
	RegisterUser(credentials models.Register) (int, error)
}

// token service is written in login handler
type RegisterHandler struct{
	Store 	RegisterStore
	Token 	TokenService
}

func (rh *RegisterHandler) HandleRegister(w http.ResponseWriter, r *http.Request){
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
	checkUserExistence, id := rh.Store.CheckPastInitialization(registerData.Email)
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
	id, err = rh.Store.RegisterUser(registerData)
	if err != nil{
		response := models.Response{
			Success: false,
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// provide the user the access-token and refresh-token
	refreshToken,accessToken, err := rh.Token.GenerateTokens(int64(id), registerData.Email, "both")


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
		Message: "Successfully registered in.",
		Data: tokenResponse{
			AccessToken: accessToken,
			RefreshToken: refreshToken,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}