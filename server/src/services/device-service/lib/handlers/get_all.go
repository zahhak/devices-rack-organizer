package handlers

import (
	"net/http"

	"github.com/TsvetanMilanov/go-lambda-workflow/workflow"
)

// GetAll ...
func GetAll(ctx workflow.Context) error {
	res := []struct{}{}
	ctx.SetResponse(res).SetResponseStatusCode(http.StatusOK)
	return nil
}
