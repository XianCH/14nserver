package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const uploadPath = "./upload"

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(100)
	if err != nil {
		fmt.Println("over mixsize", err)
		return
	}
	// fmt.Printf("http request:%#v\n", *r)
	mForm := r.MultipartForm

	for k := range mForm.File {
		file, fileHeader, err := r.FormFile(k)
		if err != nil {
			fmt.Println("invoke formfile error", err)
			return
		}
		defer file.Close()
		log.Printf("the upload file : name:[%s],size[%d],header[%#v]\n", fileHeader.Filename, fileHeader.Size, fileHeader.Header)

		localFileName := uploadPath + "/" + fileHeader.Filename
		out, err := os.Create(localFileName)
		if err != nil {
			fmt.Printf("create localpath file %s for warning", err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)

		if err != nil {
			fmt.Printf("copy file error:%s", err)
		}
		fmt.Printf("upload file %s is ok", fileHeader.Filename)
	}
}

func FileDownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed for file download", http.StatusMethodNotAllowed)
		return
	}

	// 获取客户端传递的文件名
	fileName := r.FormValue("filename")
	if fileName == "" {
		http.Error(w, "Please provide a valid filename", http.StatusBadRequest)
		return
	}

	// 构建文件的完整路径
	filePath := uploadPath + "/" + fileName

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to open file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 设置响应头，指定文件的 Content-Disposition
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")

	// 将文件内容写入响应主体
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to write file content to response: %v", err), http.StatusInternalServerError)
		return
	}
}
