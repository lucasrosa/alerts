package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lucasrosa/alerts/src/corelogic/alerts"
	//databaseadapterdynamodb
)

// AlertsAdapter is the interface that defines the entrypoints to this adapter
type AlertsAdapter interface {
	GetActive(request events.APIGatewayProxyRequest) (Response, error)
}

type alertsAdapter struct {
	alertsService alerts.AlertsPrimaryPort
}

func NewAlertsAdapter(alertsService alerts.AlertsPrimaryPort) AlertsAdapter {
	return &alertsAdapter{
		alertsService,
	}
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// GetActive receives the request, processes it and returns a Response or an error
func (a *alertsAdapter) GetActive(request events.APIGatewayProxyRequest) (Response, error) {

	// Verifying the body of the request
	// alert := alerts.Alert{}
	// err := json.Unmarshal([]byte(request.Body), &alert)
	// if err != nil {
	// 	return Response{StatusCode: 400}, nil
	// }

	alerts, err := a.alertsService.GetActive()
	fmt.Println(alerts)
	if err != nil {
		return Response{StatusCode: 502}, err
	}

	return successfulResponse(), nil
}

func successfulResponse() Response {
	return Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Methods":     "GET",
			"Access-Control-Allow-Headers":     "application/json",
		},
	}
}
