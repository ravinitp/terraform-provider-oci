// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// JsonFormatAttribute The JSON file format attribute.
type JsonFormatAttribute struct {

	// Defines whether a file pattern is supported.
	IsFilePattern *bool `mandatory:"false" json:"isFilePattern"`

	// The encoding for the file.
	Encoding *string `mandatory:"false" json:"encoding"`
}

//GetIsFilePattern returns IsFilePattern
func (m JsonFormatAttribute) GetIsFilePattern() *bool {
	return m.IsFilePattern
}

func (m JsonFormatAttribute) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m JsonFormatAttribute) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJsonFormatAttribute JsonFormatAttribute
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeJsonFormatAttribute
	}{
		"JSON_FORMAT",
		(MarshalTypeJsonFormatAttribute)(m),
	}

	return json.Marshal(&s)
}
