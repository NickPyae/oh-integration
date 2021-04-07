// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package models

type TemperatureRange struct {
	MinTemperature int64 `json:"minTemperature"`
	MaxTemperature int64 `json:"maxTemperature"`
}
