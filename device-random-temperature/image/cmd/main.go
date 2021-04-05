// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	device_random_temperature "device-random-temperature"
	"device-random-temperature/driver"

	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
)

const (
	serviceName string = "device-random-temperature"
)

func main() {
	d := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, device_random_temperature.Version, d)
}
