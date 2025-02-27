// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_file_storage "github.com/oracle/oci-go-sdk/v55/filestorage"

	oci_common "github.com/oracle/oci-go-sdk/v55/common"
)

func init() {
	RegisterOracleClient("oci_file_storage.FileStorageClient", &OracleClient{InitClientFn: initFilestorageFileStorageClient})
}

func initFilestorageFileStorageClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_file_storage.NewFileStorageClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) fileStorageClient() *oci_file_storage.FileStorageClient {
	return m.GetClient("oci_file_storage.FileStorageClient").(*oci_file_storage.FileStorageClient)
}
