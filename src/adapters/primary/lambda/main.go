package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	databaseadapterdynamodb "github.com/lucasrosa/alerts/src/adapters/secondary/database"
	"github.com/lucasrosa/alerts/src/corelogic/alerts"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	alertsRepo := databaseadapterdynamodb.NewDynamoAlertsRepository()
	alertsService := alerts.NewAlertsService(alertsRepo)
	alertsAdapter := NewAlertsAdapter(alertsService)

	return alertsAdapter.GetActive(request)
}

func main() {
	lambda.Start(Handler)
}
