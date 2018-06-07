package handlers

import (
	"fmt"
	"net/http"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/requests"
)

// UserAction is the handler for Get/Return device.
func UserAction(ctx workflow.Context, req requests.UserAction) error {
	fmt.Println(req)
	ctx.SetResponseStatusCode(http.StatusNoContent)
	return nil
}
