// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// DatabaseToolsConnectionCollection List of DatabaseToolsConnectionSummary items.
type DatabaseToolsConnectionCollection struct {

	// Array of DatabaseToolsConnectionSummary.
	Items []DatabaseToolsConnectionSummary `mandatory:"true" json:"items"`
}

func (m DatabaseToolsConnectionCollection) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DatabaseToolsConnectionCollection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Items []databasetoolsconnectionsummary `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Items = make([]DatabaseToolsConnectionSummary, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(DatabaseToolsConnectionSummary)
		} else {
			m.Items[i] = nil
		}
	}

	return
}
