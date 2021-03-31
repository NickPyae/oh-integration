// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a simple example of a device service.
package main

import (
        "device-random-temperature"
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