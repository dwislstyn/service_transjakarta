package exceptions

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(response http.ResponseWriter, message string, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	resultResponse := ExceptionResponse{
		ResponseCode: SuccessRuleCode,
		ResponseDesc: message,
		ResponseData: data,
	}
	json.NewEncoder(response).Encode(resultResponse)
}
