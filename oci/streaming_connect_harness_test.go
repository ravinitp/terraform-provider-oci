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
	oci_streaming "github.com/oracle/oci-go-sdk/v55/streaming"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ConnectHarnessRequiredOnlyResource = ConnectHarnessResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_streaming_connect_harness", "test_connect_harness", Required, Create, connectHarnessRepresentation)

	ConnectHarnessResourceConfig = ConnectHarnessResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_streaming_connect_harness", "test_connect_harness", Optional, Update, connectHarnessRepresentation)

	connectHarnessSingularDataSourceRepresentation = map[string]interface{}{
		"connect_harness_id": Representation{RepType: Required, Create: `${oci_streaming_connect_harness.test_connect_harness.id}`},
	}

	connectHarnessDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"id":             Representation{RepType: Optional, Create: `${oci_streaming_connect_harness.test_connect_harness.id}`},
		"name":           Representation{RepType: Optional, Create: `mynewconnectharness`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, connectHarnessDataSourceFilterRepresentation}}
	connectHarnessDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_streaming_connect_harness.test_connect_harness.id}`}},
	}

	connectHarnessRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"name":           Representation{RepType: Required, Create: `mynewconnectharness`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	ConnectHarnessResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: streaming/default
func TestStreamingConnectHarnessResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStreamingConnectHarnessResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_streaming_connect_harness.test_connect_harness"
	datasourceName := "data.oci_streaming_connect_harnesses.test_connect_harnesses"
	singularDatasourceName := "data.oci_streaming_connect_harness.test_connect_harness"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ConnectHarnessResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_streaming_connect_harness", "test_connect_harness", Optional, Create, connectHarnessRepresentation), "streaming", "connectHarness", t)

	ResourceTest(t, testAccCheckStreamingConnectHarnessDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ConnectHarnessResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_streaming_connect_harness", "test_connect_harness", Required, Create, connectHarnessRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "mynewconnectharness"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ConnectHarnessResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ConnectHarnessResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_streaming_connect_harness", "test_connect_harness", Optional, Create, connectHarnessRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "mynewconnectharness"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ConnectHarnessResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_streaming_connect_harness", "test_connect_harness", Optional, Create,
					RepresentationCopyWithNewProperties(connectHarnessRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "mynewconnectharness"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + ConnectHarnessResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_streaming_connect_harness", "test_connect_harness", Optional, Update, connectHarnessRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "mynewconnectharness"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				GenerateDataSourceFromRepresentationMap("oci_streaming_connect_harnesses", "test_connect_harnesses", Optional, Update, connectHarnessDataSourceRepresentation) +
				compartmentIdVariableStr + ConnectHarnessResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_streaming_connect_harness", "test_connect_harness", Optional, Update, connectHarnessRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", "mynewconnectharness"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "connect_harness.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "connect_harness.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "connect_harness.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "connect_harness.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "connect_harness.0.name", "mynewconnectharness"),
				resource.TestCheckResourceAttrSet(datasourceName, "connect_harness.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "connect_harness.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_streaming_connect_harness", "test_connect_harness", Required, Create, connectHarnessSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ConnectHarnessResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connect_harness_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "mynewconnectharness"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ConnectHarnessResourceConfig,
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

func testAccCheckStreamingConnectHarnessDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).streamAdminClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_streaming_connect_harness" {
			noResourceFound = false
			request := oci_streaming.GetConnectHarnessRequest{}

			tmp := rs.Primary.ID
			request.ConnectHarnessId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "streaming")

			response, err := client.GetConnectHarness(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_streaming.ConnectHarnessLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
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
	if !InSweeperExcludeList("StreamingConnectHarness") {
		resource.AddTestSweepers("StreamingConnectHarness", &resource.Sweeper{
			Name:         "StreamingConnectHarness",
			Dependencies: DependencyGraph["connectHarness"],
			F:            sweepStreamingConnectHarnessResource,
		})
	}
}

func sweepStreamingConnectHarnessResource(compartment string) error {
	streamAdminClient := GetTestClients(&schema.ResourceData{}).streamAdminClient()
	connectHarnessIds, err := getConnectHarnessIds(compartment)
	if err != nil {
		return err
	}
	for _, connectHarnessId := range connectHarnessIds {
		if ok := SweeperDefaultResourceId[connectHarnessId]; !ok {
			deleteConnectHarnessRequest := oci_streaming.DeleteConnectHarnessRequest{}

			deleteConnectHarnessRequest.ConnectHarnessId = &connectHarnessId

			deleteConnectHarnessRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "streaming")
			_, error := streamAdminClient.DeleteConnectHarness(context.Background(), deleteConnectHarnessRequest)
			if error != nil {
				fmt.Printf("Error deleting ConnectHarness %s %s, It is possible that the resource is already deleted. Please verify manually \n", connectHarnessId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &connectHarnessId, connectHarnessSweepWaitCondition, time.Duration(3*time.Minute),
				connectHarnessSweepResponseFetchOperation, "streaming", true)
		}
	}
	return nil
}

func getConnectHarnessIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "ConnectHarnessId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	streamAdminClient := GetTestClients(&schema.ResourceData{}).streamAdminClient()

	listConnectHarnessesRequest := oci_streaming.ListConnectHarnessesRequest{}
	listConnectHarnessesRequest.CompartmentId = &compartmentId
	listConnectHarnessesRequest.LifecycleState = oci_streaming.ConnectHarnessSummaryLifecycleStateActive
	listConnectHarnessesResponse, err := streamAdminClient.ListConnectHarnesses(context.Background(), listConnectHarnessesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ConnectHarness list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, connectHarness := range listConnectHarnessesResponse.Items {
		id := *connectHarness.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "ConnectHarnessId", id)
	}
	return resourceIds, nil
}

func connectHarnessSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if connectHarnessResponse, ok := response.Response.(oci_streaming.GetConnectHarnessResponse); ok {
		return connectHarnessResponse.LifecycleState != oci_streaming.ConnectHarnessLifecycleStateDeleted
	}
	return false
}

func connectHarnessSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.streamAdminClient().GetConnectHarness(context.Background(), oci_streaming.GetConnectHarnessRequest{
		ConnectHarnessId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
