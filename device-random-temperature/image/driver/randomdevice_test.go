// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package driver

import (
        "testing"

        "github.com/edgexfoundry/device-sdk-go/pkg/models"
)

var device = newRandomDevice()


func TestValue_Int32(t *testing.T) {
        valueType := models.Int32

        val, err := device.value(valueType)

        if err != nil {
                t.Fatalf("Failed to generate random %v value", valueType)
        }
        if val <= defMinTemperature || val >= defMaxTemperature {                            
                t.Fatalf("Unexpected test result. %v is not in %v value range", val, valueType)
        }
}
