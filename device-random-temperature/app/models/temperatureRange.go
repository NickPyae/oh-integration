// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package models

type TemperatureRange struct {
	MinTemperature int64 `json:"minTemperature" validate:"required,max=10000,min=-459"`
	MaxTemperature int64 `json:"maxTemperature" validate:"required,max=10000,min=-459,gtefield=MinTemperature"`
}
