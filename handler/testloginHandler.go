package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RequestData struct {
	Message string
}

type Usertest struct {
	ID       int16
	Name     string
	Password string
}

func LoginHandlerTest(w http.ResponseWriter, r *http.Request) {
	var user Usertest
	if r.Method != http.MethodPost {
		log.Println(r.UserAgent(), "方法不正确")
		// 创建一个响应数据结构
		response := &HttpResponse{
			Code: "a111",         // 设置响应码
			Msg:  GetMsg("a111"), // 获取对应的响应消息
			Data: nil,            // 设置响应数据，如果有的话
		}
		responseJson, _ := json.Marshal(response)

		w.Write(responseJson)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	err2 := json.Unmarshal(data, &user)
	if err2 != nil {
		log.Println("反序列化失败", err2)
		return
	}
	fmt.Println(user.Name)

	httpResponse := &HttpResponse{
		Code: "200",
		Msg:  "OK",
		Data: user,
	}

	b, err3 := json.Marshal(httpResponse)
	if err3 != nil {
		log.Println("fuck out")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}
