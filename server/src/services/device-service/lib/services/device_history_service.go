package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/types"
)

// DeviceHistoryService ...
type DeviceHistoryService struct {
}

// UpdateHistory ...
func (service *DeviceHistoryService) UpdateHistory(oldOwner, newOwner, deviceID string) error {
	client := s3.New(session.New())
	historyKey := service.getHistoryKey(deviceID)
	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String("ns-devices-rack-organizer"),
		Key:    historyKey,
	}
	res, err := client.GetObject(getObjectInput)
	if err != nil {
		if e, ok := err.(awserr.Error); !ok {
			return err
		} else if e.Code() != s3.ErrCodeNoSuchKey {
			return err
		}
	}

	parsedHistory, err := service.getHistory(res)
	if err != nil {
		return err
	}
	record := types.DeviceHistoryRecord{
		OldOwner: oldOwner,
		NewOwner: newOwner,
		Date:     time.Now().UTC(),
	}
	newHistory := append(*parsedHistory, record)
	newContent, err := json.Marshal(newHistory)
	if err != nil {
		return err
	}

	putObjectInput := &s3.PutObjectInput{
		Bucket: aws.String("ns-devices-rack-organizer"),
		Body:   bytes.NewReader(newContent),
		Key:    historyKey,
		ACL:    aws.String("public-read"),
	}
	_, err = client.PutObject(putObjectInput)
	if err != nil {
		return err
	}

	return nil
}

func (service *DeviceHistoryService) getHistoryKey(deviceID string) *string {
	return aws.String(fmt.Sprintf("%s/history.json", deviceID))
}

func (service *DeviceHistoryService) getHistory(s3Res *s3.GetObjectOutput) (*[]types.DeviceHistoryRecord, error) {
	hasItem := s3Res != nil && s3Res.Body != nil
	defer func() {
		if hasItem {
			s3Res.Body.Close()
		}
	}()

	res := new([]types.DeviceHistoryRecord)
	if hasItem {
		historyContent, err := ioutil.ReadAll(s3Res.Body)
		if err != nil {
			return nil, nil
		}

		err = json.Unmarshal(historyContent, res)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
