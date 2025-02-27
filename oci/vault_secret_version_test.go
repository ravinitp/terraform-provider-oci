// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	secretVersionSingularDataSourceRepresentation = map[string]interface{}{
		"secret_id":             Representation{RepType: Required, Create: `${oci_vault_secret.test_secret.id`},
		"secret_version_number": Representation{RepType: Required, Create: `1`},
	}

	SecretVersionResourceConfig = ``
)

// issue-routing-tag: vault/default
func TestVaultSecretVersionResource_basic(t *testing.T) {
	t.Skip("Skip this test till Secret Management service provides a better way of testing this.")
	httpreplay.SetScenario("TestVaultSecretVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_vault_secret_version.test_secret_version"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_vault_secret_version", "test_secret_version", Required, Create, secretVersionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SecretVersionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "secret_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "secret_version_number", "1"),

				//resource.TestCheckResourceAttrSet(singularDatasourceName, "content_type"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stages.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_current_version_expiry"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_of_deletion"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version_number"),
			),
		},
	})
}
