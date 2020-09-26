package handlers

import (
	middleware "github.com/go-openapi/runtime/middleware"

	"strings"

	"github.com/phoax/godemo/config"
	"github.com/phoax/godemo/models"
	op "github.com/phoax/godemo/restapi/operations/homepage"
)

var apiconfig = config.Config()

// Homepage handler
func HomepageHandler() middleware.Responder {
	serviceName := apiconfig.GetString("SERVICE_NAME")

	// Response
	payload := models.Ack{Status: "success", Message: strings.Join([]string{"welcome to ", serviceName, "api"}, " ")}
	return op.NewHomepageOK().WithPayload(&payload)
}
