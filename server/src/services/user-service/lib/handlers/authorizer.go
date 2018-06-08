package handlers

import (
	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/aws/aws-lambda-go/events"
	"github.com/zahhak/devices-rack-organizer/server/src/services/user-service/lib/services"
)

// Authorizer is the handler for the API Gateway Custom Authorizer.
func Authorizer(ctx workflow.Context, evt events.APIGatewayCustomAuthorizerRequest) (err error) {
	authService := new(services.AuthService)
	res, err := authService.GetAuthorizerResponse(evt)
	if err != nil {
		return err
	}

	ctx.SetRawResponse(res)
	return nil
}
