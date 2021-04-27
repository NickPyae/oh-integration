// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"

	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-controller-ui/models"

	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-controller-ui/helpers"
)

func SetRoutes() {
	fileServer := http.FileServer(http.Dir("./static/"))

	r := mux.NewRouter()

	// POST
	r.HandleFunc("/api/v1/submitRequest", SubmitRequestHandler)

	r.PathPrefix("/").Handler(fileServer)

	log.Println("Listening on :49990")
	log.Fatal(http.ListenAndServe(":49990", r))
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func getPutCommandUrl() (url string) {
	// get device commands from core command
	path := "/api/v1/device"

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", helpers.CoreCommandURL+path, nil)

	if err != nil {
		return url
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return url
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return url
	}

	// get PUT url from response body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return url
	}

	var commandInfo []models.CommandInfo
	unmarshallErr := json.Unmarshal(body, &commandInfo)

	if unmarshallErr != nil {
		return url
	}

	for i := range commandInfo {
		info := commandInfo[i]
		if info.Name == helpers.DeviceName {
			for j := range info.Commands {
				command := info.Commands[j]
				url = command.Put.Url
				break
			}
		}
	}

	return url
}

func executePutCommand(url string, form url.Values) (err error) {
	idx := strings.Index(url, "/api/v1/device")
	path := url[idx:]

	duration := form["duration"][0]
	minTemperature := form["minTemperature"][0]
	maxTemperature := form["maxTemperature"][0]

	var jsonStr = []byte(`{
		"duration": ` + duration + `,
		"range": {"MinTemperature": ` + minTemperature + `, "MaxTemperature": ` + maxTemperature + `}
	   }`)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "PUT", helpers.CoreCommandURL+path, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("")
	}

	_, err = ioutil.ReadAll(resp.Body)

	return err
}

func SubmitRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		url := getPutCommandUrl()
		if url == "" {
			fmt.Fprintf(w, "Unable to submit request")
			return
		}

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Unable to submit request")
			return
		}

		err := executePutCommand(url, r.PostForm)
		if err != nil {
			fmt.Fprintf(w, "Unable to submit request")
			return
		}

		fmt.Fprint(w, "Request submitted")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
