// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v55/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v55/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	hostnameDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           RepresentationGroup{Required, hostnameDataSourceFilterRepresentation}}
	hostnameDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `name`},
		"values": Representation{RepType: Required, Create: []string{`${oci_load_balancer_hostname.test_hostname.name}`}},
	}

	hostnameRepresentation = map[string]interface{}{
		"hostname":         Representation{RepType: Required, Create: `app.example.com`, Update: `hostname2`},
		"load_balancer_id": Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             Representation{RepType: Required, Create: `example_hostname_001`},
	}

	HostnameResourceDependencies = GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", Required, Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerHostnameResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerHostnameResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_hostname.test_hostname"
	datasourceName := "data.oci_load_balancer_hostnames.test_hostnames"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+HostnameResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_load_balancer_hostname", "test_hostname", Required, Create, hostnameRepresentation), "loadbalancer", "hostname", t)

	ResourceTest(t, testAccCheckLoadBalancerHostnameDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + HostnameResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_hostname", "test_hostname", Required, Create, hostnameRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "hostname", "app.example.com"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_hostname_001"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + HostnameResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_hostname", "test_hostname", Optional, Update, hostnameRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "hostname", "hostname2"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_hostname_001"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_load_balancer_hostnames", "test_hostnames", Optional, Update, hostnameDataSourceRepresentation) +
				compartmentIdVariableStr + HostnameResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_load_balancer_hostname", "test_hostname", Optional, Update, hostnameRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttr(datasourceName, "hostnames.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "hostnames.0.hostname", "hostname2"),
				resource.TestCheckResourceAttr(datasourceName, "hostnames.0.name", "example_hostname_001"),
			),
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"state",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckLoadBalancerHostnameDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_hostname" {
			noResourceFound = false
			request := oci_load_balancer.GetHostnameRequest{}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.Name = &value
			}

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "load_balancer")

			_, err := client.GetHostname(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("LoadBalancerHostname") {
		resource.AddTestSweepers("LoadBalancerHostname", &resource.Sweeper{
			Name:         "LoadBalancerHostname",
			Dependencies: DependencyGraph["hostname"],
			F:            sweepLoadBalancerHostnameResource,
		})
	}
}

func sweepLoadBalancerHostnameResource(compartment string) error {
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()
	hostnameIds, err := getHostnameIds(compartment)
	if err != nil {
		return err
	}
	for _, hostnameId := range hostnameIds {
		if ok := SweeperDefaultResourceId[hostnameId]; !ok {
			deleteHostnameRequest := oci_load_balancer.DeleteHostnameRequest{}

			deleteHostnameRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "load_balancer")
			_, error := loadBalancerClient.DeleteHostname(context.Background(), deleteHostnameRequest)
			if error != nil {
				fmt.Printf("Error deleting Hostname %s %s, It is possible that the resource is already deleted. Please verify manually \n", hostnameId, error)
				continue
			}
		}
	}
	return nil
}

func getHostnameIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "HostnameId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loadBalancerClient := GetTestClients(&schema.ResourceData{}).loadBalancerClient()

	listHostnamesRequest := oci_load_balancer.ListHostnamesRequest{}

	loadBalancerIds, error := getLoadBalancerIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting loadBalancerId required for Hostname resource requests \n")
	}
	for _, loadBalancerId := range loadBalancerIds {
		listHostnamesRequest.LoadBalancerId = &loadBalancerId

		listHostnamesResponse, err := loadBalancerClient.ListHostnames(context.Background(), listHostnamesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting Hostname list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, hostname := range listHostnamesResponse.Items {
			id := *hostname.Name
			resourceIds = append(resourceIds, id)
			AddResourceIdToSweeperResourceIdMap(compartmentId, "HostnameId", id)
		}

	}
	return resourceIds, nil
}
