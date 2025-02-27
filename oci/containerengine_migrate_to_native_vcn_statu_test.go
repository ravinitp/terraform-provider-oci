// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	migrateToNativeVCNSingularDataSourceRepresentation = map[string]interface{}{
		"cluster_id": Representation{RepType: Required, Create: `${oci_containerengine_cluster.test_cluster.id}`},
	}
)

// issue-routing-tag: containerengine/default
func TestContainerengineMigrateToNativeVcnStatusResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestContainerengineMigrateToNativeVcnStatusResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_containerengine_cluster.test_cluster"
	singularDatasourceName := "data.oci_containerengine_migrate_to_native_vcn_status.test_migrate_to_native_vcn_status"

	var resId, resId2 string

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// Create V1 Cluster
		{
			Config: config + compartmentIdVariableStr + ClusterResourceDependencies + GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Required, Create, RepresentationCopyWithRemovedProperties(clusterRepresentation, []string{"kms_key_id", "options", "image_policy_config"})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify V1 Cluster migrates to V2
		{
			Config: config + compartmentIdVariableStr + ClusterResourceDependencies + GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Optional, Update, RepresentationCopyWithRemovedProperties(clusterRepresentation, []string{"kms_key_id", "options", "image_policy_config"})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "endpoint_config.0.is_public_ip_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_config.0.subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "endpoint_config.0.nsg_ids.#"),
				resource.TestCheckResourceAttrSet(resourceName, "kubernetes_version"),
				resource.TestCheckResourceAttrSet(resourceName, "vcn_id"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + ClusterResourceDependencies + GenerateResourceFromRepresentationMap("oci_containerengine_cluster", "test_cluster", Optional, Update, RepresentationCopyWithRemovedProperties(clusterRepresentation, []string{"kms_key_id", "options", "image_policy_config"})) + GenerateDataSourceFromRepresentationMap(
				"oci_containerengine_migrate_to_native_vcn_status", "test_migrate_to_native_vcn_status",
				Optional, Create, migrateToNativeVCNSingularDataSourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_decommission_scheduled"),
			),
		},
	})
}
