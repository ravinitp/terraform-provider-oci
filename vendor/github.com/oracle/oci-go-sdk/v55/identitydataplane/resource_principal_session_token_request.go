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

// ResourcePrincipalSessionTokenRequest The representation of ResourcePrincipalSessionTokenRequest
type ResourcePrincipalSessionTokenRequest struct {

	// The resource principal token.
	ResourcePrincipalToken *string `mandatory:"true" json:"resourcePrincipalToken"`

	// The service principal session token.
	ServicePrincipalSessionToken *string `mandatory:"true" json:"servicePrincipalSessionToken"`

	// The session public key.
	SessionPublicKey *string `mandatory:"true" json:"sessionPublicKey"`
}

func (m ResourcePrincipalSessionTokenRequest) String() string {
	return common.PointerString(m)
}
