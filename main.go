package main

import (
	"Authorization/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"
)

func main() {
	router := gin.Default()

	var conf model.Config

	byte, err := os.ReadFile("./configuration.json")
	if err != nil {
		fmt.Println("Error Read File:", err)
		return
	}

	err = json.Unmarshal(byte, &conf)
	if err != nil {
		fmt.Println("Error Unmarshal:", err)
		return
	}

	_, err = sqlx.Open("mysql", conf.DataSourceName)
	if err != nil {
		panic(err)
		return
	}

	port := ":" + strconv.Itoa(conf.Port)

	router.GET("", Hand)

	err = router.Run(port)
	if err != nil {
		panic(err)
		return
	}
}

func Hand(context *gin.Context) {
	context.Writer.WriteString("Oke")
	fmt.Println("Ok")
}
