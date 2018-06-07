package services

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
