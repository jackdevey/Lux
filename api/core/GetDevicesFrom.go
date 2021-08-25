package core

import (
	"github.com/bandev/lux/api/general"
	"github.com/bandev/lux/commands/devices"
	"github.com/fatih/color"
	"os"
	"strconv"
)

func GetDevicesFrom(arg string, c general.Connection) []devices.Device {
	var ds devices.Devices
	ds.Get(c)
	if arg == "@a" {
		return ds.Data.Devices
	} else {
		var dID, _ = strconv.Atoi(arg)
		if len(ds.Data.Devices) <= dID || dID < 0 {
			general.PrintHeading("Device id provided is invalid", color.FgRed)
			os.Exit(1)
		}
		return []devices.Device{ds.Data.Devices[dID]}
	}
}