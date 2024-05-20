package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	httpresp "myapp/utils/httpResp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Boys' handlers
func AddStudent(w http.ResponseWriter, r *http.Request) {
	var stud model.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	if err := stud.AddStudent(); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpresp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "student added"})
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	sid := mux.Vars(r)["enrollment_no"]
	enrollmentNo, err := strconv.Atoi(sid)
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	s := model.Student{EnrollmentNo: enrollmentNo}
	if err := s.GetStudent(); err != nil {
		if err == sql.ErrNoRows {
			httpresp.RespondWithError(w, http.StatusNotFound, "Student not found")
		} else {
			httpresp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	httpresp.RespondWithJSON(w, http.StatusOK, s)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	oldSid := mux.Vars(r)["enrollment_no"]
	oldEnrollmentNo, err := strconv.Atoi(oldSid)
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var stud model.Student
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	if err := stud.UpdateStudent(int64(oldEnrollmentNo)); err != nil {
		if err == sql.ErrNoRows {
			httpresp.RespondWithError(w, http.StatusNotFound, "Student not found")
		} else {
			httpresp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	httpresp.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"status": "student updated", "updated": stud})
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["enrollment_no"]
	enrollmentNo, err := strconv.Atoi(sid)
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	s := model.Student{EnrollmentNo: enrollmentNo}
	if err := s.DeleteStudent(); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpresp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}

	students, err := model.GetAllStudents()
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpresp.RespondWithJSON(w, http.StatusOK, students)
}

// Girls' handlers
func AddStudentG(w http.ResponseWriter, r *http.Request) {
	var stud model.StudentG
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	if err := stud.AddStudentG(); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpresp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "student added"})
}

func GetStudentG(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	sid := mux.Vars(r)["enrollment_no"]
	enrollmentNo, err := strconv.Atoi(sid)
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	s := model.StudentG{EnrollmentNo: enrollmentNo}
	if err := s.GetStudentG(); err != nil {
		if err == sql.ErrNoRows {
			httpresp.RespondWithError(w, http.StatusNotFound, "Student not found")
		} else {
			httpresp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	httpresp.RespondWithJSON(w, http.StatusOK, s)
}

func UpdateStudentG(w http.ResponseWriter, r *http.Request) {
	oldSid := mux.Vars(r)["enrollment_no"]
	oldEnrollmentNo, err := strconv.Atoi(oldSid)
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var stud model.StudentG
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stud); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	if err := stud.UpdateStudentG(int64(oldEnrollmentNo)); err != nil {
		if err == sql.ErrNoRows {
			httpresp.RespondWithError(w, http.StatusNotFound, "Student not found")
		} else {
			httpresp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	httpresp.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"status": "student updated", "updated": stud})
}

func DeleteStudentG(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["enrollment_no"]
	enrollmentNo, err := strconv.Atoi(sid)
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	s := model.StudentG{EnrollmentNo: enrollmentNo}
	if err := s.DeleteStudentG(); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpresp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func GetAllStudentsG(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}

	students, err := model.GetAllStudentsG()
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpresp.RespondWithJSON(w, http.StatusOK, students)
}
