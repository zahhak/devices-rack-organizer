package services

import "github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/types"

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
