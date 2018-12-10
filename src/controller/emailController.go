package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	dto "../dto"
	es "../service"
	u "../utils"
)

var emailService = es.EmailService{}

// SendEmail ...
var SendEmail = func(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body)
	var emailDto dto.EmailDTO
	json.Unmarshal(b, &emailDto)

	resp := emailService.SendEmail(emailDto)
	u.Respond(w, resp)

}

//GetEmails method
func GetEmails(w http.ResponseWriter, r *http.Request) {
	emails := emailService.GetEmails()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emails)
}
