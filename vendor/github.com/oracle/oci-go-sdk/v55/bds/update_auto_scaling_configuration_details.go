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

// UpdateAutoScalingConfigurationDetails The information about the autoscale configuration.
type UpdateAutoScalingConfigurationDetails struct {

	// A user-friendly name. The name does not have to be unique, and it may be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether the autoscale configuration is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Base-64 encoded password for the cluster (and Cloudera Manager) admin user.
	ClusterAdminPassword *string `mandatory:"false" json:"clusterAdminPassword"`

	Policy *AutoScalePolicy `mandatory:"false" json:"policy"`
}

func (m UpdateAutoScalingConfigurationDetails) String() string {
	return common.PointerString(m)
}
