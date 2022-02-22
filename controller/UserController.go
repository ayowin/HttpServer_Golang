package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"http-server-golang/db"
)

type UserControllerS struct {
	Users func(context *gin.Context)
	Login func(context *gin.Context)
}

var UserController UserControllerS

func init() {

	UserController.Users = func(context *gin.Context) {
		rows, _ := db.Interactor.Query("select * from user;")
		if rows == nil {
			context.String(http.StatusOK, "none users")
			return
		}

		response := ""
		for rows.Next() {
			id := ""
			username := ""
			password := ""
			updateTime := ""

			rows.Scan(&id, &username, &password, &updateTime)
			response = fmt.Sprintf("{'id':'%s','username':'%s','password':'%s','updateTime':'%s'}\n%s",
				id,
				username,
				password,
				updateTime,
				response)
		}

		context.String(http.StatusOK, response)
	}

	UserController.Login = func(context *gin.Context) {
		response := ""

		requestData := make(map[string]interface{})
		context.BindJSON(&requestData)
		username := requestData["username"]
		password := requestData["password"]

		if username != nil && password != nil {
			sql := fmt.Sprintf("select * from user where username='%s' and password='%s';",
				username,
				password)
			rows, _ := db.Interactor.Query(sql)
			if rows == nil || !rows.Next() {
				response = "username or password incorrect"
			} else {
				response = "login success"
			}

		}

		context.String(http.StatusOK, response)
	}

}
