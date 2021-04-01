// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package driver

import (
        "fmt"
        "math/rand"
        "time"

        "github.com/edgexfoundry/device-sdk-go/pkg/models"
)

const (
        defMinTemperature, defMaxTemperature = 50, 200
)

type randomDevice struct {
        minTemperature int64
        maxTemperature int64
}

func (d *randomDevice) value(valueType models.ValueType) (int64, error) {
        // This code block checks the max and min value integrity every time because device-random allows users to modify
        // the max and min values at runtime by Put commands
        switch valueType {
        case models.Int32:
                if d.maxTemperature <= d.minTemperature {
                        return 0, fmt.Errorf("randomDevice.value: maximum: %d <= minimum : %d", d.maxTemperature, d.minTemperature)
                } else {
                        return random(d.minTemperature, d.maxTemperature), nil
                }
        default:
                return 0, fmt.Errorf("randomDevice.value: wrong value type: %v", valueType)
        }
}

func newRandomDevice() *randomDevice {
        return &randomDevice{
                minTemperature: defMinTemperature,
                maxTemperature: defMaxTemperature,
        }
}

func random(min int64, max int64) int64 {
        rand.Seed(time.Now().UnixNano())
        return rand.Int63n(max-min) + min
}
