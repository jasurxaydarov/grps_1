CREATE DATABASE grps_1;


CREATE TABLE users(

    user_id     UUID            PRIMARY KEY,
    surname     VARCHAR(254)    NOT NULL,
    name        VARCHAR(124)    NOT NULL,
    password    VARCHAR(69)     NOT NULL

);