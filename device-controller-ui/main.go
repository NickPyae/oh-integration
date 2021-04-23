// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
	"os"

	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-controller-ui/api"
	"eos2git.cec.lab.emc.com/ISG-Edge/HelloSally/device-controller-ui/helpers"
)

func main() {
	if os.Getenv("CORE_SVCS_IP") == "" {
		log.Println("Please ensure env variables 'CORE_SVCS_IP' is present.")
		os.Exit(1)
	}

	helpers.CoreCommandPort = "48082"

	// override default values if env var is set
	if os.Getenv("CORE_COMMAND_PORT") != "" {
		helpers.CoreCommandPort = os.Getenv("CORE_COMMAND_PORT")
	}

	helpers.CoreServicesIP = os.Getenv("CORE_SVCS_IP")
	helpers.CoreServicesBaseURL = "http://" + os.Getenv("CORE_SVCS_IP")

	helpers.CoreCommandURL = helpers.CoreServicesBaseURL + ":" + helpers.CoreCommandPort

	api.SetRoutes()
}
