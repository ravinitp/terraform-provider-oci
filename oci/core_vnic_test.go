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
	vnicSingularDataSourceRepresentation = map[string]interface{}{
		"vnic_id": Representation{RepType: Required, Create: `${lookup(data.oci_core_vnic_attachments.t.vnic_attachments[0],"vnic_id")}`},
	}

	VnicResourceConfig             = ``
	VnicResourceConfigDependencies = OciImageIdsVariable +
		GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, instanceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, RepresentationCopyWithNewProperties(subnetRepresentation, map[string]interface{}{
			"dns_label": Representation{RepType: Required, Create: `dnslabel`},
		})) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, RepresentationCopyWithNewProperties(vcnRepresentation, map[string]interface{}{
			"dns_label": Representation{RepType: Required, Create: `dnslabel`},
		})) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: core/virtualNetwork
func TestCoreVnicResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreVnicResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_core_vnic.test_vnic"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + `

data "oci_core_vnic_attachments" "t" {
	compartment_id = "${var.compartment_id}"
	instance_id = "${oci_core_instance.test_instance.id}"
}` +
				GenerateDataSourceFromRepresentationMap("oci_core_vnic", "test_vnic", Required, Create, vnicSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VnicResourceConfig + VnicResourceConfigDependencies,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vnic_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hostname_label"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_primary"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mac_address"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nsg_ids.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "private_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "public_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "skip_source_dest_check"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
	})
}
