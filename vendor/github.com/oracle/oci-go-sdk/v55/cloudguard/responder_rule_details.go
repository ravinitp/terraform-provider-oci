// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// ResponderRuleDetails Details of ResponderRule.
type ResponderRuleDetails struct {

	// Identifies state for ResponderRule
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	Condition Condition `mandatory:"false" json:"condition"`

	// ResponderRule configurations
	Configurations []ResponderConfiguration `mandatory:"false" json:"configurations"`

	// Execution Mode for ResponderRule
	Mode ResponderModeTypesEnum `mandatory:"false" json:"mode,omitempty"`
}

func (m ResponderRuleDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ResponderRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Condition      condition                `json:"condition"`
		Configurations []ResponderConfiguration `json:"configurations"`
		Mode           ResponderModeTypesEnum   `json:"mode"`
		IsEnabled      *bool                    `json:"isEnabled"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Condition.UnmarshalPolymorphicJSON(model.Condition.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Condition = nn.(Condition)
	} else {
		m.Condition = nil
	}

	m.Configurations = make([]ResponderConfiguration, len(model.Configurations))
	for i, n := range model.Configurations {
		m.Configurations[i] = n
	}

	m.Mode = model.Mode

	m.IsEnabled = model.IsEnabled

	return
}
