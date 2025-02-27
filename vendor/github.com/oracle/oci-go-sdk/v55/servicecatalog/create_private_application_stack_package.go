// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// CreatePrivateApplicationStackPackage An object for creating a private application stack package.
type CreatePrivateApplicationStackPackage struct {

	// The package version.
	Version *string `mandatory:"true" json:"version"`

	// Base-64 payload of the Terraform zip package.
	ZipFileBase64Encoded *string `mandatory:"false" json:"zipFileBase64Encoded"`
}

//GetVersion returns Version
func (m CreatePrivateApplicationStackPackage) GetVersion() *string {
	return m.Version
}

func (m CreatePrivateApplicationStackPackage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreatePrivateApplicationStackPackage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreatePrivateApplicationStackPackage CreatePrivateApplicationStackPackage
	s := struct {
		DiscriminatorParam string `json:"packageType"`
		MarshalTypeCreatePrivateApplicationStackPackage
	}{
		"STACK",
		(MarshalTypeCreatePrivateApplicationStackPackage)(m),
	}

	return json.Marshal(&s)
}
