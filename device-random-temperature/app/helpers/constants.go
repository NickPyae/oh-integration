// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package helpers

const (
	CoreServicesBaseURL = "https://10.244.14.32"
	CoreDataPort        = "3080"
	CoreMetadataPort    = "3081"

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
