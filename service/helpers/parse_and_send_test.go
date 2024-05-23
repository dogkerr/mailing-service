package helpers

import (
	"testing"

	"github.com/dogkerr/mailing-service/m/v2/domain"
)

func TestParseAndSend(t *testing.T) {
	mockData := domain.VerificationData{
		Name:             "John Doe",
		Email:            "davidlou0810@gmail.com",
		VerificationLink: "https://example.com/verify",
	}
	ParseAndSend("templates/verification.html", mockData, mockData.Email, "Account Verification Email")
}
