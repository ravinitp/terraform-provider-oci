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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// UpdateDhcpDetails The representation of UpdateDhcpDetails
type UpdateDhcpDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	Options []DhcpOption `mandatory:"false" json:"options"`

	// The search domain name type of DHCP options
	DomainNameType UpdateDhcpDetailsDomainNameTypeEnum `mandatory:"false" json:"domainNameType,omitempty"`
}

func (m UpdateDhcpDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDhcpDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags    map[string]map[string]interface{}   `json:"definedTags"`
		DisplayName    *string                             `json:"displayName"`
		FreeformTags   map[string]string                   `json:"freeformTags"`
		Options        []dhcpoption                        `json:"options"`
		DomainNameType UpdateDhcpDetailsDomainNameTypeEnum `json:"domainNameType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.Options = make([]DhcpOption, len(model.Options))
	for i, n := range model.Options {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Options[i] = nn.(DhcpOption)
		} else {
			m.Options[i] = nil
		}
	}

	m.DomainNameType = model.DomainNameType

	return
}

// UpdateDhcpDetailsDomainNameTypeEnum Enum with underlying type: string
type UpdateDhcpDetailsDomainNameTypeEnum string

// Set of constants representing the allowable values for UpdateDhcpDetailsDomainNameTypeEnum
const (
	UpdateDhcpDetailsDomainNameTypeSubnetDomain UpdateDhcpDetailsDomainNameTypeEnum = "SUBNET_DOMAIN"
	UpdateDhcpDetailsDomainNameTypeVcnDomain    UpdateDhcpDetailsDomainNameTypeEnum = "VCN_DOMAIN"
	UpdateDhcpDetailsDomainNameTypeCustomDomain UpdateDhcpDetailsDomainNameTypeEnum = "CUSTOM_DOMAIN"
)

var mappingUpdateDhcpDetailsDomainNameType = map[string]UpdateDhcpDetailsDomainNameTypeEnum{
	"SUBNET_DOMAIN": UpdateDhcpDetailsDomainNameTypeSubnetDomain,
	"VCN_DOMAIN":    UpdateDhcpDetailsDomainNameTypeVcnDomain,
	"CUSTOM_DOMAIN": UpdateDhcpDetailsDomainNameTypeCustomDomain,
}

// GetUpdateDhcpDetailsDomainNameTypeEnumValues Enumerates the set of values for UpdateDhcpDetailsDomainNameTypeEnum
func GetUpdateDhcpDetailsDomainNameTypeEnumValues() []UpdateDhcpDetailsDomainNameTypeEnum {
	values := make([]UpdateDhcpDetailsDomainNameTypeEnum, 0)
	for _, v := range mappingUpdateDhcpDetailsDomainNameType {
		values = append(values, v)
	}
	return values
}
