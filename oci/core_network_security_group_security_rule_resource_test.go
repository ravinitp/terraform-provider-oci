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
	networkSecurityGroupSecurityRuleResourceRepresentation = map[string]interface{}{
		"network_security_group_id": Representation{RepType: Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
		"direction":                 Representation{RepType: Required, Create: `EGRESS`},
		"protocol":                  Representation{RepType: Required, Create: `1`},
		"description":               Representation{RepType: Optional, Create: `description`, Update: `updated description`},
		"destination":               Representation{RepType: Optional, Create: `10.0.0.0/24`},
	}

	networkSecurityGroupIngressSecurityRuleResourceRepresentation = map[string]interface{}{
		"network_security_group_id": Representation{RepType: Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
		"direction":                 Representation{RepType: Required, Create: `INGRESS`},
		"protocol":                  Representation{RepType: Required, Create: `1`},
		"source":                    Representation{RepType: Optional, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"icmp_options":              RepresentationGroup{Optional, nsgSecurityRulesIcmpOptionsRepresentation},
		"source_type":               Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":                 Representation{RepType: Optional, Create: `false`, Update: `true`},
	}

	networkSecurityGroupIngressSecurityRuleUDPResourceRepresentation = map[string]interface{}{
		"network_security_group_id": Representation{RepType: Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
		"direction":                 Representation{RepType: Required, Create: `INGRESS`},
		"protocol":                  Representation{RepType: Required, Create: `17`},
		"source":                    Representation{RepType: Optional, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"source_type":               Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":                 Representation{RepType: Optional, Create: `false`},
		"udp_options":               RepresentationGroup{Optional, securityRulesUdpOptionsRepresentation},
	}

	nsgSecurityRulesIcmpOptionsRepresentation = map[string]interface{}{
		"type": Representation{RepType: Required, Create: `3`},
	}
)

// issue-routing-tag: core/virtualNetwork
func TestAccResourceCoreNetworkSecurityGroupSecurityRule_scenarios(t *testing.T) {
	httpreplay.SetScenario("TestAccResourceCoreNetworkSecurityGroupSecurityRule_multipleRules")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_network_security_group_security_rule.test_network_security_group_security_rule"

	var resId1, resId2 [10]string

	ResourceTest(t, nil, []resource.TestStep{

		//verify Create 10 rules
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", Optional, Create,
					RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleResourceRepresentation, map[string]interface{}{
						"count": Representation{RepType: Optional, Create: `10`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					for i := 0; i < 10; i++ {
						resId, err := FromInstanceState(s, fmt.Sprintf("%s.%d", resourceName, i), "id")
						if resId == "" {
							return err
						}
						resId1[i] = resId
					}
					return nil
				},
			),
		},
		//verify Update 10 rules
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", Optional, Update,
					RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleResourceRepresentation, map[string]interface{}{
						"count": Representation{RepType: Optional, Create: `10`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					for i := 0; i < 10; i++ {

						resId, err := FromInstanceState(s, fmt.Sprintf("%s.%d", resourceName, i), "id")
						if resId == "" {
							return err
						}
						resId2[i] = resId

						if resId1[i] != resId2[i] {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						description, err := FromInstanceState(s, fmt.Sprintf("%s.%d", resourceName, i), "description")
						if description == "" {
							return err
						}
						if description != "updated description" {
							return fmt.Errorf("%s: Attribute 'description' expected \"updated description\", got %s", fmt.Sprintf("%s.%d", resourceName, i), description)
						}
					}
					return nil
				},
			),
		},
		// delete
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies,
		},
		// Create rule without specifying `code` in icmp options
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", Optional, Create, networkSecurityGroupIngressSecurityRuleResourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "icmp_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "icmp_options.0.code", "-1"),
			),
		},
		// Update rule without specifying code in icmp options
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", Optional, Update, networkSecurityGroupIngressSecurityRuleResourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "icmp_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "icmp_options.0.code", "-1"),
			),
		},
		// delete
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies,
		},
		// Create rule without specifying `code` in udp options
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", Optional, Create, networkSecurityGroupIngressSecurityRuleUDPResourceRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "udp_options.#", "1"),
			),
		},
	})
}
