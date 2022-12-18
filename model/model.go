package model

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	ID         int       `db:"id" json:"id"`
	Login      string    `db:"login" json:"login"`
	HashedPass string    `db:"hashPass" json:"hashedPass"`
	Token      string    `db:"token" json:"token"`
	Time       time.Time `db:"time" json:"time"`
}

type CheckTokenResponse struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
}

type Config struct {
	DataSourceName string `json:"dataSourceName"`
	Port           int    `json:"port"`
}

type AuthResp struct {
	Token string `json:"token"`
}

type Err struct {
	Error string `json:"error"`
}

type Request struct {
	Login string `json:"login"`
	Pass  string `json:"password"`
}
