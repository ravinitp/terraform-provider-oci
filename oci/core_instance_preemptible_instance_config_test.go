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
	instanceWithPreemptibleInstanceConfigRepresentation = map[string]interface{}{
		"availability_domain":         Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":              Representation{RepType: Required, Create: `${var.compartment_id}`},
		"shape":                       Representation{RepType: Required, Create: `VM.Standard2.1`},
		"image":                       Representation{RepType: Required, Create: `${var.InstanceImageOCID[var.region]}`},
		"preemptible_instance_config": RepresentationGroup{Required, instancePreemptibleInstanceConfigRepresentation},
		"subnet_id":                   Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
	}

	InstanceWithPreemptibleInstanceConfigResourceConfig = InstanceResourceDependenciesWithoutDHV +
		GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceWithPreemptibleInstanceConfigRepresentation)

	instanceWithPreemtibleInstanceConfigDataSourceRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"availability_domain": Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"state":               Representation{RepType: Required, Create: `RUNNING`},
		"filter":              RepresentationGroup{Required, instanceWithPreemtibleInstanceConfigDataSourceFilterRepresentation}}
	instanceWithPreemtibleInstanceConfigDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_core_instance.test_instance.id}`}},
	}

	instancePreemptibleInstanceConfigRepresentation = map[string]interface{}{
		"preemption_action": RepresentationGroup{Required, instancePreemptibleInstanceConfigPreemptionActionRepresentation},
	}
	instancePreemptibleInstanceConfigPreemptionActionRepresentation = map[string]interface{}{
		"type":                 Representation{RepType: Required, Create: `TERMINATE`},
		"preserve_boot_volume": Representation{RepType: Required, Create: `false`},
	}
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestResourceCoreInstancePreemptibleInstanceConfig_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreInstancePreemptibleInstanceConfigResource_basic")
	defer httpreplay.SaveScenario()

	config := `
		provider oci {
			test_time_maintenance_reboot_due = "2030-01-01 00:00:00"
		}
	` + commonTestVariables()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_instance.test_instance"
	datasourceName := "data.oci_core_instances.test_instances"

	ResourceTest(t, testAccCheckCoreInstanceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: testProviderConfig() + compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV +
				GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceWithPreemptibleInstanceConfigRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttr(resourceName, "time_maintenance_reboot_due", ""),
				resource.TestCheckResourceAttr(resourceName, "preemptible_instance_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "preemptible_instance_config.0.preemption_action.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "preemptible_instance_config.0.preemption_action.0.preserve_boot_volume", "false"),
				resource.TestCheckResourceAttr(resourceName, "preemptible_instance_config.0.preemption_action.0.type", "TERMINATE"),

				func(s *terraform.State) (err error) {
					_, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_core_instances", "test_instances", Required, Create, instanceWithPreemtibleInstanceConfigDataSourceRepresentation) +
				compartmentIdVariableStr + InstanceResourceDependenciesWithoutDHV +
				GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceWithPreemptibleInstanceConfigRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "state", "RUNNING"),

				resource.TestCheckResourceAttr(datasourceName, "instances.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.image"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.preemptible_instance_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.preemptible_instance_config.0.preemption_action.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.preemptible_instance_config.0.preemption_action.0.preserve_boot_volume", "false"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.preemptible_instance_config.0.preemption_action.0.type", "TERMINATE"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.region"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.shape", "VM.Standard2.1"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.shape_config.0.ocpus", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.shape_config.0.processor_description"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.source_details.0.source_id"),
				resource.TestCheckResourceAttr(datasourceName, "instances.0.source_details.0.source_type", "image"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "instances.0.time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + InstanceWithPreemptibleInstanceConfigResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
