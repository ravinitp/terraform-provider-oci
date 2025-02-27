// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// OperationsInsightsWarehouseSummary Summary of a Operations Insights Warehouse resource.
type OperationsInsightsWarehouseSummary struct {

	// OPSI Warehouse OCID
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// User-friedly name of Operations Insights Warehouse that does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Number of OCPUs allocated to OPSI Warehouse ADW.
	CpuAllocated *float64 `mandatory:"true" json:"cpuAllocated"`

	// The time at which the resource was first created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time at which the resource was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Possible lifecycle states
	LifecycleState OperationsInsightsWarehouseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Number of OCPUs used by OPSI Warehouse ADW. Can be fractional.
	CpuUsed *float64 `mandatory:"false" json:"cpuUsed"`

	// Storage allocated to OPSI Warehouse ADW.
	StorageAllocatedInGBs *float64 `mandatory:"false" json:"storageAllocatedInGBs"`

	// Storage by OPSI Warehouse ADW in GB.
	StorageUsedInGBs *float64 `mandatory:"false" json:"storageUsedInGBs"`

	// OCID of the dynamic group created for the warehouse
	DynamicGroupId *string `mandatory:"false" json:"dynamicGroupId"`

	// Tenancy Identifier of Operations Insights service
	OperationsInsightsTenancyId *string `mandatory:"false" json:"operationsInsightsTenancyId"`

	// The time at which the ADW wallet was last rotated for the Operations Insights Warehouse. An RFC3339 formatted datetime string
	TimeLastWalletRotated *common.SDKTime `mandatory:"false" json:"timeLastWalletRotated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m OperationsInsightsWarehouseSummary) String() string {
	return common.PointerString(m)
}
