package controller

import (
	"net/http"

	"github.com/draganaF/gled.io/EmailService/dto"
	"github.com/draganaF/gled.io/EmailService/helpers"
	"github.com/draganaF/gled.io/EmailService/services"
)

var SendEmail = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var email *dto.Email
	err := helpers.ReadJSONBody(r, &email)

	if err != nil {
		helpers.JSONResponse(w, 400, "Failed to deserialize request body.")
	}

	emailSender := services.NewEmailSender()
	hasSent := emailSender.Send(email)
	if !hasSent {
		helpers.JSONResponse(w, 400, "Failed to send email.")
		return
	}

	helpers.JSONResponse(w, 200, email)
})
