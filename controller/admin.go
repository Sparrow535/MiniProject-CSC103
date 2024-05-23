package controller

import (
	"encoding/json"
	"myapp/model"
	httpresp "myapp/utils/httpResp"
	"net/http"
	"time"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&admin); err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	saveErr := admin.CreateAdmin()
	if saveErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}
	httpresp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "admin added"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	getErr := admin.GetAdmin()
	if getErr != nil {
		httpresp.RespondWithError(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	// Creating an instance of http.Cookie type
	cookie := http.Cookie{
		Name:    "my-cookie",
		Value:   "my-value",
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
	}

	// Setting cookie
	http.SetCookie(w, &cookie)

	// Respond with JSON containing the gender
	response := map[string]string{
		"message": "success",
		"gender":  admin.Gender,
	}
	httpresp.RespondWithJSON(w, http.StatusOK, response)
}

func GetDetails(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	getErr := admin.GetAdmin()
	if getErr != nil {
		httpresp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpresp.RespondWithJSON(w, http.StatusOK, admin)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "my-cookie",
		Expires: time.Now(),
	})

	httpresp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "cookie deleted"})
}

func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			httpresp.RespondWithError(w, http.StatusSeeOther, "cookie not found")
			return false
		}
		httpresp.RespondWithError(w, http.StatusInternalServerError, "internal server error")
		return false
	}
	if cookie.Value != "my-value" {
		httpresp.RespondWithError(w, http.StatusUnauthorized, "cookie does not match")
		return false
	}
	return true
}
