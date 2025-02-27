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

// Backup The representation of Backup
type Backup struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"false" json:"databaseId"`

	// The user-friendly name for the backup. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The type of backup.
	Type BackupTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The date and time the backup started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the backup was completed.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The name of the availability domain where the database backup is stored.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The current state of the backup.
	LifecycleState BackupLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The Oracle Database edition of the DB system from which the database backup was taken.
	DatabaseEdition BackupDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`

	// The size of the database in gigabytes at the time the backup was taken.
	DatabaseSizeInGBs *float64 `mandatory:"false" json:"databaseSizeInGBs"`

	// Shape of the backup's source database.
	Shape *string `mandatory:"false" json:"shape"`

	// Version of the backup's source database
	Version *string `mandatory:"false" json:"version"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`
}

func (m Backup) String() string {
	return common.PointerString(m)
}

// BackupTypeEnum Enum with underlying type: string
type BackupTypeEnum string

// Set of constants representing the allowable values for BackupTypeEnum
const (
	BackupTypeIncremental BackupTypeEnum = "INCREMENTAL"
	BackupTypeFull        BackupTypeEnum = "FULL"
	BackupTypeVirtualFull BackupTypeEnum = "VIRTUAL_FULL"
)

var mappingBackupType = map[string]BackupTypeEnum{
	"INCREMENTAL":  BackupTypeIncremental,
	"FULL":         BackupTypeFull,
	"VIRTUAL_FULL": BackupTypeVirtualFull,
}

// GetBackupTypeEnumValues Enumerates the set of values for BackupTypeEnum
func GetBackupTypeEnumValues() []BackupTypeEnum {
	values := make([]BackupTypeEnum, 0)
	for _, v := range mappingBackupType {
		values = append(values, v)
	}
	return values
}

// BackupLifecycleStateEnum Enum with underlying type: string
type BackupLifecycleStateEnum string

// Set of constants representing the allowable values for BackupLifecycleStateEnum
const (
	BackupLifecycleStateCreating  BackupLifecycleStateEnum = "CREATING"
	BackupLifecycleStateActive    BackupLifecycleStateEnum = "ACTIVE"
	BackupLifecycleStateDeleting  BackupLifecycleStateEnum = "DELETING"
	BackupLifecycleStateDeleted   BackupLifecycleStateEnum = "DELETED"
	BackupLifecycleStateFailed    BackupLifecycleStateEnum = "FAILED"
	BackupLifecycleStateRestoring BackupLifecycleStateEnum = "RESTORING"
)

var mappingBackupLifecycleState = map[string]BackupLifecycleStateEnum{
	"CREATING":  BackupLifecycleStateCreating,
	"ACTIVE":    BackupLifecycleStateActive,
	"DELETING":  BackupLifecycleStateDeleting,
	"DELETED":   BackupLifecycleStateDeleted,
	"FAILED":    BackupLifecycleStateFailed,
	"RESTORING": BackupLifecycleStateRestoring,
}

// GetBackupLifecycleStateEnumValues Enumerates the set of values for BackupLifecycleStateEnum
func GetBackupLifecycleStateEnumValues() []BackupLifecycleStateEnum {
	values := make([]BackupLifecycleStateEnum, 0)
	for _, v := range mappingBackupLifecycleState {
		values = append(values, v)
	}
	return values
}

// BackupDatabaseEditionEnum Enum with underlying type: string
type BackupDatabaseEditionEnum string

// Set of constants representing the allowable values for BackupDatabaseEditionEnum
const (
	BackupDatabaseEditionStandardEdition                     BackupDatabaseEditionEnum = "STANDARD_EDITION"
	BackupDatabaseEditionEnterpriseEdition                   BackupDatabaseEditionEnum = "ENTERPRISE_EDITION"
	BackupDatabaseEditionEnterpriseEditionHighPerformance    BackupDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	BackupDatabaseEditionEnterpriseEditionExtremePerformance BackupDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingBackupDatabaseEdition = map[string]BackupDatabaseEditionEnum{
	"STANDARD_EDITION":                       BackupDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     BackupDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    BackupDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": BackupDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetBackupDatabaseEditionEnumValues Enumerates the set of values for BackupDatabaseEditionEnum
func GetBackupDatabaseEditionEnumValues() []BackupDatabaseEditionEnum {
	values := make([]BackupDatabaseEditionEnum, 0)
	for _, v := range mappingBackupDatabaseEdition {
		values = append(values, v)
	}
	return values
}
