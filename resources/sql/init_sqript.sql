CREATE TABLE `user`
(
    `id`       INT AUTO_INCREMENT PRIMARY KEY,
    `login`    VARCHAR(25) UNIQUE CHECK (`login` != '')      NOT NULL,
    `hashedPass` VARCHAR(300) UNIQUE CHECK ( `hashedPass` != '') NOT NULL,
    `token`    CHAR(36) UNIQUE,
    `tokenTTL`     DATETIME
);
