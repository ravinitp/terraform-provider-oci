// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetric, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For information about monitoring, see Monitoring Overview (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm).
//

package monitoring

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// Datapoint Metric value for a specific timestamp.
type Datapoint struct {

	// Timestamp for this metric value. Format defined by RFC3339.
	// Example: `2019-02-01T01:02:29.600Z`
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// Numeric value of the metric.
	// Example: `10.23`
	Value *float64 `mandatory:"true" json:"value"`

	// The number of occurrences of the associated value in the set of data.
	// Default is 1. Value must be greater than zero.
	Count *int `mandatory:"false" json:"count"`
}

func (m Datapoint) String() string {
	return common.PointerString(m)
}
