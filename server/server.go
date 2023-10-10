package server

import (
	"fmt"
	"log"
	"os"
	"time"

	"net/http"

	"github.com/x14n/14nserver/handler"
	logger "github.com/x14n/14nserver/log"
	"github.com/x14n/14nserver/utils"
)

var accessLog *logger.Log

func ServerStart() {

	InitLoger()
	defer accessLog.Close()

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

func InitLoger() {
	accessLog = &logger.Log{
		EntriesNum: 1024,
		Writer:     os.Stdout,
		Interval:   1 * time.Second,
	}

	accessLog.InitLog()
	go accessLog.Loop()
}
