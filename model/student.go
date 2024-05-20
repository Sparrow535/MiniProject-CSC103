package model

import "myapp/dataBase/postgres"

type Student struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EnrollmentNo int    `json:"enrollment_no"`
	ContactNo    int    `json:"contact_no"`
	Room         string `json:"room"`
}

// Boys' queries and methods
const (
	queryInsertStudent = "INSERT INTO student(fname, lname, enrollment_no, contact_no, room) VALUES ($1, $2, $3, $4, $5) RETURNING enrollment_no;"
	queryGetStudent    = "SELECT fname, lname, enrollment_no, contact_no, room FROM student WHERE enrollment_no = $1;"
	queryUpdateStudent = "UPDATE student SET enrollment_no=$1, fname = $2, lname = $3, contact_no = $4, room = $5 WHERE enrollment_no = $6 RETURNING enrollment_no;"
	queryDeleteStudent = "DELETE FROM student WHERE enrollment_no = $1 RETURNING enrollment_no;"
)

func (student *Student) AddStudent() error {
	_, err := postgres.Db.Exec(queryInsertStudent, student.FirstName, student.LastName, student.EnrollmentNo, student.ContactNo, student.Room)
	return err
}

func (student *Student) GetStudent() error {
	return postgres.Db.QueryRow(queryGetStudent, student.EnrollmentNo).Scan(&student.FirstName, &student.LastName, &student.EnrollmentNo, &student.ContactNo, &student.Room)
}

func (student *Student) UpdateStudent(oldEnrollmentNo int64) error {
	return postgres.Db.QueryRow(queryUpdateStudent, student.EnrollmentNo, student.FirstName, student.LastName, student.ContactNo, student.Room, oldEnrollmentNo).Scan(&student.EnrollmentNo)
}

func (student *Student) DeleteStudent() error {
	return postgres.Db.QueryRow(queryDeleteStudent, student.EnrollmentNo).Scan(&student.EnrollmentNo)
}

func GetAllStudents() ([]Student, error) {
	rows, err := postgres.Db.Query("SELECT fname, lname, enrollment_no, contact_no, room FROM student;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.FirstName, &s.LastName, &s.EnrollmentNo, &s.ContactNo, &s.Room); err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}

// Girls' queries and methods
const (
	queryInsertStudentG = "INSERT INTO studentg(fname, lname, enrollment_no, contact_no, room) VALUES ($1, $2, $3, $4, $5) RETURNING enrollment_no;"
	queryGetStudentG    = "SELECT fname, lname, enrollment_no, contact_no, room FROM studentg WHERE enrollment_no = $1;"
	queryUpdateStudentG = "UPDATE studentg SET enrollment_no=$1, fname = $2, lname = $3, contact_no = $4, room = $5 WHERE enrollment_no = $6 RETURNING enrollment_no;"
	queryDeleteStudentG = "DELETE FROM studentg WHERE enrollment_no = $1 RETURNING enrollment_no;"
)

type StudentG struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EnrollmentNo int    `json:"enrollment_no"`
	ContactNo    int    `json:"contact_no"`
	Room         string `json:"room"`
}

func (student *StudentG) AddStudentG() error {
	_, err := postgres.Db.Exec(queryInsertStudentG, student.FirstName, student.LastName, student.EnrollmentNo, student.ContactNo, student.Room)
	return err
}

func (student *StudentG) GetStudentG() error {
	return postgres.Db.QueryRow(queryGetStudentG, student.EnrollmentNo).Scan(&student.FirstName, &student.LastName, &student.EnrollmentNo, &student.ContactNo, &student.Room)
}

func (student *StudentG) UpdateStudentG(oldEnrollmentNo int64) error {
	return postgres.Db.QueryRow(queryUpdateStudentG, student.EnrollmentNo, student.FirstName, student.LastName, student.ContactNo, student.Room, oldEnrollmentNo).Scan(&student.EnrollmentNo)
}

func (student *StudentG) DeleteStudentG() error {
	return postgres.Db.QueryRow(queryDeleteStudentG, student.EnrollmentNo).Scan(&student.EnrollmentNo)
}

func GetAllStudentsG() ([]StudentG, error) {
	rows, err := postgres.Db.Query("SELECT fname, lname, enrollment_no, contact_no, room FROM studentg;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []StudentG
	for rows.Next() {
		var s StudentG
		if err := rows.Scan(&s.FirstName, &s.LastName, &s.EnrollmentNo, &s.ContactNo, &s.Room); err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}
