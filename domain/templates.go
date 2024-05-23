package domain

type TemplateType string

const (
	Verification  TemplateType = "verification"
	BillingNotice TemplateType = "billing-notice"
)

type BillingNoticeData struct {
	Name                 string  `json:"name" validate:"required"`
	Email                string  `json:"email" validate:"required,email"`
	ContainerID          string  `json:"container_id"`
	TotalCPUUsage        float32 `json:"total_cpu_usage"`
	TotalMemoryUsage     float32 `json:"total_memory_usage"`
	TotalNetIngressUsage float32 `json:"total_net_ingress_usage"`
	TotalNetEgressUsage  float32 `json:"total_net_egress_usage"`
	Timestamp            string  `json:"timestamp"`
	TotalCost            float32 `json:"total_cost"`
}

type VerificationData struct {
	Name             string `json:"name" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	VerificationLink string `json:"verificationLink" validate:"required,url"`
}

type ContainerDownNoticeData struct {
	Name             string `json:"name" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	SwarmServiceName string `json:"swarm_service_name"`
	SwarmServiceID   string `json:"swarm_service_id"`
}

func (d *BillingNoticeData) Validate() error {
	return Validator.Struct(d)
}

func (d *VerificationData) Validate() error {
	return Validator.Struct(d)
}
