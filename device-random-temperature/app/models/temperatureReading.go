// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package models

type TemperatureReading struct {
	Name            string `json:"name"`
	Value           string `json:"value"`
	ValueType       string `json:"valueType"`
	Created         int64  `json:"created"`
	CreatedDateTime string `json:"createdDateTime"`
}
