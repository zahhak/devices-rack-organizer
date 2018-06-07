package services

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/types/db"
)

// DBService ...
type DBService struct {
}

// GetAllDevices returns all devices stored in the database.
func (service *DBService) GetAllDevices() ([]db.Device, error) {
	return []db.Device{}, nil
}

// UpdateDevice ...
func (service *DBService) UpdateDevice(deviceID, userID string) error {
	client := dynamodb.New(session.New())
	item, err := dynamodbattribute.MarshalMap(db.Device{
		DeviceID: deviceID,
		UserID:   userID,
		Date:     time.Now().UTC(),
	})
	if err != nil {
		return err
	}
	req := &dynamodb.PutItemInput{
		TableName: aws.String("RackOrganizer"),
		Item:      item,
	}
	_, err = client.PutItem(req)
	if err != nil {
		return err
	}
	return nil
}

// GetDeviceByID ...
func (service *DBService) GetDeviceByID(id string) (*db.Device, error) {
	deviceKey, err := dynamodbattribute.MarshalMap(struct{ DeviceID string }{DeviceID: id})
	if err != nil {
		return nil, err
	}

	client := dynamodb.New(session.New())
	req := &dynamodb.GetItemInput{
		TableName: aws.String("RackOrganizer"),
		Key:       deviceKey,
	}
	res, err := client.GetItem(req)
	if err != nil {
		return nil, err
	}

	device := new(db.Device)
	err = dynamodbattribute.UnmarshalMap(res.Item, device)
	if err != nil {
		return nil, err
	}
	return device, nil
}
