// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package functionaltests includes tests that make http requests to the handlers using net/http/test
package functionaltests

import (
	"math/rand"

	"github.com/docker/docker/api/types"
)

func getMockStats() *types.Stats {
	return &types.Stats{
		CPUStats: types.CPUStats{
			SystemUsage: uint64(rand.Intn(10000)),
		},
	}
}
