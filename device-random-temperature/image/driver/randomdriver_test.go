// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package driver

import (
        "testing"
        "time"

        dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
        "github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
        "github.com/edgexfoundry/go-mod-core-contracts/models"
)

var d *RandomDriver

func init() {
        d = new(RandomDriver)
        d.lc = logger.NewClient("devicerandom", false, "", "DEBUG")
}

func TestHandleReadCommands(t *testing.T) {
        deviceName := "testDevice"
        protocols := map[string]models.ProtocolProperties{
                "other": {
                        "Address": "simple01",
                        "Port":    "300",
                },
        }

        requests := []dsModels.CommandRequest{
                {
                        DeviceResourceName: "RandomTemperature",
                        Type:               dsModels.Int32,
                },
        }

        res, err := d.HandleReadCommands(deviceName, protocols, requests)

        if err != nil {
                t.Fatalf("Failed to read command, %v", err)
        }
        if len(res) != len(requests) {
                t.Fatalf("Number of results fetched '%v' should match '%v'", len(res), len(requests))
        }
        if res[0].DeviceResourceName != "RandomTemperature" {
                t.Fatalf("Unexpected test result. Wrong resource object.")
        }
        if res[0].Type != dsModels.Int32 {
                t.Fatalf("Unexpected test result. Wrong value type.")
        }
}

func TestHandleWriteCommands(t *testing.T) {
        deviceName := "testDevice"
        protocols := map[string]models.ProtocolProperties{
                "other": {
                        "Address": "simple01",
                        "Port":    "300",
                },
        }
        var requests []dsModels.CommandRequest

        now := time.Now().UnixNano()
        cv, err := dsModels.NewInt32Value("Max_Temperature", now, int32(127))
        if err != nil {
                t.Fatalf("Failed to create command value, %v", err)
        }
        params := []*dsModels.CommandValue{cv}

        err = d.HandleWriteCommands(deviceName, protocols, requests, params)

        if err != nil {
                t.Fatalf("Failed to write command, %v", err)
        }
}
