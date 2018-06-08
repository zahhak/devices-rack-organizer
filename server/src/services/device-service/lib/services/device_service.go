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
func (service *DeviceService) UpdateDevice(deviceID, userID string) error {
	dbService := new(DBService)
	device, err := dbService.GetDeviceByID(deviceID)
	if err != nil {
		return err
	}

	var dbUserID string
	if userID == device.UserID {
		// Return device.
		dbUserID = rackUserID
	} else {
		// Get device.
		dbUserID = userID
	}

	err = dbService.UpdateDevice(deviceID, dbUserID)
	if err != nil {
		return err
	}
	return nil
}

// GetFullDevice ...
func (service *DeviceService) GetFullDevice(deviceID string) (*types.DeviceExt, error) {
	dbService := new(DBService)
	device, err := dbService.GetDeviceByID(deviceID)
	if err != nil {
		return nil, err
	}
	deviceExt := types.DeviceExt{
		Device: *device,
		Image:  fmt.Sprintf("https://s3-eu-west-1.amazonaws.com/ns-devices-rack-organizer/%s/image.jpg", device.DeviceID),
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
			Device: device,
			Image:  fmt.Sprintf("https://s3-eu-west-1.amazonaws.com/ns-devices-rack-organizer/%s/image.jpg", device.DeviceID),
		}
		devicesExt = append(devicesExt, deviceExt)
	}
	return &devicesExt, nil
}
