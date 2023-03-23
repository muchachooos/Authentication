package handler

import (
	"Authorization/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (s *Server) RegistrationHandler(context *gin.Context) {
	bodyInBytes, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Read body error: " + err.Error()})
		return
	}

	err = context.Request.Body.Close()
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Close body error: " + err.Error()})
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

	err = context.Request.Body.Close()
	if err != nil {
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Close body error: " + err.Error()})
		return
	}

	var authReq model.Request

	err = json.Unmarshal(bodyInBytes, &authReq)
	if err != nil {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Unmarshal request body error: " + err.Error()})
		return
	}

	token, err := s.Storage.AuthorizationUserInDB(authReq.Login, authReq.Pass)
	if err != nil {
		if err == model.ErrorAuthorized {
			context.JSON(http.StatusUnauthorized, model.Err{Error: err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, token)
}

func (s *Server) CheckTokenHandler(context *gin.Context) {
	keyFromHeader := context.Request.Header.Get("Authorization")

	if s.Key != keyFromHeader {
		context.JSON(http.StatusUnauthorized, model.Err{Error: "Auth Key is wrong"})
		return
	}

	token, ok := context.GetQuery("token")
	if token == "" || !ok {
		context.JSON(http.StatusBadRequest, model.Err{Error: "Token is missing"})
		return
	}

	user, err := s.Storage.CheckTokenInDB(token)
	if err != nil {
		if err == model.ErrorCheckToken || err == model.ErrorTokenTTLisOver {
			context.JSON(http.StatusUnauthorized, model.Err{Error: err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, model.Err{Error: "Database error: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}
