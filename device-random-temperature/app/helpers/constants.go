// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package helpers

const (
	coreServicesBaseURL = "http://192.168.56.144"
	coreDataPort        = "30800"
	coreMetadataPort    = "30801"

	coreDataURL     = coreServicesBaseURL + ":" + coreDataPort
	coreMetadataURL = coreServicesBaseURL + ":" + coreMetadataPort

	deviceName        = "Random-Temperature-Generator01"
	devicePort        = "30080"
	resourceName      = "RandomTemperature"
	deviceServiceName = "device-random-temperature"
	deviceProfileName = "Random-Temperature-Generator"

	defaultMinTemperature = "50"
	defaultMaxTemperature = "200"
)
