package model

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	ID              int       `db:"id" json:"id"`
	Login           string    `db:"login" json:"login"`
	HashedPass      string    `db:"hashPass" json:"hashedPass"`
	Token           string    `db:"token" json:"token"`
	TokenTimeToLive time.Time `db:"tokenTTL" json:"tokenTTL"`
}

type CheckTokenResponse struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
}

type Config struct {
	DataSourceName string `json:"dataSourceName"`
	Port           int    `json:"port"`
	Key            string `json:"auth_key"`
}

type Err struct {
	Error string `json:"error"`
}

type Request struct {
	Login string `json:"login"`
	Pass  string `json:"password"`
}

var ErrorAuthorized = errors.New("authorization unsuccessful")
var ErrorCheckToken = errors.New("no such token")
var ErrorTokenTTLisOver = errors.New("token TTL is over")
