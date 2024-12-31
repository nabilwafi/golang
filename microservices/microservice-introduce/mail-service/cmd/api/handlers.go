package main

import (
	"log"
	"net/http"
)

func (c *Config) sendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage

	err := c.ReadJSON(w, r, &requestPayload)
	if err != nil {
		log.Println(err)
		c.ErrorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	err = c.Mailer.SendSMTPMessage(msg)
	if err != nil {
		log.Println(err)
		c.ErrorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "sent to " + requestPayload.To,
	}

	c.WriteJSON(w, http.StatusAccepted, payload)
}
