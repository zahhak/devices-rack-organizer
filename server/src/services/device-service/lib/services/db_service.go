package services

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/types/db"
)

// DBService ...
type DBService struct {
}

// GetAllDevices returns all devices stored in the database.
func (service *DBService) GetAllDevices() ([]db.Device, error) {
	sess := session.New()
	client := dynamodb.New(sess)
	dbInput := &dynamodb.ScanInput{
		TableName: aws.String("RackOrganizer"),
	}
	allItems, err := client.Scan(dbInput)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	devices := []db.Device{}
	device := new(db.Device)
	for _, Item := range allItems.Items {
		err = dynamodbattribute.UnmarshalMap(Item, device)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		devices = append(devices, *device)
		fmt.Println(devices)
	}
	return devices, nil
}
