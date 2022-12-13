package storage

import (
	"Authorization/model"
	"time"
)

func (u *UserStorage) RegistrationUserInBD(log, pass, token string, time time.Time) error {

	_, err := u.DataBase.Exec("INSERT INTO user (`login`, `password`, `token`, `time`) VALUES (?,?,?,?)", log, pass, token, time)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserStorage) AuthorizationUserInDB(log, pass, token string, time time.Time) ([]model.Data, bool, error) {

	var resultTable []model.Data

	res, err := u.DataBase.Exec("UPDATE user SET `token` = ?, `time` = ? WHERE `login` = ? AND `password` = ?", token, time, log, pass)
	if err != nil {
		return nil, false, err
	}

	err = u.DataBase.Select(&resultTable, "SELECT * FROM user WHERE `login` = ? AND `password` = ?", log, pass)
	if err != nil {
		return nil, false, err
	}

	if len(resultTable) == 0 {
		return nil, false, err
	}

	countOfChangedRows, err := res.RowsAffected()
	if err != nil {
		return nil, false, err
	}

	if countOfChangedRows == 0 {
		return nil, false, err
	}

	return resultTable, true, nil
}

func (u *UserStorage) CheckTokenInDB() {

}
