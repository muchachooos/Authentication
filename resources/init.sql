CREATE TABLE `user`
(
    `id`       INT AUTO_INCREMENT PRIMARY KEY,
    `login`    VARCHAR(25) CHECK (`login` != '')   NOT NULL,
    `password` VARCHAR(30) CHECK ( password != '') NOT NULL,
    `token`    CHAR(36),
    `time`     DATETIME
);

SHOW DATABASES;
USE auth_data;
SHOW TABLES;
SELECT *FROM user