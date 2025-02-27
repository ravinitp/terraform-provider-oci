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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// UpdateHostInsightDetails The information to be updated.
type UpdateHostInsightDetails interface {

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatehostinsightdetails struct {
	JsonData     []byte
	FreeformTags map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags  map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	EntitySource string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *updatehostinsightdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatehostinsightdetails updatehostinsightdetails
	s := struct {
		Model Unmarshalerupdatehostinsightdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatehostinsightdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "EM_MANAGED_EXTERNAL_HOST":
		mm := UpdateEmManagedExternalHostInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MACS_MANAGED_EXTERNAL_HOST":
		mm := UpdateMacsManagedExternalHostInsightDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetFreeformTags returns FreeformTags
func (m updatehostinsightdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m updatehostinsightdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatehostinsightdetails) String() string {
	return common.PointerString(m)
}
