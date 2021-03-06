// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
	"os"
	"time"

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
	helpers.CoreCommandPort = "48082"
	helpers.AddressablePort = "49989"
	helpers.CoreDataURL = helpers.CoreServicesBaseURL + ":" + helpers.CoreDataPort
	helpers.CoreMetadataURL = helpers.CoreServicesBaseURL + ":" + helpers.CoreMetadataPort
	helpers.CoreCommandURL = helpers.CoreServicesBaseURL + ":" + helpers.CoreCommandPort

	initDevice()

	addDeviceReadings()
	api.SetRoutes()
}

func addDeviceReadings() {
	i := 1000
	d := time.Duration(i)
	ticker := time.NewTicker(d * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				success := api.AddDeviceReading()
				if !success {
					initDevice()
				}
			}
		}
	}()
}

func initDevice() {
	// (blocking) ensure connection to core services
	scripts.CheckConnection()

	scripts.CreateAddressables()
	scripts.CreateValueDescriptors()
	scripts.UploadDeviceProfile()
	scripts.CreateDeviceService()
	scripts.CreateDevice()
}
