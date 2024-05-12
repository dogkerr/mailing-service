package structs

type Data struct {
	*VerificationData
	*BillingNoticeData
}

type Message struct {
	TemplateType string   `json:"templateType"`
	Data         Data     `json:"data"`
	Subject      string   `json:"subject"`
	To           []string `json:"to"`
}
