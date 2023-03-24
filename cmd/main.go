package main

import (
	"Authorization/handler"
	"Authorization/model"
	"Authorization/storage"
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

	configInBytes, err := os.ReadFile("./configuration.json")
	if err != nil {
		panic(err)
	}

	var config model.Config

	err = json.Unmarshal(configInBytes, &config)
	if err != nil {
		panic(err)
	}

	dataBase, err := sqlx.Open("mysql", getDSN(config.DBConf))
	if err != nil {
		panic(err)
	}

	if dataBase == nil {
		panic("Database nil")
	}

	server := handler.Server{
		Storage: &storage.UserStorage{
			DataBase: dataBase,
		},
		Key: config.Key,
	}

	router.POST("/registration", server.RegistrationHandler)
	router.POST("/authorization", server.AuthorizationHandler)
	router.GET("/check_token", server.CheckTokenHandler)

	port := ":" + strconv.Itoa(config.Port)

	err = router.Run(port)
	if err != nil {
		panic(err)
	}
}

func getDSN(conf model.DBConf) string {
	return fmt.Sprintf("%s:%s@http(%s:%d)/%s?parseTime=true",
		conf.User, conf.Password, conf.Host, &conf.Port, conf.DBName)
}
