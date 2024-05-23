package domain

import "testing"

func TestMessageValidation(t *testing.T) {
	m := Message{
		TemplateType: "verification",
		Data: Data{
			VerificationData: &VerificationData{
				Name:             "test",
				VerificationLink: "http://example.com",
			},
		},
		Subject: "Test",
		To:      []string{"davidlou0810@gmail.com"}}

	err := m.Validate()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestBillingNoticeValidation(t *testing.T) {
	m := Message{
		TemplateType: "billing-notice",
		Data: Data{
			BillingNoticeData: &BillingNoticeData{
				Username: "test",
				Balance:  100,
			},
		},
		Subject: "Test",
		To:      []string{"davidlou0810@gmail.com"}}

	err := m.Validate()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestVerificationDataValidation(t *testing.T) {
	d := VerificationData{
		Name:             "test",
		VerificationLink: "http://example.com",
	}

	err := d.Validate()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
