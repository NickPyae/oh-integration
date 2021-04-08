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

	if os.Getenv("CORE_SVCS_IP") == "" || os.Getenv("DEVICE_SVC_IP") == "" {
		log.Println("Please ensure env variables 'CORE_SVCS_IP' and 'DEVICE_SVC_IP' is present.")
		os.Exit(1)
	}

	// override default values if env var is set
	if os.Getenv("CORE_DATA_PORT") != "" {
		helpers.CoreDataPort = os.Getenv("CORE_DATA_PORT")
	}

	if os.Getenv("CORE_METADATA_PORT") != "" {
		helpers.CoreMetadataPort = os.Getenv("CORE_METADATA_PORT")
	}

	if os.Getenv("ADDRESSABLE_PORT") != "" {
		helpers.AddressablePort = os.Getenv("ADDRESSABLE_PORT")
	}

	helpers.DeviceServiceIP = os.Getenv("DEVICE_SVC_IP")
	helpers.CoreServicesIP = os.Getenv("CORE_SVCS_IP")
	helpers.CoreServicesBaseURL = "http://" + os.Getenv("CORE_SVCS_IP")
	helpers.CoreDataPort = "48080"
	helpers.CoreMetadataPort = "48081"
	helpers.AddressablePort = "49989"
	helpers.CoreDataURL = helpers.CoreServicesBaseURL + ":" + helpers.CoreDataPort
	helpers.CoreMetadataURL = helpers.CoreServicesBaseURL + ":" + helpers.CoreMetadataPort

	scripts.CreateAddressables()
	scripts.CreateValueDescriptors()
	scripts.UploadDeviceProfile()
	scripts.CreateDeviceService()
	scripts.CreateDevice()

	api.SetRoutes()
}
