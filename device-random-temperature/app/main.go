// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
	"os"

	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-random-temperature/api"
	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-random-temperature/helpers"
	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-random-temperature/scripts"
)

func main() {

	if os.Getenv("CORE_SVCS_IP") == "" || os.Getenv("CORE_DATA_PORT") == "" || os.Getenv("CORE_METADATA_PORT") == "" || os.Getenv("ADDRESSABLE_PORT") == "" {
		log.Println("Please ensure env variables 'CORE_SVCS_IP', 'CORE_DATA_PORT', 'CORE_METADATA_PORT', 'ADDRESSABLE_PORT' is present.")
		os.Exit(1)
	}

	helpers.CoreServicesIP = os.Getenv("CORE_SVCS_IP")
	helpers.CoreServicesBaseURL = "https://" + os.Getenv("CORE_SVCS_IP")
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
