package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/x14n/14nserver/utils"
)

type User struct {
	ID       int
	Name     string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		RespondWithError(w, "AAA", "bad method")
		log.Println("method bad")
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, "AAA", "bad")
		log.Println("read resquest body failed", err)
		return
	}
	var user User

	errorUnmarshal := json.Unmarshal(data, &user)
	if errorUnmarshal != nil {
		RespondWithError(w, "aaa", "bad")
		log.Println("unmarshal failed", errorUnmarshal)
	}

	//查询数据库
	//需要修改的数据库对应记录的user结构体，id不能为空
	u := QueryRow(user)

	RespondWithSuccess(w, "RRRR", "ok", u)

	//生成token

	//将token和用户数据保存到redis中

	//将token返回

}

// 查询数据
func QueryRow(queryUser User) User {
	sqlStr := "SELECT id,name,passwd from User WHERE id = ?"
	row := utils.Db.QueryRow(sqlStr, queryUser.ID)
	var u User
	//然后使用Scan()方法给对应类型变量赋值，以便取出结果,注意传入的是指针
	err := row.Scan(&u.ID, &u.Name, &u.Password)
	if err != nil {
		fmt.Printf("获取数据错误, err:%v\n", err)
		return u
	}
	fmt.Printf("查询数据成功%#v", u)
	return u

}
