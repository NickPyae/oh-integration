// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-random-temperature/scripts"

	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-random-temperature/api"
)

func main() {
	scripts.CreateAddressables()
	scripts.CreateValueDescriptors()
	scripts.UploadDeviceProfile()
	scripts.CreateDeviceService()
	scripts.CreateDevice()

	api.SetRoutes()
}
