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

// ParavirtualizedVolumeAttachment A paravirtualized volume attachment.
type ParavirtualizedVolumeAttachment struct {

	// The availability domain of an instance.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the volume attachment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the instance the volume is attached to.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The date and time the volume was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the volume.
	VolumeId *string `mandatory:"true" json:"volumeId"`

	// The device name.
	Device *string `mandatory:"false" json:"device"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether the attachment was created in read-only mode.
	IsReadOnly *bool `mandatory:"false" json:"isReadOnly"`

	// Whether the attachment should be created in shareable mode. If an attachment
	// is created in shareable mode, then other instances can attach the same volume, provided
	// that they also create their attachments in shareable mode. Only certain volume types can
	// be attached in shareable mode. Defaults to false if not specified.
	IsShareable *bool `mandatory:"false" json:"isShareable"`

	// Whether in-transit encryption for the data volume's paravirtualized attachment is enabled or not.
	IsPvEncryptionInTransitEnabled *bool `mandatory:"false" json:"isPvEncryptionInTransitEnabled"`

	// Whether the attachment is multipath or not.
	IsMultipath *bool `mandatory:"false" json:"isMultipath"`

	// The current state of the volume attachment.
	LifecycleState VolumeAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The iscsi login state of the volume attachment. For a multipath volume attachment,
	// all iscsi sessions need to be all logged-in or logged-out to be in logged-in or logged-out state.
	IscsiLoginState VolumeAttachmentIscsiLoginStateEnum `mandatory:"false" json:"iscsiLoginState,omitempty"`
}

//GetAvailabilityDomain returns AvailabilityDomain
func (m ParavirtualizedVolumeAttachment) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

//GetCompartmentId returns CompartmentId
func (m ParavirtualizedVolumeAttachment) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDevice returns Device
func (m ParavirtualizedVolumeAttachment) GetDevice() *string {
	return m.Device
}

//GetDisplayName returns DisplayName
func (m ParavirtualizedVolumeAttachment) GetDisplayName() *string {
	return m.DisplayName
}

//GetId returns Id
func (m ParavirtualizedVolumeAttachment) GetId() *string {
	return m.Id
}

//GetInstanceId returns InstanceId
func (m ParavirtualizedVolumeAttachment) GetInstanceId() *string {
	return m.InstanceId
}

//GetIsReadOnly returns IsReadOnly
func (m ParavirtualizedVolumeAttachment) GetIsReadOnly() *bool {
	return m.IsReadOnly
}

//GetIsShareable returns IsShareable
func (m ParavirtualizedVolumeAttachment) GetIsShareable() *bool {
	return m.IsShareable
}

//GetLifecycleState returns LifecycleState
func (m ParavirtualizedVolumeAttachment) GetLifecycleState() VolumeAttachmentLifecycleStateEnum {
	return m.LifecycleState
}

//GetTimeCreated returns TimeCreated
func (m ParavirtualizedVolumeAttachment) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetVolumeId returns VolumeId
func (m ParavirtualizedVolumeAttachment) GetVolumeId() *string {
	return m.VolumeId
}

//GetIsPvEncryptionInTransitEnabled returns IsPvEncryptionInTransitEnabled
func (m ParavirtualizedVolumeAttachment) GetIsPvEncryptionInTransitEnabled() *bool {
	return m.IsPvEncryptionInTransitEnabled
}

//GetIsMultipath returns IsMultipath
func (m ParavirtualizedVolumeAttachment) GetIsMultipath() *bool {
	return m.IsMultipath
}

//GetIscsiLoginState returns IscsiLoginState
func (m ParavirtualizedVolumeAttachment) GetIscsiLoginState() VolumeAttachmentIscsiLoginStateEnum {
	return m.IscsiLoginState
}

func (m ParavirtualizedVolumeAttachment) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ParavirtualizedVolumeAttachment) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeParavirtualizedVolumeAttachment ParavirtualizedVolumeAttachment
	s := struct {
		DiscriminatorParam string `json:"attachmentType"`
		MarshalTypeParavirtualizedVolumeAttachment
	}{
		"paravirtualized",
		(MarshalTypeParavirtualizedVolumeAttachment)(m),
	}

	return json.Marshal(&s)
}
