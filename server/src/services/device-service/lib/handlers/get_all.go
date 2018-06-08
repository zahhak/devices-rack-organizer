package handlers

import (
	"net/http"

	"github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/requests"
	"github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/services"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
)

// GetAll ...
func GetAll(ctx workflow.Context) error {
	deviceService := services.DeviceService{}
	res, err := deviceService.GetFullDevices()
	if err != nil {
		return err
	}
	// res := []struct{}{}
	ctx.SetResponse(res).SetResponseStatusCode(http.StatusOK)
	return nil
}

// GetSingle ...
func GetSingle(ctx workflow.Context, req requests.UserAction) error {
	deviceService := services.DeviceService{}
	res, err := deviceService.GetFullDevice(req.DeviceID)
	if err != nil {
		return err
	}
	ctx.SetResponse(res).SetResponseStatusCode(http.StatusOK)
	return nil
}
