CREATE DATABASE UserDummy;
CREATE TABLE users
(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_name varchar(64),
    first_name varchar(64),
    last_name varchar(64)
);

-- Insert rows into table 'TableName'
INSERT INTO users
( -- columns to insert data into
 user_name,
 first_name,
 last_name
)
VALUES
( -- first row: values for the columns in the list above
 "Ignaciocd", "Carvajal", "Duran"
);

