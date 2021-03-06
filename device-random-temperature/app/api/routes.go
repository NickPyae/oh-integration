// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"

	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-random-temperature/models"

	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-random-temperature/helpers"
)

var (
	minTemperature, maxTemperature, duration int64 = 50, 200, 10
	timer                                    *time.Ticker
	validate                                 *validator.Validate
)

func SetRoutes() {
	fileServer := http.FileServer(http.Dir("./static/"))

	r := mux.NewRouter()

	// GET
	r.HandleFunc("/getTemperatureRange", GetTemperatureRangeHandler)
	r.HandleFunc("/getDeviceReadings", GetDeviceReadingsHandler)

	// POST
	r.HandleFunc("/api/v1/device/register", RegisterDeviceHandler)

	// PUT
	r.HandleFunc("/api/v1/device/{deviceId}/GenerateRandomTemperature", ChangeTemperatureRangeHandler)

	r.PathPrefix("/").Handler(fileServer)

	log.Println("Listening on :49989")
	log.Fatal(http.ListenAndServe(":49989", r))
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func GetTemperatureRangeHandler(w http.ResponseWriter, r *http.Request) {
	t := models.TemperatureRange{
		MinTemperature: minTemperature,
		MaxTemperature: maxTemperature,
	}

	if err := json.NewEncoder(w).Encode(t); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
	}
}

func RegisterDeviceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "Device registered")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func ChangeTemperatureRangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		decoder := json.NewDecoder(r.Body)
		var t models.TemperatureRequest
		decodeErr := decoder.Decode(&t)

		if decodeErr != nil {
			// checks for field type
			http.Error(w, "invalid parameters", http.StatusInternalServerError)
			return
		}

		validate = validator.New()
		validateErr := validate.Struct(t)
		if validateErr != nil {
			// checks for validators
			http.Error(w, "invalid parameters", http.StatusInternalServerError)
			return
		}

		minTemperature = t.TemperatureRange.MinTemperature
		maxTemperature = t.TemperatureRange.MaxTemperature
		duration = t.Duration // in seconds

		d := time.Duration(duration)
		if timer != nil {
			timer.Stop()
		}

		timer := time.NewTimer(d * time.Second)
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-timer.C:
					minTemperature, _ = strconv.ParseInt(helpers.DefaultMinTemperature, 10, 64)
					maxTemperature, _ = strconv.ParseInt(helpers.DefaultMaxTemperature, 10, 64)

					timer.Stop()
				case <-done:
					return
				}
			}
		}()

		fmt.Fprint(w, "Command accepted")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func GetDeviceReadingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		path := "/api/v1/reading/device/"
		limit := "20"

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, "GET", helpers.CoreDataURL+path+helpers.DeviceName+"/"+limit, nil)

		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}

		var readings []models.TemperatureReading
		unmarshallErr := json.Unmarshal(body, &readings)

		if unmarshallErr != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}

		for i := range readings {
			dt := time.Unix(0, readings[i].Created*int64(time.Millisecond))
			readings[i].CreatedDateTime = dt.String()
		}

		if err := json.NewEncoder(w).Encode(readings); err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
		}
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func AddDeviceReading() bool {
	success := false

	path := "/api/v1/event"

	var jsonStr = []byte(`{
                        "device": "` + helpers.DeviceName + `",
                        "readings": [{"name": "` + helpers.ResourceName + `", "value":"` + helpers.RandomIntStr(minTemperature, maxTemperature) + `"}]
       				}`)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", helpers.CoreDataURL+path, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return success
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return success
	}
	defer resp.Body.Close()

	success = true
	return success
}
