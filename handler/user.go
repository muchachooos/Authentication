package handler

import (
	"Authorization/model"
	"fmt"
	"io"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) RegistrationHandler(context *gin.Context) {

	bodyInBytes, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	var regReq model.RegRequest

	err = json.Unmarshal(bodyInBytes, &regReq)
	if err != nil {
		fmt.Println("Error Unmarshal:", err)
		return
	}

	err = s.Storage.RegistrationUserInBD(regReq.Login, regReq.Pass)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	context.Status(http.StatusOK)
	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) AuthorizationHandler(context *gin.Context) {

	error := model.Err{
		Error: "Bad Request",
	}

	errInByte, err := json.Marshal(error)
	if err != nil {
		return
	}

	bodyInBytes, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.Status(http.StatusBadRequest)
		context.Writer.Write(errInByte)
		return
	}

	var regReq model.RegRequest

	err = json.Unmarshal(bodyInBytes, &regReq)
	if err != nil {
		fmt.Println("Error Unmarshal:", err)
		return
	}

	resultTable, ok, err := s.Storage.AuthorizationUserInDB(regReq.Login, regReq.Pass)
	if err != nil {
		fmt.Println(err)
		context.Status(http.StatusInternalServerError)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if ok == false {
		context.Status(http.StatusUnauthorized)
		context.Writer.WriteString("Something went wrong")
		return
	}

	jsonInByte, err := json.Marshal(resultTable)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		context.Writer.WriteString("json creating error")
		return
	}

	context.Status(http.StatusOK)
	context.Writer.Write(jsonInByte)
}

func (s *Server) CheckTokenHandler(context *gin.Context) {

	token, ok := context.GetQuery("token")
	if token == "" || !ok {
		context.Status(http.StatusBadRequest)
		context.Writer.WriteString("token is missing")
		return
	}

	resultTable, connect, err := s.Storage.CheckTokenInDB(token)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		context.Writer.WriteString("Something went wrong")
		return
	}

	if connect == false {
		context.Status(http.StatusUnauthorized)
		context.Writer.WriteString("Session time is over")
		return
	}

	jsonInByte, err := json.Marshal(resultTable)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		context.Writer.WriteString("json creating error")
		return
	}

	context.Status(http.StatusOK)
	context.Writer.Write(jsonInByte)
}
