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

// CloudAutonomousVmCluster Details of the cloud Autonomous VM cluster.
type CloudAutonomousVmCluster struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Cloud Autonomous VM cluster.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the availability domain that the cloud Autonomous VM cluster is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the cloud Autonomous VM Cluster is associated with.
	// **Subnet Restrictions:**
	// - For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and backup subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The current state of the cloud Autonomous VM cluster.
	LifecycleState CloudAutonomousVmClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-friendly name for the cloud Autonomous VM cluster. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
	CloudExadataInfrastructureId *string `mandatory:"true" json:"cloudExadataInfrastructureId"`

	// User defined description of the cloud Autonomous VM cluster.
	Description *string `mandatory:"false" json:"description"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last maintenance update history. This value is updated when a maintenance update starts.
	LastUpdateHistoryEntryId *string `mandatory:"false" json:"lastUpdateHistoryEntryId"`

	// The date and time that the cloud Autonomous VM cluster was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last date and time that the cloud Autonomous VM cluster was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The hostname for the cloud Autonomous VM cluster.
	Hostname *string `mandatory:"false" json:"hostname"`

	// The domain name for the cloud Autonomous VM cluster.
	Domain *string `mandatory:"false" json:"domain"`

	// The model name of the Exadata hardware running the cloud Autonomous VM cluster.
	Shape *string `mandatory:"false" json:"shape"`

	// The number of database servers in the cloud VM cluster.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The total data storage allocated, in terabytes (TB).
	DataStorageSizeInTBs *float64 `mandatory:"false" json:"dataStorageSizeInTBs"`

	// The total data storage allocated, in gigabytes (GB).
	DataStorageSizeInGBs *float64 `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The number of CPU cores enabled on the cloud Autonomous VM cluster.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The number of CPU cores enabled on the cloud Autonomous VM cluster. Only 1 decimal place is allowed for the fractional part.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// The memory allocated in GBs.
	MemorySizeInGBs *int `mandatory:"false" json:"memorySizeInGBs"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm), this attribute must be null because the attribute is already set at the
	// Autonomous Exadata Infrastructure level. When using shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`.
	LicenseModel CloudAutonomousVmClusterLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last maintenance run.
	LastMaintenanceRunId *string `mandatory:"false" json:"lastMaintenanceRunId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the next maintenance run.
	NextMaintenanceRunId *string `mandatory:"false" json:"nextMaintenanceRunId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CloudAutonomousVmCluster) String() string {
	return common.PointerString(m)
}

// CloudAutonomousVmClusterLifecycleStateEnum Enum with underlying type: string
type CloudAutonomousVmClusterLifecycleStateEnum string

// Set of constants representing the allowable values for CloudAutonomousVmClusterLifecycleStateEnum
const (
	CloudAutonomousVmClusterLifecycleStateProvisioning          CloudAutonomousVmClusterLifecycleStateEnum = "PROVISIONING"
	CloudAutonomousVmClusterLifecycleStateAvailable             CloudAutonomousVmClusterLifecycleStateEnum = "AVAILABLE"
	CloudAutonomousVmClusterLifecycleStateUpdating              CloudAutonomousVmClusterLifecycleStateEnum = "UPDATING"
	CloudAutonomousVmClusterLifecycleStateTerminating           CloudAutonomousVmClusterLifecycleStateEnum = "TERMINATING"
	CloudAutonomousVmClusterLifecycleStateTerminated            CloudAutonomousVmClusterLifecycleStateEnum = "TERMINATED"
	CloudAutonomousVmClusterLifecycleStateFailed                CloudAutonomousVmClusterLifecycleStateEnum = "FAILED"
	CloudAutonomousVmClusterLifecycleStateMaintenanceInProgress CloudAutonomousVmClusterLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingCloudAutonomousVmClusterLifecycleState = map[string]CloudAutonomousVmClusterLifecycleStateEnum{
	"PROVISIONING":            CloudAutonomousVmClusterLifecycleStateProvisioning,
	"AVAILABLE":               CloudAutonomousVmClusterLifecycleStateAvailable,
	"UPDATING":                CloudAutonomousVmClusterLifecycleStateUpdating,
	"TERMINATING":             CloudAutonomousVmClusterLifecycleStateTerminating,
	"TERMINATED":              CloudAutonomousVmClusterLifecycleStateTerminated,
	"FAILED":                  CloudAutonomousVmClusterLifecycleStateFailed,
	"MAINTENANCE_IN_PROGRESS": CloudAutonomousVmClusterLifecycleStateMaintenanceInProgress,
}

// GetCloudAutonomousVmClusterLifecycleStateEnumValues Enumerates the set of values for CloudAutonomousVmClusterLifecycleStateEnum
func GetCloudAutonomousVmClusterLifecycleStateEnumValues() []CloudAutonomousVmClusterLifecycleStateEnum {
	values := make([]CloudAutonomousVmClusterLifecycleStateEnum, 0)
	for _, v := range mappingCloudAutonomousVmClusterLifecycleState {
		values = append(values, v)
	}
	return values
}

// CloudAutonomousVmClusterLicenseModelEnum Enum with underlying type: string
type CloudAutonomousVmClusterLicenseModelEnum string

// Set of constants representing the allowable values for CloudAutonomousVmClusterLicenseModelEnum
const (
	CloudAutonomousVmClusterLicenseModelLicenseIncluded     CloudAutonomousVmClusterLicenseModelEnum = "LICENSE_INCLUDED"
	CloudAutonomousVmClusterLicenseModelBringYourOwnLicense CloudAutonomousVmClusterLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCloudAutonomousVmClusterLicenseModel = map[string]CloudAutonomousVmClusterLicenseModelEnum{
	"LICENSE_INCLUDED":       CloudAutonomousVmClusterLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CloudAutonomousVmClusterLicenseModelBringYourOwnLicense,
}

// GetCloudAutonomousVmClusterLicenseModelEnumValues Enumerates the set of values for CloudAutonomousVmClusterLicenseModelEnum
func GetCloudAutonomousVmClusterLicenseModelEnumValues() []CloudAutonomousVmClusterLicenseModelEnum {
	values := make([]CloudAutonomousVmClusterLicenseModelEnum, 0)
	for _, v := range mappingCloudAutonomousVmClusterLicenseModel {
		values = append(values, v)
	}
	return values
}
