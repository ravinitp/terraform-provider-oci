// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// StorageWorkRequestCollection This is a list of work request summaries.
type StorageWorkRequestCollection struct {

	// This is an array of work request summaries.
	Items []StorageWorkRequestSummary `mandatory:"true" json:"items"`
}

func (m StorageWorkRequestCollection) String() string {
	return common.PointerString(m)
}
