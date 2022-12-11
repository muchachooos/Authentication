package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	token := uuid.NewString()

	err := s.Storage.RegistrationUserInBD(log, pass, token)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	context.Status(http.StatusOK)
	context.Writer.WriteString("Welcome to the club Body")
}

//Login handler

//Check handler

