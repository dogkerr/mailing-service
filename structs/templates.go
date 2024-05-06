package structs

type TemplateType string

const (
	Verification  TemplateType = "verification"
	BillingNotice TemplateType = "billing-notice"
)

type BillingNoticeData struct {
	Name       string
	ExpiryDate string
}

type VerificationData struct {
	Name             string
	VerificationLink string
}
