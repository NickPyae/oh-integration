// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"device-random-temperature/api"
	"device-random-temperature/scripts"
)

func main() {
	scripts.CreateAddressables()
	scripts.CreateValueDescriptors()
	scripts.UploadDeviceProfile()
	scripts.CreateDeviceService()
	scripts.CreateDevice()

	api.SetRoutes()
}
