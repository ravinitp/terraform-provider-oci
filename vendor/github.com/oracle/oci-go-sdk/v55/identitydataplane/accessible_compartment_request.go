// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// AccessibleCompartmentRequest The representation of AccessibleCompartmentRequest
type AccessibleCompartmentRequest struct {

	// The authorization request.
	AuthorizationRequest *AuthorizationRequest `mandatory:"true" json:"authorizationRequest"`

	// The list of compartment ids.
	CompartmentIds []string `mandatory:"true" json:"compartmentIds"`
}

func (m AccessibleCompartmentRequest) String() string {
	return common.PointerString(m)
}
