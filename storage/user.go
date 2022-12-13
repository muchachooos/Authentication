package storage

import (
	"Authorization/model"
	"github.com/google/uuid"
	"time"
)

func (u *UserStorage) RegistrationUserInBD(log, pass string) error {

	token := uuid.NewString()

	time := time.Now()

	_, err := u.DataBase.Exec("INSERT INTO user (`login`, `password`, `token`, `time`) VALUES (?,?,?,?)", log, pass, token, time)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserStorage) AuthorizationUserInDB(log, pass string) ([]model.Data, bool, error) {

	token := uuid.NewString()

	time := time.Now()

	res, err := u.DataBase.Exec("UPDATE user SET `token` = ?, `time` = ? WHERE `login` = ? AND `password` = ?", token, time, log, pass)
	if err != nil {
		return nil, false, err
	}

	var resultTable []model.Data

	err = u.DataBase.Select(&resultTable, "SELECT * FROM user WHERE `login` = ? AND `password` = ?", log, pass)
	if err != nil {
		return nil, false, err
	}

	if len(resultTable) == 0 {
		return nil, false, nil
	}

	countOfChangedRows, err := res.RowsAffected()
	if err != nil {
		return nil, false, nil
	}

	if countOfChangedRows == 0 {
		return nil, false, nil
	}

	return resultTable, true, nil
}

func (u *UserStorage) CheckTokenInDB(token string) (model.CheckTokenResponse, bool, error) {

	var resultTable []model.Data

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
