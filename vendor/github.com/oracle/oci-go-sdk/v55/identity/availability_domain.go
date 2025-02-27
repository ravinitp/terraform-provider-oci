// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// AvailabilityDomain One or more isolated, fault-tolerant Oracle data centers that host cloud resources such as instances, volumes,
// and subnets. A region contains several Availability Domains. For more information, see
// Regions and Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm).
type AvailabilityDomain struct {

	// The name of the Availability Domain.
	Name *string `mandatory:"false" json:"name"`

	// The OCID of the Availability Domain.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the tenancy.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m AvailabilityDomain) String() string {
	return common.PointerString(m)
}
