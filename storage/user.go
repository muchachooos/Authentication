package storage

import (
	"Authorization/model"
	"Authorization/utilities"
	"errors"
	"github.com/google/uuid"
	"time"
)

func (u *UserStorage) RegistrationUserInBD(log, pass string) error {

	token := uuid.NewString()

	time := time.Now()

	hashedPass, err := utilities.GenerateHashPassword(pass)
	if err != nil {
		return err
	}

	_, err = u.DataBase.Exec("INSERT INTO user (`login`, `hashedPass`, `token`, `time`) VALUES (?,?,?,?)", log, hashedPass, token, time)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserStorage) AuthorizationUserInDB(log, pass string) (string, bool, error) {

	var result []string

	err := u.DataBase.Select(&result, "SELECT `hashedPass` FROM user WHERE `login` = ?", log)
	if err != nil {
		return "", false, err
	}

	if len(result) == 0 {
		return "", false, nil
	}

	err = utilities.CompareHashPassword(result[0], pass)
	if err != nil {
		return "", false, err
	}

	time := time.Now()
	token := uuid.NewString()

	res, err := u.DataBase.Exec("UPDATE user SET `token` = ?, `time` = ? WHERE `login` = ? AND `hashedPass` = ?", token, time, log, result[0])
	if err != nil {
		return "", false, err
	}

	countOfChangedRows, err := res.RowsAffected()
	if err != nil {
		return "", false, err
	}

	if countOfChangedRows == 0 {
		return "", false, errors.New("failed set token")
	}

	return token, true, nil
}

func (u *UserStorage) CheckTokenInDB(token string) (model.CheckTokenResponse, bool, error) {

	var resultTable []model.User

	err := u.DataBase.Select(&resultTable, "SELECT `id`, `login`, `time` FROM user WHERE `token` = ?", token)
	if err != nil {
		return model.CheckTokenResponse{}, false, err
	}

	if len(resultTable) == 0 {
		return model.CheckTokenResponse{}, false, nil
	}

	data := resultTable[0]

	if time.Since(data.Time) > 15*time.Minute {
		return model.CheckTokenResponse{}, false, nil
	}

	resp := model.CheckTokenResponse{
		ID:    data.ID,
		Login: data.Login,
	}
	return resp, true, nil
}
