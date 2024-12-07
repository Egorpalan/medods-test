package utils

import (
	"log"
)

func SendEmailWarning(email, message string) {
	log.Printf("Sending email warning to %s: %s", email, message)
}
