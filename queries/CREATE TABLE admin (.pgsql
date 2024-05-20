CREATE TABLE admin (
    fname varchar (255) NOT NULL,
    lname varchar(255) DEFAULT NULL,
    email varchar(255) NOT NULL,
    gender varchar(255) NOT NULL,
    Password varchar(255),
    PRIMARY KEY(email)
);
