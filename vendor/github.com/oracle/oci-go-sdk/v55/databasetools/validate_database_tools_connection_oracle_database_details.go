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

// ValidateDatabaseToolsConnectionOracleDatabaseDetails Connection validation details for Oracle Database.
type ValidateDatabaseToolsConnectionOracleDatabaseDetails struct {
}

func (m ValidateDatabaseToolsConnectionOracleDatabaseDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ValidateDatabaseToolsConnectionOracleDatabaseDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeValidateDatabaseToolsConnectionOracleDatabaseDetails ValidateDatabaseToolsConnectionOracleDatabaseDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeValidateDatabaseToolsConnectionOracleDatabaseDetails
	}{
		"ORACLE_DATABASE",
		(MarshalTypeValidateDatabaseToolsConnectionOracleDatabaseDetails)(m),
	}

	return json.Marshal(&s)
}
