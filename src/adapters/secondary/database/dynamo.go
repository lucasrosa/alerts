package databaseadapterdynamodb

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/lucasrosa/alerts/src/corelogic/alerts"
)

type alertsRepository struct{}

// NewDynamoAlertsRepository instantiates the repository for this adapter
func NewDynamoAlertsRepository() alerts.AlertsSecondaryPort {
	return &alertsRepository{}
}

func (r *alertsRepository) GetActive() ([]alerts.Alert, error) {
	fmt.Println("retrieving")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		return []alerts.Alert{alerts.Alert{}}, err
	}

	//alerts := []alerts.Alert{alerts.Alert{ID: "1", Start: 10, End: 20}}
	svc := dynamodb.New(sess)

	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		IndexName: aws.String("start"),
		KeyConditions: map[string]*dynamodb.Condition{
			"end": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("50"),
					},
				},
			},
		},
	}

	var resp1, err1 = svc.Query(queryInput)
	if err1 != nil {
		fmt.Println(err1)
	}

	alerts := []alerts.Alert{}
	err = dynamodbattribute.UnmarshalListOfMaps(resp1.Items, &alerts)
	fmt.Println(alerts)

	return alerts, nil
}

// func (r *checkoutRepository) Save(order *cart.Order) error {
// 	fmt.Println("saving order", order)

// 	sess, err := session.NewSession(&aws.Config{
// 		Region: aws.String("us-east-1")},
// 	)

// 	svc := dynamodb.New(sess)

// 	persistedOrder := PersistedOrder{
// 		ID:        order.ID,
// 		Email:     order.Email,
// 		Amount:    order.Amount,
// 		Currency:  order.Currency,
// 		ProductID: order.ProductID,
// 	}
// 	fmt.Println("Persisting order:", persistedOrder)

// 	// Marshall the Item into a Map DynamoDB can deal with
// 	av, err := dynamodbattribute.MarshalMap(persistedOrder)
// 	if err != nil {
// 		fmt.Println("Got error marshalling map:")
// 		fmt.Println(err.Error())
// 		return err
// 	}

// 	// Create Item in table and return
// 	input := &dynamodb.PutItemInput{
// 		Item:      av,
// 		TableName: aws.String(os.Getenv("TABLE_NAME")),
// 	}

// 	_, err = svc.PutItem(input)
// 	if err != nil {
// 		fmt.Println("Error while sending message to sqs", err)
// 	} else {
// 		fmt.Println("Success while sending message to sqs")
// 	}

// 	return err
// }
