package main

import (
	"bytes"

	"gobot.io/x/gobot"
	ble "gobot.io/x/gobot/platforms/ble"
)

// from https://github.com/kvantetore/igrill/blob/713f4cb058cfff0807c7507541468b7c463dd3ca/igrill.py#L15
var (
	probe1_temp_uuid      = "06ef0002-2e06-4b79-9e33-fce2c42805ec"
	probe1_threshold_uuid = "06ef0003-2e06-4b79-9e33-fce2c42805ec"
	probe2_temp_uuid      = "06ef0004-2e06-4b79-9e33-fce2c42805ec"
	probe2_threshold_uuid = "06ef0005-2e06-4b79-9e33-fce2c42805ec"
	probe3_temp_uuid      = "06ef0006-2e06-4b79-9e33-fce2c42805ec"
	probe3_threshold_uuid = "06ef0007-2e06-4b79-9e33-fce2c42805ec"
	probe4_temp_uuid      = "06ef0008-2e06-4b79-9e33-fce2c42805ec"
	probe4_threshold_uuid = "06ef0009-2e06-4b79-9e33-fce2c42805ec"
)

// IGrillDriver represents the Battery Service for a BLE Peripheral
type IGrillDriver struct {
	name       string
	connection gobot.Connection
	gobot.Eventer
}

// NewIGrillDriver creates a BatteryDriver
func NewIGrillDriver(a ble.BLEConnector) *IGrillDriver {
	n := &IGrillDriver{
		name:       gobot.DefaultName("Battery"),
		connection: a,
		Eventer:    gobot.NewEventer(),
	}

	return n
}

// Connection returns the Driver's Connection to the associated Adaptor
func (b *IGrillDriver) Connection() gobot.Connection { return b.connection }

// Name returns the Driver name
func (b *IGrillDriver) Name() string { return b.name }

// SetName sets the Driver name
func (b *IGrillDriver) SetName(n string) { b.name = n }

// adaptor returns BLE adaptor
func (b *IGrillDriver) adaptor() ble.BLEConnector {
	return b.Connection().(ble.BLEConnector)
}

// Start tells driver to get ready to do work
func (b *IGrillDriver) Start() (err error) {
	return
}

// Halt stops battery driver (void)
func (b *IGrillDriver) Halt() (err error) { return }

func (b *IGrillDriver) getValue(uuid string) uint8 {
	var l uint8
	c, _ := b.adaptor().ReadCharacteristic(uuid)
	buf := bytes.NewBuffer(c)
	val, _ := buf.ReadByte()
	l = uint8(val)
	return l
}

// GetProbe1Temp reads and returns the tempature level for probe 1
func (b *IGrillDriver) GetProbe1Temp() uint8 {
	return b.getValue(probe1_temp_uuid)
}

// GetProbe2Temp reads and returns the tempature level for probe 2
func (b *IGrillDriver) GetProbe2Temp() uint8 {
	return b.getValue(probe2_temp_uuid)
}

// GetProbe3Temp reads and returns the tempature level for probe 3
func (b *IGrillDriver) GetProbe3Temp() uint8 {
	return b.getValue(probe3_temp_uuid)
}

// GetProbe4Temp reads and returns the tempature level for probe 4
func (b *IGrillDriver) GetProbe4Temp() uint8 {
	return b.getValue(probe4_temp_uuid)
}

// GetThreshold1 read the current threshold value for probe 1
func (b IGrillDriver) GetThreshold1() uint8 {
	return b.getValue(probe1_threshold_uuid)
}

// GetThreshold2 read the current threshold value for probe 2
func (b IGrillDriver) GetThreshold2() uint8 {
	return b.getValue(probe2_threshold_uuid)
}

// GetThreshold3 read the current threshold value for probe 3
func (b IGrillDriver) GetThreshold3() uint8 {
	return b.getValue(probe3_threshold_uuid)
}

// GetThreshold4 read the current threshold value for probe 4
func (b IGrillDriver) GetThreshold4() uint8 {
	return b.getValue(probe4_threshold_uuid)
}
