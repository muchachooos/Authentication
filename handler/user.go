package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) RegistrationHandler(context *gin.Context) {

	log, ok := context.GetQuery("login")
	if log == "" || !ok {
		context.Status(http.StatusBadRequest)
		context.Writer.WriteString("Login is missing")
		return
	}

	pass, ok := context.GetQuery("password")
	if pass == "" || !ok {
		context.Status(http.StatusBadRequest)
		context.Writer.WriteString("Password is missing")
		return
	}

	err := s.Storage.RegistrationUserInBD(log, pass)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	context.Status(http.StatusOK)
	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) AuthorizationHandler(context *gin.Context) {

	log, ok := context.GetQuery("login")
	if log == "" || !ok {
		context.Status(http.StatusBadRequest)
		context.Writer.WriteString("No login")
		return
	}

	pass, ok := context.GetQuery("password")
	if pass == "" || !ok {
		context.Status(http.StatusBadRequest)
		context.Writer.WriteString("No password")
		return
	}

	resultTable, isChanged, err := s.Storage.AuthorizationUserInDB(log, pass)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if isChanged == false {
		context.Status(http.StatusInternalServerError)
		context.Writer.WriteString("Something went wrong")
		return
	}

	if len(resultTable) == 0 {
		context.Status(http.StatusNotFound)
		context.Writer.WriteString("Wrong login or password. Try again")
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

	if len(resultTable) == 0 {
		context.Status(http.StatusNotFound)
		context.Writer.WriteString("The User is not found.")
		return
	}

	context.Status(http.StatusOK)
	context.Writer.WriteString("Welcome to the club Body")
}
