package handlers

import (
	"net/http"

	"github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/services"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
)

// GetAll ...
func GetAll(ctx workflow.Context) error {
	dbService := services.DBService{}
	res, err := dbService.GetAllDevices()
	if err != nil {
		return err
	}
	// res := []struct{}{}
	ctx.SetResponse(res).SetResponseStatusCode(http.StatusOK)
	return nil
}
