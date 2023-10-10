package handler

import (
	"encoding/json"
	"fmt"
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
		RespondWithError(w, "请求方法错误")
		log.Println("method bad")
		return
	}

	var user User

	err := JsonDecode(r, &user)
	if err != nil {
		RespondWithError(w, "")
		log.Println(err)
		return
	}

	err = QueryRow(user)
	if err != nil {
		RespondWithError(w, "查不到")
		log.Println(err)
		return
	}

	//生成token
	token, err := utils.GetToken(user.Name, 5)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	responseToken := map[string]string{"token": token}

	json.Marshal(responseToken)

	//将token返回

	RespondWithSuccess(w, "ok", responseToken)
}

// 查询数据
func QueryRow(queryUser User) error {
	sqlStr := "SELECT id, name, passwd from User WHERE id = ?"
	row := utils.Db.QueryRow(sqlStr, queryUser.ID)
	var u User

	err := row.Scan(&u.ID, &u.Name, &u.Password)
	if err != nil {
		fmt.Printf("获取数据错误, err:%v\n", err)
		// 返回一个空的 User 结构体和一个错误
		return err
	}

	fmt.Printf("查询数据成功%#v", u)

	// 返回查询结果和 nil 错误
	return nil
}

func JsonDecode(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(data)
}
