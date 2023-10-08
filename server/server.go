package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/x14n/14nserver/handler"
	"github.com/x14n/14nserver/utils"
)

func ServerStart() {

	DBerr := utils.InitDB()
	if DBerr != nil {
		fmt.Println("sql 启动失败")
		return
	}

	http.HandleFunc("/file/upload", handler.FileUploadHandler)
	http.HandleFunc("/file/download", handler.FileDownloadHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("服务启动失败", err)
	}
	log.Println("server start")
}
