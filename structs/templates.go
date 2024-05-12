package structs

type TemplateType string

const (
	Verification  TemplateType = "verification"
	BillingNotice TemplateType = "billing-notice"
)

type BillingNoticeData struct {
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

type VerificationData struct {
	Name             string `json:"name"`
	VerificationLink string `json:"verificationLink"`
}
