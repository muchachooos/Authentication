CREATE TABLE `user`
(
    `id`       INT AUTO_INCREMENT PRIMARY KEY,
    `login`    VARCHAR(25) UNIQUE CHECK (`login` != '')      NOT NULL,
    `hashedPass` VARCHAR(300) UNIQUE CHECK ( `hashedPass` != '') NOT NULL,
    `token`    CHAR(36) UNIQUE,
    `time`     DATETIME
);

DROP TABLE `user`;
SHOW DATABASES;
USE auth_data;
SHOW TABLES;
SELECT *
FROM user;

# SELECT EXISTS(SELECT * FROM user WHERE time = ?)