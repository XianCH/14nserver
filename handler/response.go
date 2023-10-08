package handler

import (
	"encoding/json"
	"net/http"
)

var CodeMsgMap = map[string]string{
	"a111": "权限不足",
}

type HttpResponse struct {
	Code string
	Msg  string
	Data interface{}
}

func GetMsg(code string) string {
	codeMsg := CodeMsgMap[code]
	return codeMsg
}

func RespondWithError(w http.ResponseWriter, code string, message string) {
	errorResponse := HttpResponse{
		Code: code,
		Msg:  message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorResponse)
}

func RespondWithSuccess(w http.ResponseWriter, code string, message string, data interface{}) {
	Response := HttpResponse{
		Code: code,
		Msg:  message,
		Data: data,
	}
	w.Header().Set("COntent-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response)
}
