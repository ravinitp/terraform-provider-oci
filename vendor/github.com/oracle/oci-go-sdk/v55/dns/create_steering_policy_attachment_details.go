// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// CreateSteeringPolicyAttachmentDetails The body for defining an attachment between a steering policy and a domain.
//
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateSteeringPolicyAttachmentDetails struct {

	// The OCID of the attached steering policy.
	SteeringPolicyId *string `mandatory:"true" json:"steeringPolicyId"`

	// The OCID of the attached zone.
	ZoneId *string `mandatory:"true" json:"zoneId"`

	// The attached domain within the attached zone.
	DomainName *string `mandatory:"true" json:"domainName"`

	// A user-friendly name for the steering policy attachment.
	// Does not have to be unique and can be changed.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateSteeringPolicyAttachmentDetails) String() string {
	return common.PointerString(m)
}
