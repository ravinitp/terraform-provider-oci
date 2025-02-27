// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// DiffChunk Details about a group of changes.
type DiffChunk struct {

	// Line number in base version where changes begin.
	BaseLine *int `mandatory:"false" json:"baseLine"`

	// Number of lines chunk spans in base version.
	BaseSpan *int `mandatory:"false" json:"baseSpan"`

	// Line number in target version where changes begin.
	TargetLine *int `mandatory:"false" json:"targetLine"`

	// Number of lines chunk spans in target version.
	TargetSpan *int `mandatory:"false" json:"targetSpan"`

	// List of difference section.
	DiffSections []DiffSection `mandatory:"false" json:"diffSections"`
}

func (m DiffChunk) String() string {
	return common.PointerString(m)
}
