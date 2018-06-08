package main

import (
	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/handlers"
)

func main() {
	wf := workflow.NewAPIGWProxyWorkflowBuilder().
		AddGetHandler("/device", handlers.GetAll).
		AddGetHandler("/device/{deviceId}", handlers.GetSingle).
		AddPutHandler("/device/{deviceId}", handlers.UserAction).
		Build()
	lambda.Start(wf.GetLambdaHandler())
}
