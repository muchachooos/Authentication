package model

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Data struct {
	ID       int       `db:"id" json:"id"`
	Login    string    `db:"login" json:"login"`
	Password string    `db:"password" json:"password"`
	Token    string    `db:"token" json:"token"`
	Time     time.Time `db:"time" json:"time"`
}

type CheckTokenResponse struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
}

type Config struct {
	DataSourceName string `json:"dataSourceName"`
	Port           int    `json:"port"`
}
