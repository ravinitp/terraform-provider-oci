// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// PatchHistoryEntrySummary The record of a patch action on a specified target.
type PatchHistoryEntrySummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the patch history entry.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the patch.
	PatchId *string `mandatory:"true" json:"patchId"`

	// The current state of the action.
	LifecycleState PatchHistoryEntrySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the patch action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The action being performed or was completed.
	Action PatchHistoryEntrySummaryActionEnum `mandatory:"false" json:"action,omitempty"`

	// A descriptive text associated with the lifecycleState.
	// Typically contains additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the patch action completed
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`
}

func (m PatchHistoryEntrySummary) String() string {
	return common.PointerString(m)
}

// PatchHistoryEntrySummaryActionEnum Enum with underlying type: string
type PatchHistoryEntrySummaryActionEnum string

// Set of constants representing the allowable values for PatchHistoryEntrySummaryActionEnum
const (
	PatchHistoryEntrySummaryActionApply    PatchHistoryEntrySummaryActionEnum = "APPLY"
	PatchHistoryEntrySummaryActionPrecheck PatchHistoryEntrySummaryActionEnum = "PRECHECK"
)

var mappingPatchHistoryEntrySummaryAction = map[string]PatchHistoryEntrySummaryActionEnum{
	"APPLY":    PatchHistoryEntrySummaryActionApply,
	"PRECHECK": PatchHistoryEntrySummaryActionPrecheck,
}

// GetPatchHistoryEntrySummaryActionEnumValues Enumerates the set of values for PatchHistoryEntrySummaryActionEnum
func GetPatchHistoryEntrySummaryActionEnumValues() []PatchHistoryEntrySummaryActionEnum {
	values := make([]PatchHistoryEntrySummaryActionEnum, 0)
	for _, v := range mappingPatchHistoryEntrySummaryAction {
		values = append(values, v)
	}
	return values
}

// PatchHistoryEntrySummaryLifecycleStateEnum Enum with underlying type: string
type PatchHistoryEntrySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for PatchHistoryEntrySummaryLifecycleStateEnum
const (
	PatchHistoryEntrySummaryLifecycleStateInProgress PatchHistoryEntrySummaryLifecycleStateEnum = "IN_PROGRESS"
	PatchHistoryEntrySummaryLifecycleStateSucceeded  PatchHistoryEntrySummaryLifecycleStateEnum = "SUCCEEDED"
	PatchHistoryEntrySummaryLifecycleStateFailed     PatchHistoryEntrySummaryLifecycleStateEnum = "FAILED"
)

var mappingPatchHistoryEntrySummaryLifecycleState = map[string]PatchHistoryEntrySummaryLifecycleStateEnum{
	"IN_PROGRESS": PatchHistoryEntrySummaryLifecycleStateInProgress,
	"SUCCEEDED":   PatchHistoryEntrySummaryLifecycleStateSucceeded,
	"FAILED":      PatchHistoryEntrySummaryLifecycleStateFailed,
}

// GetPatchHistoryEntrySummaryLifecycleStateEnumValues Enumerates the set of values for PatchHistoryEntrySummaryLifecycleStateEnum
func GetPatchHistoryEntrySummaryLifecycleStateEnumValues() []PatchHistoryEntrySummaryLifecycleStateEnum {
	values := make([]PatchHistoryEntrySummaryLifecycleStateEnum, 0)
	for _, v := range mappingPatchHistoryEntrySummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
