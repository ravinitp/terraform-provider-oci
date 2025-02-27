// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v55/common"
	oci_devops "github.com/oracle/oci-go-sdk/v55/devops"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DeployEnvironmentRequiredOnlyResource = DeployEnvironmentResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Required, Create, deployEnvironmentRepresentation)

	DeployEnvironmentResourceConfig = DeployEnvironmentResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Optional, Update, deployEnvironmentRepresentation)

	deployEnvironmentSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_environment_id": Representation{RepType: Required, Create: `${oci_devops_deploy_environment.test_deploy_environment.id}`},
	}

	deployEnvironmentDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":             Representation{RepType: Optional, Create: `${oci_devops_deploy_environment.test_deploy_environment.id}`},
		"project_id":     Representation{RepType: Optional, Create: `${oci_devops_project.test_project.id}`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, deployEnvironmentDataSourceFilterRepresentation}}
	deployEnvironmentDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_devops_deploy_environment.test_deploy_environment.id}`}},
	}

	cluster_fake_id                 = "ocid1.cluster.oc1.us-ashburn-1.aaaaaaaaafqtkm3fg4zwgnlggmywkzdemi2dcyzymfrdqojygcstocluster1"
	deployEnvironmentRepresentation = map[string]interface{}{
		"deploy_environment_type": Representation{RepType: Required, Create: `OKE_CLUSTER`},
		"project_id":              Representation{RepType: Required, Create: `${oci_devops_project.test_project.id}`},
		"cluster_id":              Representation{RepType: Required, Create: cluster_fake_id},
		"defined_tags":            Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"display_name":            Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":           Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":               RepresentationGroup{Required, ignoreDefinedTagsDifferencesRepresentation},
	}

	DeployEnvironmentResourceDependencies = AvailabilityDomainConfig +
		GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		GenerateResourceFromRepresentationMap("oci_logging_log_group", "test_devops_log_group", Required, Create, devopsLogGroupRepresentation) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: devops/default
func TestDevopsDeployEnvironmentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployEnvironmentResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_environment.test_deploy_environment"
	datasourceName := "data.oci_devops_deploy_environments.test_deploy_environments"
	singularDatasourceName := "data.oci_devops_deploy_environment.test_deploy_environment"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DeployEnvironmentResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Optional, Create, deployEnvironmentRepresentation), "devops", "deployEnvironment", t)

	ResourceTest(t, testAccCheckDevopsDeployEnvironmentDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeployEnvironmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Required, Create, deployEnvironmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "OKE_CLUSTER"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DeployEnvironmentResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DeployEnvironmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Optional, Create, deployEnvironmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "OKE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
			Config: config + compartmentIdVariableStr + DeployEnvironmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Optional, Update, deployEnvironmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_environment_type", "OKE_CLUSTER"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

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
				GenerateDataSourceFromRepresentationMap("oci_devops_deploy_environments", "test_deploy_environments", Optional, Update, deployEnvironmentDataSourceRepresentation) +
				compartmentIdVariableStr + DeployEnvironmentResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Optional, Update, deployEnvironmentRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "deploy_environment_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_environment", Required, Create, deployEnvironmentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployEnvironmentResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_environment_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_environment_type", "OKE_CLUSTER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DeployEnvironmentResourceConfig,
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

func testAccCheckDevopsDeployEnvironmentDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).devopsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_devops_deploy_environment" {
			noResourceFound = false
			request := oci_devops.GetDeployEnvironmentRequest{}

			tmp := rs.Primary.ID
			request.DeployEnvironmentId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")

			response, err := client.GetDeployEnvironment(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_devops.DeployEnvironmentLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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
	if !InSweeperExcludeList("DevopsDeployEnvironment") {
		resource.AddTestSweepers("DevopsDeployEnvironment", &resource.Sweeper{
			Name:         "DevopsDeployEnvironment",
			Dependencies: DependencyGraph["deployEnvironment"],
			F:            sweepDevopsDeployEnvironmentResource,
		})
	}
}

func sweepDevopsDeployEnvironmentResource(compartment string) error {
	deployEnvironmentClient := GetTestClients(&schema.ResourceData{}).devopsClient()
	deployEnvironmentIds, err := getDeployEnvironmentIds(compartment)
	if err != nil {
		return err
	}
	for _, deployEnvironmentId := range deployEnvironmentIds {
		if ok := SweeperDefaultResourceId[deployEnvironmentId]; !ok {
			deleteDeployEnvironmentRequest := oci_devops.DeleteDeployEnvironmentRequest{}

			deleteDeployEnvironmentRequest.DeployEnvironmentId = &deployEnvironmentId

			deleteDeployEnvironmentRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "devops")
			_, error := deployEnvironmentClient.DeleteDeployEnvironment(context.Background(), deleteDeployEnvironmentRequest)
			if error != nil {
				fmt.Printf("Error deleting DeployEnvironment %s %s, It is possible that the resource is already deleted. Please verify manually \n", deployEnvironmentId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &deployEnvironmentId, deployEnvironmentSweepWaitCondition, time.Duration(3*time.Minute),
				deployEnvironmentSweepResponseFetchOperation, "devops", true)
		}
	}
	return nil
}

func getDeployEnvironmentIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "DeployEnvironmentId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	deployEnvironmentClient := GetTestClients(&schema.ResourceData{}).devopsClient()

	listDeployEnvironmentsRequest := oci_devops.ListDeployEnvironmentsRequest{}
	listDeployEnvironmentsRequest.CompartmentId = &compartmentId
	listDeployEnvironmentsRequest.LifecycleState = oci_devops.DeployEnvironmentLifecycleStateActive
	listDeployEnvironmentsResponse, err := deployEnvironmentClient.ListDeployEnvironments(context.Background(), listDeployEnvironmentsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DeployEnvironment list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, deployEnvironment := range listDeployEnvironmentsResponse.Items {
		id := *deployEnvironment.GetId()
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "DeployEnvironmentId", id)
	}
	return resourceIds, nil
}

func deployEnvironmentSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if deployEnvironmentResponse, ok := response.Response.(oci_devops.GetDeployEnvironmentResponse); ok {
		return deployEnvironmentResponse.GetLifecycleState() != oci_devops.DeployEnvironmentLifecycleStateDeleted
	}
	return false
}

func deployEnvironmentSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.devopsClient().GetDeployEnvironment(context.Background(), oci_devops.GetDeployEnvironmentRequest{
		DeployEnvironmentId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
