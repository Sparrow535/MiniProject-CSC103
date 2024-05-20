CREATE TABLE student (
    Fname varchar(255) NOT NULL,
    Lname varchar(255) DEFAULT NULL,
    Enrollment_no varchar(255) NOT NULL,
    Contact_no int NOT NULL,
    Room varchar(255),
    PRIMARY KEY(Enrollment_no),
    FOREIGN KEY (Room) REFERENCES room(Room_no)
)