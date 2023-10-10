package handler

import (
	"encoding/json"
	"net/http"
)

const (
	serverInterError = "服务器网络异常"
	BadRequest       = "请求参数异常"
	JsonDecodeFaild  = "json解析错误"
	//...
)

type HttpResponse struct {
	Msg  string
	Data interface{}
}

func RespondWithError(w http.ResponseWriter, message string) {
	errorResponse := &HttpResponse{
		Msg: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorResponse)
}

func RespondWithSuccess(w http.ResponseWriter, message string, data interface{}) {
	Response := HttpResponse{
		Msg:  message,
		Data: data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response)
}
