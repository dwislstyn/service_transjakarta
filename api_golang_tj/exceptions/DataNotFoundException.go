package exceptions

import (
	"encoding/json"
	"net/http"
)

func DataNotFoundException(response http.ResponseWriter, message string, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	resultResponse := ExceptionResponse{
		ResponseCode: DataNotFoundRuleCode,
		ResponseDesc: message,
		ResponseData: data,
	}
	json.NewEncoder(response).Encode(resultResponse)
}
