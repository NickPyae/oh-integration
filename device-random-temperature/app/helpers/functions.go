// Copyright 2021 Dell Inc, or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func RandomIntStr(min int64, max int64) string {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return strconv.FormatInt(n+min, 10)
}
