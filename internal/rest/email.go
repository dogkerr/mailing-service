package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/dogkerr/mailing-service/m/v2/domain"
	"github.com/dogkerr/mailing-service/m/v2/internal/repository"
	"github.com/dogkerr/mailing-service/m/v2/service/helpers"
)

type EmailService interface {
	SendEmailContainerDown(ctx context.Context, d domain.Message) error
}

type EmailHandler struct {
}

func EmailRouter(r *server.Hertz) {
	handler := &EmailHandler{}

	root := r.Group("/api/v1")
	{
		eH := root.Group("/email")
		{
			eH.POST("/down", handler.SendDownEmail)
		}
	}
}

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

type commonLabels struct {
	Alertname                       string `json:"alertname"`
	ContainerSwarmServiceID         string `json:"container_label_com_docker_swarm_service_id"`
	ContainerDockerSwarmServiceName string `json:"container_label_com_docker_swarm_service_name"`
	ContainerLabelUserID            string `json:"container_label_user_id"`
}

type promeWebhookReq struct {
	Receiver     string       `json:"receiver"`
	CommonLabels commonLabels `json:"commonLabels"`
}

type promeWebhookRes struct {
	Message string `json:"message"`
}

func (h *EmailHandler) SendDownEmail(ctx context.Context, c *app.RequestContext) {
	var req promeWebhookReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	fmt.Println(req)
	fmt.Println(req.CommonLabels.ContainerLabelUserID)
	fmt.Println(req.CommonLabels.ContainerDockerSwarmServiceName)
	fmt.Println(req.CommonLabels.ContainerSwarmServiceID)

	// Handle send email
	user, err := repository.NewService().GetUserById(req.CommonLabels.ContainerLabelUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	templateData := domain.ContainerDownNoticeData{
		Name:             user.GetFullname(),
		Email:            user.GetEmail(),
		SwarmServiceName: req.CommonLabels.ContainerDockerSwarmServiceName,
		SwarmServiceID:   req.CommonLabels.ContainerSwarmServiceID,
	}

	err = helpers.ParseAndSend("templates/service-down-notice.html", templateData, user.GetEmail(), "Swarm Service Down Notice")
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, promeWebhookRes{Message: "ok"})
}
