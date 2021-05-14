CREATE DATABASE xfers;

USE xfers;

CREATE TABLE `kurs` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `name` varchar(3) NOT NULL,
                        PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE `e_rate` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `kurs_id` int NOT NULL,
                          `buy` decimal(18,2) NOT NULL,
                          `sell` decimal(18,2) NOT NULL,
                          `valid_date` date NOT NULL,
                          PRIMARY KEY (id),
                          FOREIGN KEY (kurs_id) REFERENCES kurs(id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `tt_counter` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `kurs_id` int NOT NULL,
                              `buy` decimal(18,2) NOT NULL,
                              `sell` decimal(18,2) NOT NULL,
                              `valid_date` date NOT NULL,
                              PRIMARY KEY (id),
                              FOREIGN KEY (kurs_id) REFERENCES kurs(id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `bank_notes` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `kurs_id` int NOT NULL,
                              `buy` decimal(18,2) NOT NULL,
                              `sell` decimal(18,2) NOT NULL,
                              `valid_date` date NOT NULL,
                              PRIMARY KEY (id),
                              FOREIGN KEY (kurs_id) REFERENCES kurs(id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
