// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v55/common"
)

// RepositoryPathSummary Object containing information about files and directories in a repository.
type RepositoryPathSummary struct {

	// File or directory.
	Type *string `mandatory:"false" json:"type"`

	// Size of file or directory.
	SizeInBytes *int64 `mandatory:"false" json:"sizeInBytes"`

	// Name of file or directory.
	Name *string `mandatory:"false" json:"name"`

	// Path to file or directory in a repository.
	Path *string `mandatory:"false" json:"path"`

	// SHA-1 checksum of blob or tree.
	Sha *string `mandatory:"false" json:"sha"`

	// The git URL of the submodule.
	SubmoduleGitUrl *string `mandatory:"false" json:"submoduleGitUrl"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m RepositoryPathSummary) String() string {
	return common.PointerString(m)
}
