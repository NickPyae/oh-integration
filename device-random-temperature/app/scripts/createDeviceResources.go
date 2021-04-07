// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package scripts

import (
	"bytes"
	"net/http"
)

// todo: refactor when resolve issue with importing package 'helpers'
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

func CreateAddressables() {
	path := "/api/v1/addressable"

	var jsonStr = []byte(`{
			"name": "` + deviceServiceName + `",
			"protocol": "HTTP",
			"address": "` + coreServicesBaseURL + `",
			"port": ` + devicePort + `,
			"path": "/api/v1/device/register"
		}`)
	req, err := http.NewRequest("POST", coreMetadataURL+path, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}

func CreateValueDescriptors() {
	path := "/api/v1/valuedescriptor"

	var jsonStr = []byte(`{
			"name": "` + resourceName + `",
			"description": "Random temperature readings in Fahrenheit",
			"type": "Int32",
			"defaultValue": "` + defaultMinTemperature + `",
			"labels":["` + deviceServiceName + `"]
		}`)
	req, err := http.NewRequest("POST", coreDataURL+path, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}

func CreateDeviceService() {
	path := "/api/v1/deviceservice"

	var jsonStr = []byte(`{
			"name": "` + deviceServiceName + `",
			"description": "Generates random temperature readings in Fahrenheit",
			"labels":["` + deviceServiceName + `"],
			"adminState":"unlocked",
			"operatingState":"enabled",
			"addressable": {
				"name":"` + deviceServiceName + `"
			}
		}`)
	req, err := http.NewRequest("POST", coreMetadataURL+path, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}

func CreateDevice() {
	path := "/api/v1/device"

	var jsonStr = []byte(`{
			"name": "` + deviceName + `",
			"description": "Generates random temperature readings in Fahrenheit",
			"adminState": "unlocked",
			"operatingState": "enabled",
			"protocols": {
				"example": {
				"host": "localhost",
				"port": "0",
				"unitID": "1"
				}
			},
			"addressable": {
				"name": "` + deviceServiceName + `"
			},
			"labels": [
				"` + deviceServiceName + `"
			],
			"service": {
				"name": "` + deviceServiceName + `" 
			},
			"profile": {
				"name": "` + deviceProfileName + `"
			}
		}`)
	req, err := http.NewRequest("POST", coreMetadataURL+path, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}
