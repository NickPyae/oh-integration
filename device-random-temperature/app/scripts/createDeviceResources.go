// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package scripts

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-random-temperature/helpers"
)

func postJSONCoreSvc(baseurl string, path string, jsonArr []byte) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", baseurl+path, bytes.NewBuffer(jsonArr))
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

func CreateAddressables() {
	path := "/api/v1/addressable"

	var jsonArr = []byte(`{
			"name": "` + helpers.DeviceServiceName + `",
			"protocol": "HTTP",
			"address": "` + helpers.CoreServicesBaseURL + `",
			"port": ` + helpers.AddressablePort + `,
			"path": "/api/v1/device/register"
		}`)

	postJSONCoreSvc(helpers.CoreMetadataURL, path, jsonArr)
}

func CreateValueDescriptors() {
	path := "/api/v1/valuedescriptor"

	var jsonArr = []byte(`{
			"name": "` + helpers.ResourceName + `",
			"description": "Random temperature readings in Fahrenheit",
			"type": "Int32",
			"defaultValue": "` + helpers.DefaultMinTemperature + `",
			"labels":["` + helpers.DeviceServiceName + `"]
		}`)

	postJSONCoreSvc(helpers.CoreDataURL, path, jsonArr)
}

func CreateDeviceService() {
	path := "/api/v1/deviceservice"

	var jsonArr = []byte(`{
			"name": "` + helpers.DeviceServiceName + `",
			"description": "Generates random temperature readings in Fahrenheit",
			"labels":["` + helpers.DeviceServiceName + `"],
			"adminState":"unlocked",
			"operatingState":"enabled",
			"addressable": {
				"name":"` + helpers.DeviceServiceName + `"
			}
		}`)

	postJSONCoreSvc(helpers.CoreMetadataURL, path, jsonArr)
}

func CreateDevice() {
	path := "/api/v1/device"

	var jsonArr = []byte(`{
			"name": "` + helpers.DeviceName + `",
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
				"name": "` + helpers.DeviceServiceName + `"
			},
			"labels": [
				"` + helpers.DeviceServiceName + `"
			],
			"service": {
				"name": "` + helpers.DeviceServiceName + `" 
			},
			"profile": {
				"name": "` + helpers.DeviceProfileName + `"
			}
		}`)

	postJSONCoreSvc(helpers.CoreMetadataURL, path, jsonArr)
}

func UploadDeviceProfile() {
	absPath, _ := filepath.Abs("./scripts/device-profile.yaml")

	// prepare the reader instances to encode
	values := map[string]io.Reader{
		"file": mustOpen(absPath),
	}
	client := &http.Client{}
	path := "/api/v1/deviceprofile/uploadfile"
	url := helpers.CoreMetadataURL + path
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
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", url, &b)
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
	return err
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}
