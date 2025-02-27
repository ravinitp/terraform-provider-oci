// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// UpdateInstancePoolPlacementConfigurationDetails The location for where an instance pool will place instances.
type UpdateInstancePoolPlacementConfigurationDetails struct {

	// The availability domain to place instances.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the primary subnet to place instances.
	PrimarySubnetId *string `mandatory:"true" json:"primarySubnetId"`

	// The fault domains to place instances.
	// If you don't provide any values, the system makes a best effort to distribute
	// instances across all fault domains based on capacity.
	// To distribute the instances evenly across selected fault domains, provide a
	// set of fault domains. For example, you might want instances to be evenly
	// distributed if your applications require high availability.
	// To get a list of fault domains, use the
	// ListFaultDomains operation
	// in the Identity and Access Management Service API.
	// Example: `[FAULT-DOMAIN-1, FAULT-DOMAIN-2, FAULT-DOMAIN-3]`
	FaultDomains []string `mandatory:"false" json:"faultDomains"`

	// The set of subnet OCIDs for secondary VNICs for instances in the pool.
	SecondaryVnicSubnets []InstancePoolPlacementSecondaryVnicSubnet `mandatory:"false" json:"secondaryVnicSubnets"`
}

func (m UpdateInstancePoolPlacementConfigurationDetails) String() string {
	return common.PointerString(m)
}
