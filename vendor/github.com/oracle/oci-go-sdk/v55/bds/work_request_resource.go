// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// WorkRequestResource A resource created or operated on by a work request.
type WorkRequestResource struct {

	// The resource type the work request affects.
	EntityType *string `mandatory:"true" json:"entityType"`

	// The way in which this resource is affected by the work tracked in the work request. A resource being created, updated, or deleted remains in the IN_PROGRESS state until work is complete for that resource, at which point it will transition to CREATED, UPDATED, or DELETED, respectively.
	ActionType ActionTypesEnum `mandatory:"true" json:"actionType"`

	// The OCID of the resource the work request affects.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The URI path that the user can do a GET on to access the resource metadata.
	EntityUri *string `mandatory:"false" json:"entityUri"`
}

func (m WorkRequestResource) String() string {
	return common.PointerString(m)
}
