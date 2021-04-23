// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package models

type CommandInfo struct {
	Name     string          `json:"name" validate:"required"`
	Commands []DeviceCommand `json:"commands" validate:"required"`
}
