package structs

type Data struct {
	*VerificationData
	*BillingNoticeData
}

type Message struct {
	TemplateType string   `json:"templateType" validate:"required,oneof=verification billing-notice"`
	Data         Data     `json:"data" validate:"required"`
	Subject      string   `json:"subject" validate:"required"`
	To           []string `json:"to" validate:"gt=0,dive,email"`
}

func (m *Message) GetTemplateData() interface{} {
	switch m.TemplateType {
	case "verification":
		return m.Data.VerificationData
	case "billing-notice":
		return m.Data.BillingNoticeData
	}
	return nil
}

func (m *Message) Validate() error {
	return Validator.Struct(m)
}
