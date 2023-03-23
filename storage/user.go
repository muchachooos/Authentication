package storage

import (
	"Authorization/model"
	"Authorization/utilities"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"time"
)

const tokenTTL = 15 * time.Minute

func (u *UserStorage) RegistrationUserInBD(log, pass string) error {
	hashedPass, err := utilities.GenerateHashPassword(pass)
	if err != nil {
		return err
	}

	_, err = u.DataBase.Exec("INSERT INTO user (`login`, `hashedPass`, `token`, `tokenTTL`) VALUES (?,?,?,?)", log, hashedPass, uuid.NewString(), time.Now())

	return err
}

func (u *UserStorage) AuthorizationUserInDB(log, pass string) (string, error) {
	var hashedPass string

	err := u.DataBase.Get(&hashedPass, "SELECT `hashedPass` FROM user WHERE `login` = ?", log)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", model.ErrorAuthorized
		}
		return "", err
	}

	err = utilities.CompareHashPassword(hashedPass, pass)
	if err != nil {
		return "", model.ErrorAuthorized
	}

	token := uuid.NewString()

	result, err := u.DataBase.Exec("UPDATE user SET `token` = ?, `tokenTTL` = ? WHERE `login` = ? AND `hashedPass` = ?", token, time.Now(), log, hashedPass)
	if err != nil {
		return "", err
	}

	countOfChangedRows, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if countOfChangedRows == 0 {
		return "", errors.New("failed set token")
	}

	return token, nil
}

func (u *UserStorage) CheckTokenInDB(token string) (model.CheckTokenResponse, error) {
	var user model.User

	err := u.DataBase.Get(&user, "SELECT * FROM user WHERE `token` = ?", token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.CheckTokenResponse{}, model.ErrorCheckToken
		}
		return model.CheckTokenResponse{}, err
	}

	if time.Since(user.TokenTimeToLive) > tokenTTL {
		return model.CheckTokenResponse{}, model.ErrorTokenTTLisOver
	}

	result := model.CheckTokenResponse{
		ID:    user.ID,
		Login: user.Login,
	}

	return result, nil
}
