package storage

func (u *UserStorage) RegistrationUserInBD(log, pass, token string) error {

	_, err := u.DataBase.Exec("INSERT INTO user (login, password, token) VALUES (?,?,?)", log, pass, token)
	if err != nil {
		return err
	}

	return nil
}
