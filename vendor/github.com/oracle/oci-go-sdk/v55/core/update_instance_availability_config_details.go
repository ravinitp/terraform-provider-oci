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

// UpdateInstanceAvailabilityConfigDetails Options for defining the availability of a VM instance after a maintenance event that impacts the underlying
// hardware, including whether to live migrate supported VM instances when possible without sending a prior customer notification.
type UpdateInstanceAvailabilityConfigDetails struct {

	// Whether to live migrate supported VM instances to a healthy physical VM host without
	// disrupting running instances during infrastructure maintenance events. If null, Oracle
	// chooses the best option for migrating the VM during infrastructure maintenance events.
	IsLiveMigrationPreferred *bool `mandatory:"false" json:"isLiveMigrationPreferred"`

	// The lifecycle state for an instance when it is recovered after infrastructure maintenance.
	// * `RESTORE_INSTANCE` - The instance is restored to the lifecycle state it was in before the maintenance event.
	// If the instance was running, it is automatically rebooted. This is the default action when a value is not set.
	// * `STOP_INSTANCE` - The instance is recovered in the stopped state.
	RecoveryAction UpdateInstanceAvailabilityConfigDetailsRecoveryActionEnum `mandatory:"false" json:"recoveryAction,omitempty"`
}

func (m UpdateInstanceAvailabilityConfigDetails) String() string {
	return common.PointerString(m)
}

// UpdateInstanceAvailabilityConfigDetailsRecoveryActionEnum Enum with underlying type: string
type UpdateInstanceAvailabilityConfigDetailsRecoveryActionEnum string

// Set of constants representing the allowable values for UpdateInstanceAvailabilityConfigDetailsRecoveryActionEnum
const (
	UpdateInstanceAvailabilityConfigDetailsRecoveryActionRestoreInstance UpdateInstanceAvailabilityConfigDetailsRecoveryActionEnum = "RESTORE_INSTANCE"
	UpdateInstanceAvailabilityConfigDetailsRecoveryActionStopInstance    UpdateInstanceAvailabilityConfigDetailsRecoveryActionEnum = "STOP_INSTANCE"
)

var mappingUpdateInstanceAvailabilityConfigDetailsRecoveryAction = map[string]UpdateInstanceAvailabilityConfigDetailsRecoveryActionEnum{
	"RESTORE_INSTANCE": UpdateInstanceAvailabilityConfigDetailsRecoveryActionRestoreInstance,
	"STOP_INSTANCE":    UpdateInstanceAvailabilityConfigDetailsRecoveryActionStopInstance,
}

// GetUpdateInstanceAvailabilityConfigDetailsRecoveryActionEnumValues Enumerates the set of values for UpdateInstanceAvailabilityConfigDetailsRecoveryActionEnum
func GetUpdateInstanceAvailabilityConfigDetailsRecoveryActionEnumValues() []UpdateInstanceAvailabilityConfigDetailsRecoveryActionEnum {
	values := make([]UpdateInstanceAvailabilityConfigDetailsRecoveryActionEnum, 0)
	for _, v := range mappingUpdateInstanceAvailabilityConfigDetailsRecoveryAction {
		values = append(values, v)
	}
	return values
}
