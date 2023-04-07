package main

import (
	"Authorization/configuration"
	"Authorization/handler"
	"Authorization/model"
	"Authorization/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
)

const configPath = "configuration.json"

func main() {
	config := configuration.GetConfig(configPath)

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

	router := gin.Default()

	router.POST("/registration", server.RegistrationHandler)
	router.POST("/authorization", server.AuthorizationHandler)
	router.GET("/check_token", server.CheckTokenHandler)

	port := ":" + strconv.Itoa(config.Port)

	err = router.Run(port)
	if err != nil {
		panic(err)
	}
}

func getDSN(cfg model.DBConf) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.User, cfg.Password, cfg.Host, cfg.DBPort, cfg.DBName)
}
