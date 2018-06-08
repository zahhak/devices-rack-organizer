package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/aws/aws-lambda-go/events"
	"github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/requests"
	"github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/services"
	"github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/types"
)

// UserAction is the handler for Get/Return device.
func UserAction(ctx workflow.Context, req requests.UserAction) error {
	evt := new(events.APIGatewayProxyRequest)
	err := ctx.GetLambdaEvent(evt)
	if err != nil {
		return err
	}
	userInfoValue, ok := evt.RequestContext.Authorizer["User"]
	if !ok {
		return errors.New("unable to find user in the authorizer context")
	}

	user := new(types.User)
	unmarshalErr := json.Unmarshal([]byte(userInfoValue.(string)), user)
	if unmarshalErr != nil {
		return unmarshalErr
	}

	deviceService := new(services.DeviceService)
	deviceState, updateErr := deviceService.UpdateDevice(req.DeviceID, user.UserID)
	if updateErr != nil {
		return updateErr
	}
	ctx.SetResponse(deviceState).SetResponseStatusCode(http.StatusOK)
	return nil
}
