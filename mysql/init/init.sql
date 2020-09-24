DROP DATABASE IF EXISTS summer;
CREATE DATABASE summer;
USE summer;
DROP TABLE IF EXISTS `signUP`;

CREATE TABLE `signUp`(
    `id` VARCHAR(100) NOT NULL ,
    `password` VARCHAR(100) NOT NULL ,
    `token` varchar(100) NOT NULL ,
    PRIMARY KEY (`id`))
    ENGINE = InnoDB DEFAULT CHARSET=utf8;

