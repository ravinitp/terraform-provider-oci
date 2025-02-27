// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_operator_access_control "github.com/oracle/oci-go-sdk/v55/operatoraccesscontrol"

	oci_common "github.com/oracle/oci-go-sdk/v55/common"
)

func init() {
	RegisterOracleClient("oci_operator_access_control.AccessRequestsClient", &OracleClient{InitClientFn: initOperatoraccesscontrolAccessRequestsClient})
	RegisterOracleClient("oci_operator_access_control.OperatorActionsClient", &OracleClient{InitClientFn: initOperatoraccesscontrolOperatorActionsClient})
	RegisterOracleClient("oci_operator_access_control.OperatorControlClient", &OracleClient{InitClientFn: initOperatoraccesscontrolOperatorControlClient})
	RegisterOracleClient("oci_operator_access_control.OperatorControlAssignmentClient", &OracleClient{InitClientFn: initOperatoraccesscontrolOperatorControlAssignmentClient})
}

func initOperatoraccesscontrolAccessRequestsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_operator_access_control.NewAccessRequestsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) accessRequestsClient() *oci_operator_access_control.AccessRequestsClient {
	return m.GetClient("oci_operator_access_control.AccessRequestsClient").(*oci_operator_access_control.AccessRequestsClient)
}

func initOperatoraccesscontrolOperatorActionsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_operator_access_control.NewOperatorActionsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) operatorActionsClient() *oci_operator_access_control.OperatorActionsClient {
	return m.GetClient("oci_operator_access_control.OperatorActionsClient").(*oci_operator_access_control.OperatorActionsClient)
}

func initOperatoraccesscontrolOperatorControlClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_operator_access_control.NewOperatorControlClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) operatorControlClient() *oci_operator_access_control.OperatorControlClient {
	return m.GetClient("oci_operator_access_control.OperatorControlClient").(*oci_operator_access_control.OperatorControlClient)
}

func initOperatoraccesscontrolOperatorControlAssignmentClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_operator_access_control.NewOperatorControlAssignmentClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) operatorControlAssignmentClient() *oci_operator_access_control.OperatorControlAssignmentClient {
	return m.GetClient("oci_operator_access_control.OperatorControlAssignmentClient").(*oci_operator_access_control.OperatorControlAssignmentClient)
}
