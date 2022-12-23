package handler

import "Authorization/storage"

type Server struct {
	Key     string
	Storage *storage.UserStorage
}
