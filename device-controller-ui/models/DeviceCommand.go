// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package models

type DeviceCommand struct {
	Name string     `json:"name" validate:"required"`
	Put  PutCommand `json:"put" validate:"required"`
}
