// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package scripts

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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

func UploadDeviceProfile() {
	absPath, _ := filepath.Abs("./scripts/device-profile.yaml")

	//prepare the reader instances to encode
	values := map[string]io.Reader{
		"file": mustOpen(absPath),
	}
	client := &http.Client{}
	path := "/api/v1/deviceprofile/uploadfile"
	url := coreMetadataURL + path
	err := Upload(client, url, values)
	if err != nil && err.Error() == "bad status: 409 Conflict" {
		log.Println("Device profile cannot be created. Pls check if it already exist.")
	} else if err != nil {
		panic(err)
	}
}

func Upload(client *http.Client, url string, values map[string]io.Reader) (err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	return
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}
