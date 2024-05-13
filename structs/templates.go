package structs

type TemplateType string

const (
	Verification  TemplateType = "verification"
	BillingNotice TemplateType = "billing-notice"
)

type BillingNoticeData struct {
	Username string  `json:"username" validate:"required"`
	Balance  float32 `json:"balance" validate:"gt=0"`
}

type VerificationData struct {
	Name             string `json:"name" validate:"required"`
	VerificationLink string `json:"verificationLink" validate:"required,url"`
}

func (d *BillingNoticeData) Validate() error {
	return Validator.Struct(d)
}

func (d *VerificationData) Validate() error {
	return Validator.Struct(d)
}
