package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (s *Server) Hand(context *gin.Context) {
	context.Writer.WriteString("Oke")
	fmt.Println("Ok")
}

//Reg handler

//Login handler

//Check handler
