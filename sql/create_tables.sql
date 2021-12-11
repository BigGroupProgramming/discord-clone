CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL,
    name varchar(250) NOT NULL,
    password varchar(250) NOT NULL,
    PRIMARY KEY (id)
);