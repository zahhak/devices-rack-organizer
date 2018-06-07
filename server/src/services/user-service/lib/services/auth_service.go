package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/zahhak/devices-rack-organizer/server/src/services/user-service/types"
)

// AuthService ...
type AuthService struct {
}

// GetAuthorizerResponse sets the current user in the lambda context and returns the
// custom authorizer response.
func (auth *AuthService) GetAuthorizerResponse(evt events.APIGatewayCustomAuthorizerRequest) (res *events.APIGatewayCustomAuthorizerResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = getUnauthorizedError(errors.New("panic"))
		}
	}()

	userID := parseAuthHeader(evt.AuthorizationToken)
	user, err := json.Marshal(types.UserInfo{UserID: userID})
	if err != nil {
		return nil, getUnauthorizedError(err)
	}

	return &events.APIGatewayCustomAuthorizerResponse{
		PrincipalID:    userID,
		PolicyDocument: getDummyPolicyDocument(evt),
		Context: map[string]interface{}{
			"User": string(user),
		},
	}, nil
}

func getDummyPolicyDocument(evt events.APIGatewayCustomAuthorizerRequest) events.APIGatewayCustomAuthorizerPolicy {
	statement := events.IAMPolicyStatement{
		Action:   []string{"execute-api:Invoke"},
		Effect:   "Allow",
		Resource: []string{evt.MethodArn},
	}
	return events.APIGatewayCustomAuthorizerPolicy{
		Version: "2012-10-17",
		Statement: []events.IAMPolicyStatement{
			statement,
		},
	}
}

func parseAuthHeader(token string) string {
	re := regexp.MustCompile("Bearer|Basic (.*)")
	return re.FindStringSubmatch(token)[1]
}

func getUnauthorizedError(err error) error {
	fmt.Println("Authorizer error:", err)
	return errors.New("Unauthorized")
}
