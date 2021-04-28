// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package models

type PutCommand struct {
	Url string `json:"url" validate:"required"`
}
