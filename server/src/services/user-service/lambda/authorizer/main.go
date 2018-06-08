package main

import (
	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/zahhak/devices-rack-organizer/server/src/services/user-service/lib/handlers"
)

func main() {
	wf := workflow.NewAPIGWAuthorizerWorkflowBuilder().
		SetHandler(handlers.Authorizer).
		Build()
	lambda.Start(wf.GetLambdaHandler())
}
