// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package models

type TemperatureRequest struct {
	TemperatureRange TemperatureRange `json:"range" validate:"required"`
	Duration         int64            `json:"duration" validate:"required,min=2"`
}
