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
	"time"

	"github.com/gorilla/mux"

	"device-random-temperature/helpers"
	"device-random-temperature/models"
)

var (
	minTemperature, maxTemperature int64 = 50, 200
)

func SetRoutes() {
	fileServer := http.FileServer(http.Dir("./static/"))

	r := mux.NewRouter()

	// GET
	r.HandleFunc("/getTemperatureRange", GetTemperatureRangeHandler)
	r.HandleFunc("/getDeviceReadings", GetDeviceReadingsHandler)
	r.HandleFunc("/addDeviceReading", AddDeviceReadingHandler)

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
		var t models.TemperatureRange
		err := decoder.Decode(&t)

		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
		} else {
			minTemperature = t.MinTemperature
			maxTemperature = t.MaxTemperature

			fmt.Fprint(w, "Command accepted")
		}
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
			dt := time.Unix(readings[i].Created, 0)
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

func AddDeviceReadingHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
	defer resp.Body.Close()
}
