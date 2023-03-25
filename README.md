# API Documentation

###  Регистрация пользователя
*Метод*: POST  
*URL*: /registration    
*Body:*
```json
{
"login": "some_login",
"password": "some_password"
}
```

Создание нового пользователя

###  Авторизация пользователя
*Метод*: POST  
*URL*: /authorization
*Body:*
```json
{
"login": "some_login",
"password": "some_password"
}
```

Получение токена при успешной авторизации

###  Проверка по токену
*Метод*: GET  
*URL*: /check_token?token=some_token    
*Headers:* Authorization: some_internal_key

Проверка токена
Получение ID и Login пользователя по токену

# Service Configuration

Файл конфигурации называется configuration.json и должен находиться в одной директории с исполняемым файлом

```json
{
  "port": 80,
  "auth_key": "some_key",
  "DataBase": {
    "user": "admin",
    "password": "admin",
    "host": "127.0.0.1",
    "port": 3306,
    "dataBaseName": "some_DB_name"
  }
}
```

# Database Schema

Требуемая таблицы

```sql
CREATE TABLE `user`
(
    `id`       INT AUTO_INCREMENT PRIMARY KEY,
    `login`    VARCHAR(25) UNIQUE CHECK (`login` != '') NOT NULL,
    `hashedPass` VARCHAR(300) UNIQUE CHECK ( `hashedPass` != '') NOT NULL,
    `token`    CHAR(36) UNIQUE,
    `tokenTTL`     DATETIME
);
```