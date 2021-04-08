// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package helpers

const (
	CoreServicesBaseURL = "http://192.168.56.144"
	CoreDataPort        = "30800"
	CoreMetadataPort    = "30801"

	CoreDataURL     = CoreServicesBaseURL + ":" + CoreDataPort
	CoreMetadataURL = CoreServicesBaseURL + ":" + CoreMetadataPort

	DeviceName        = "Random-Temperature-Generator01"
	AddressablePort   = "49989"
	ResourceName      = "RandomTemperature"
	DeviceServiceName = "device-random-temperature"
	DeviceProfileName = "Random-Temperature-Generator"

	DefaultMinTemperature = "50"
	DefaultMaxTemperature = "200"
)
