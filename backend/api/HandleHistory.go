package api

import (
	"encoding/json"
	"net/http"

	"github.com/image-generator/internal/models"
)

type HistoryData interface{
	GetHistoryDataOfUser(userId int) (string, error)
}

type History struct{
	Data 	HistoryData
}

func(h *History) HandleHistory(w http.ResponseWriter, r *http.Request){
	// get the userId from the context
	value := r.Context().Value("metaData").(models.ContextMetaData)
	userId := value.UserId


	// based on the userId get the history data from the backend server
	historyData, err := h.Data.GetHistoryDataOfUser(userId)
	if err != nil{
		response := models.Response{
			Success: false,
			Message: "Can't retrive the historical data.",
		}

		json.NewEncoder(w).Encode(&response)
	}


	response := models.Response{
		Success: true,
		Message: "Success in retrieving the historical data.",
		Data: historyData,
	}

	json.NewEncoder(w).Encode(&response)
}