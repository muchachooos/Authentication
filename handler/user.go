package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
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

	token := uuid.NewString()

	time := time.Now()

	err := s.Storage.RegistrationUserInBD(log, pass, token, time)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	context.Status(http.StatusOK)
	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) AuthorizationHandler(context *gin.Context) {

	log, ok := context.GetQuery("login")
	if log == "" || !ok {
		context.Writer.WriteString("No login")
		return
	}

	pass, ok := context.GetQuery("password")
	if pass == "" || !ok {
		context.Writer.WriteString("No password")
		return
	}

	token := uuid.NewString()

	resultTable, isChanged, err := s.Storage.AuthorizationUserInDB(log, pass, token)
	if err != nil {
		context.Status(500)
		fmt.Println("ERROR : ", err)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if isChanged == false {
		context.Writer.WriteString("Something went wrong")
		return
	}

	if len(resultTable) == 0 {
		context.Status(404)
		context.Writer.WriteString("")
		return
	}

	jsonInByte, err := json.Marshal(resultTable)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)

}

//Check handler
