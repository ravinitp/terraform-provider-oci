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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// CreateResolverEndpointDetails The body for defining a new resolver endpoint.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateResolverEndpointDetails interface {

	// The name of the resolver endpoint. Must be unique, case-insensitive, within the resolver.
	GetName() *string

	// A Boolean flag indicating whether or not the resolver endpoint is for forwarding.
	GetIsForwarding() *bool

	// A Boolean flag indicating whether or not the resolver endpoint is for listening.
	GetIsListening() *bool

	// An IP address from which forwarded queries may be sent. For VNIC endpoints, this IP address must be part
	// of the subnet and will be assigned by the system if unspecified when isForwarding is true.
	GetForwardingAddress() *string

	// An IP address to listen to queries on. For VNIC endpoints this IP address must be part of the
	// subnet and will be assigned by the system if unspecified when isListening is true.
	GetListeningAddress() *string
}

type createresolverendpointdetails struct {
	JsonData          []byte
	Name              *string `mandatory:"true" json:"name"`
	IsForwarding      *bool   `mandatory:"true" json:"isForwarding"`
	IsListening       *bool   `mandatory:"true" json:"isListening"`
	ForwardingAddress *string `mandatory:"false" json:"forwardingAddress"`
	ListeningAddress  *string `mandatory:"false" json:"listeningAddress"`
	EndpointType      string  `json:"endpointType"`
}

// UnmarshalJSON unmarshals json
func (m *createresolverendpointdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateresolverendpointdetails createresolverendpointdetails
	s := struct {
		Model Unmarshalercreateresolverendpointdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.IsForwarding = s.Model.IsForwarding
	m.IsListening = s.Model.IsListening
	m.ForwardingAddress = s.Model.ForwardingAddress
	m.ListeningAddress = s.Model.ListeningAddress
	m.EndpointType = s.Model.EndpointType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createresolverendpointdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EndpointType {
	case "VNIC":
		mm := CreateResolverVnicEndpointDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m createresolverendpointdetails) GetName() *string {
	return m.Name
}

//GetIsForwarding returns IsForwarding
func (m createresolverendpointdetails) GetIsForwarding() *bool {
	return m.IsForwarding
}

//GetIsListening returns IsListening
func (m createresolverendpointdetails) GetIsListening() *bool {
	return m.IsListening
}

//GetForwardingAddress returns ForwardingAddress
func (m createresolverendpointdetails) GetForwardingAddress() *string {
	return m.ForwardingAddress
}

//GetListeningAddress returns ListeningAddress
func (m createresolverendpointdetails) GetListeningAddress() *string {
	return m.ListeningAddress
}

func (m createresolverendpointdetails) String() string {
	return common.PointerString(m)
}

// CreateResolverEndpointDetailsEndpointTypeEnum Enum with underlying type: string
type CreateResolverEndpointDetailsEndpointTypeEnum string

// Set of constants representing the allowable values for CreateResolverEndpointDetailsEndpointTypeEnum
const (
	CreateResolverEndpointDetailsEndpointTypeVnic CreateResolverEndpointDetailsEndpointTypeEnum = "VNIC"
)

var mappingCreateResolverEndpointDetailsEndpointType = map[string]CreateResolverEndpointDetailsEndpointTypeEnum{
	"VNIC": CreateResolverEndpointDetailsEndpointTypeVnic,
}

// GetCreateResolverEndpointDetailsEndpointTypeEnumValues Enumerates the set of values for CreateResolverEndpointDetailsEndpointTypeEnum
func GetCreateResolverEndpointDetailsEndpointTypeEnumValues() []CreateResolverEndpointDetailsEndpointTypeEnum {
	values := make([]CreateResolverEndpointDetailsEndpointTypeEnum, 0)
	for _, v := range mappingCreateResolverEndpointDetailsEndpointType {
		values = append(values, v)
	}
	return values
}
