package services

import (
	"fmt"

	"github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/types"
)

const (
	rackUserID = "rack"
)

// DeviceService ...
type DeviceService struct {
}

// UpdateDevice ...
func (service *DeviceService) UpdateDevice(deviceID, userID string) (*types.DeviceState, error) {
	dbService := new(DBService)
	device, err := dbService.GetDeviceByID(deviceID)
	if err != nil {
		return nil, err
	}

	var dbUserID string
	var deviceState = new(types.DeviceState)
	if userID == device.UserID {
		// Return device.
		dbUserID = rackUserID
		deviceState.DeviceState = "Returned"
	} else {
		// Get device.
		dbUserID = userID
		deviceState.DeviceState = "Taken"
	}

	err = dbService.UpdateDevice(deviceID, dbUserID)
	if err != nil {
		return nil, err
	}

	deviceHistoryService := new(DeviceHistoryService)
	err = deviceHistoryService.UpdateHistory(userID, dbUserID, deviceID)
	if err != nil {
		return nil, err
	}
	return deviceState, nil
}

// GetFullDevice ...
func (service *DeviceService) GetFullDevice(deviceID string) (*types.DeviceExt, error) {
	dbService := new(DBService)
	device, err := dbService.GetDeviceByID(deviceID)
	if err != nil {
		return nil, err
	}
	deviceExt := types.DeviceExt{
		DeviceID: device.DeviceID,
		UserID:   device.UserID,
		Date:     device.Date,
		ImageURL: fmt.Sprintf("https://s3-eu-west-1.amazonaws.com/ns-devices-rack-organizer/%s/image.jpg", device.DeviceID),
	}
	return &deviceExt, nil
}

// GetFullDevices ...
func (service *DeviceService) GetFullDevices() (*[]types.DeviceExt, error) {
	dbService := new(DBService)
	devices, err := dbService.GetAllDevices()
	if err != nil {
		return nil, err
	}
	devicesExt := []types.DeviceExt{}
	for _, device := range devices {
		deviceExt := types.DeviceExt{
			DeviceID: device.DeviceID,
			UserID:   device.UserID,
			Date:     device.Date,
			ImageURL: fmt.Sprintf("https://s3-eu-west-1.amazonaws.com/ns-devices-rack-organizer/%s/image.jpg", device.DeviceID),
		}
		devicesExt = append(devicesExt, deviceExt)
	}
	return &devicesExt, nil
}
