// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// AssociatedDatabaseSummary The summary of a database currently using a Database Management private endpoint.
type AssociatedDatabaseSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database.
	Id *string `mandatory:"true" json:"id"`

	// The name of the database.
	Name *string `mandatory:"true" json:"name"`

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time when Database Management was enabled for the database.
	TimeRegistered *common.SDKTime `mandatory:"true" json:"timeRegistered"`
}

func (m AssociatedDatabaseSummary) String() string {
	return common.PointerString(m)
}
