// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v55/common"
)

// DeploymentExecutionProgress The execution progress details of a deployment.
type DeploymentExecutionProgress struct {

	// Time the deployment is started. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Time the deployment is finished. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Map of stage OCIDs to deploy stage execution progress model.
	DeployStageExecutionProgress map[string]DeployStageExecutionProgress `mandatory:"false" json:"deployStageExecutionProgress"`
}

func (m DeploymentExecutionProgress) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *DeploymentExecutionProgress) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeStarted                  *common.SDKTime                         `json:"timeStarted"`
		TimeFinished                 *common.SDKTime                         `json:"timeFinished"`
		DeployStageExecutionProgress map[string]deploystageexecutionprogress `json:"deployStageExecutionProgress"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeStarted = model.TimeStarted

	m.TimeFinished = model.TimeFinished

	m.DeployStageExecutionProgress = make(map[string]DeployStageExecutionProgress)
	for k, v := range model.DeployStageExecutionProgress {
		nn, e = v.UnmarshalPolymorphicJSON(v.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DeployStageExecutionProgress[k] = nn.(DeployStageExecutionProgress)
		} else {
			m.DeployStageExecutionProgress[k] = nil
		}
	}

	return
}
