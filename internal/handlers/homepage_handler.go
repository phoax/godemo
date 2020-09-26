package handlers

import (
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/phoax/godemo/models"
	op "github.com/phoax/godemo/restapi/operations/homepage"
)

// Homepage handler
func HomepageHandler() middleware.Responder {

	// Response
	payload := models.Ack{Status: "success", Message: "welcome to the api"}
	return op.NewHomepageOK().WithPayload(&payload)
}
