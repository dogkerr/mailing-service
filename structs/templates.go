package structs

type TemplateType string

const (
	Verification  TemplateType = "verification"
	BillingNotice TemplateType = "billing-notice"
)

type BillingNoticeData struct {
	Username string `json:"username"`
	Balance  string `json:"balance"`
}

type VerificationData struct {
	Username         string `json:"username"`
	VerificationLink string `json:"verificationLink"`
}
