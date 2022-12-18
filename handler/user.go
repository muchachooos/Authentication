package handler

import (
	"Authorization/model"
	"io"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) RegistrationHandler(context *gin.Context) {

	bodyInBytes, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Read body error: " + err.Error()})
		return
	}

	var regReq model.Request

	err = json.Unmarshal(bodyInBytes, &regReq)
	if err != nil {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Unmarshal request body error: " + err.Error()})
		return
	}

	err = s.Storage.RegistrationUserInBD(regReq.Login, regReq.Pass)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	context.Status(http.StatusOK)
}

func (s *Server) AuthorizationHandler(context *gin.Context) {

	bodyInBytes, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Read body error: " + err.Error()})
		return
	}

	var authReq model.Request

	err = json.Unmarshal(bodyInBytes, &authReq)
	if err != nil {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Unmarshal request body error: " + err.Error()})
		return
	}

	resp, ok, err := s.Storage.AuthorizationUserInDB(authReq.Login, authReq.Pass)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	if ok == false {
		context.JSON(http.StatusUnauthorized, model.Err{Error: " ??? "})
		return
	}

	jsonInByte, err := json.Marshal(resp)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "json creating error: " + err.Error()})
		return
	}

	context.Status(http.StatusOK)
	context.Writer.Write(jsonInByte)

}

/*
	bodyInBytes, err := io.ReadAll(context.Request.Body)
		if err != nil {
			error := model.Err{
				Error: "Bad Request",
			}

			errInByte, err := json.Marshal(error)
			if err != nil {
				return
			}
			context.Status(http.StatusBadRequest)
			context.Writer.Write(errInByte)
			context.JSON(http.StatusBadRequest, model.Err{Error: "Bad Request"})
			return
		}
*/

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
