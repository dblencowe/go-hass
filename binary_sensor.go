package hass

import (
	"fmt"
	"time"
)

type BinarySensorAttributes struct {
	PresenceSensors []string      `json:"presence_sensors"`
	Features        []string      `json:"features"`
	ActiveSensors   []interface{} `json:"active_sensors"`
	Lights          []string      `json:"lights"`
	ClearTimeout    int           `json:"clear_timeout"`
	UpdateInterval  int           `json:"update_interval"`
	Type            string        `json:"type"`
	Climate         []interface{} `json:"climate"`
	OnStates        []string      `json:"on_states"`
	AutomaticLights string        `json:"automatic_lights"`
	Night           bool          `json:"night"`
	Sleep           bool          `json:"sleep"`
	FriendlyName    string        `json:"friendly_name"`
	Icon            string        `json:"icon"`
	DeviceClass     string        `json:"device_class"`
}

type BinarySensor struct {
	access      *Access
	id          string                 `json:"entity_id"`
	State       string                 `json:"state"`
	Attributes  BinarySensorAttributes `json:"attributes"`
	LastChanged time.Time              `json:"last_changed"`
	LastUpdated time.Time              `json:"last_updated"`
	Context     struct {
		ID       string      `json:"id"`
		ParentID interface{} `json:"parent_id"`
		UserID   interface{} `json:"user_id"`
	} `json:"context"`
}

func (a *Access) NewBinarySensor(id string) (bs *BinarySensor, err error) {
	var state BinarySensor
	state.access = a
	err = a.httpGet("/api/states/"+id, &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

// On turns on a binary sensor (Currently Unsupported)
func (bs *BinarySensor) On() (err error) {
	return fmt.Errorf("Unsupported Action: On %+v\n", bs)
}

// Off turns off a binary sensor (Currently Unsupported)
func (bs *BinarySensor) Off() (err error) {
	return fmt.Errorf("Unsupported Action: Off %+v\n", bs)
}

// Toggle toggles a switch
func (bs *BinarySensor) Toggle() (err error) {
	return bs.access.CallService("switch", "toggle", bs.id)
}

// EntityID returns the id of the device object
func (bs *BinarySensor) EntityID() string {
	return bs.id
}

// Domain returns the Home Assistant domain for the device
func (bs *BinarySensor) Domain() string {
	return "binary_sensor"
}
