// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-random-temperature/scripts"

	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-random-temperature/api"
	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-random-temperature/helpers"
)

func main() {
	os.Setenv("CORE_SVCS_IP", "https://10.244.14.32")
	os.Setenv("CORE_DATA_PORT", "3080")
	os.Setenv("CORE_METADATA_PORT", "3081")
	os.Setenv("ADDRESSABLE_PORT", "49989")

	helpers.CoreServicesBaseURL = os.Getenv("CORE_SVCS_IP")
	helpers.CoreDataPort = os.Getenv("CORE_DATA_PORT")
	helpers.CoreMetadataPort = os.Getenv("CORE_METADATA_PORT")
	helpers.AddressablePort = os.Getenv("ADDRESSABLE_PORT")
	helpers.CoreDataURL = helpers.CoreServicesBaseURL + ":" + helpers.CoreDataPort
	helpers.CoreMetadataURL = helpers.CoreServicesBaseURL + ":" + helpers.CoreMetadataPort

	scripts.CreateAddressables()
	scripts.CreateValueDescriptors()
	scripts.UploadDeviceProfile()
	scripts.CreateDeviceService()
	scripts.CreateDevice()

	api.SetRoutes()
}
