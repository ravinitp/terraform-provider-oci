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
	oci_mysql "github.com/oracle/oci-go-sdk/v55/mysql"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ChannelRequiredOnlyResource = ChannelResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_mysql_channel", "test_channel", Required, Create, channelRepresentation)

	channelSingularDataSourceRepresentation = map[string]interface{}{
		"channel_id": Representation{RepType: Required, Create: `${oci_mysql_channel.test_channel.id}`},
	}

	channelDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"channel_id":     Representation{RepType: Optional, Create: `${oci_mysql_channel.test_channel.id}`},
		"db_system_id":   Representation{RepType: Optional, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"is_enabled":     Representation{RepType: Optional, Create: `true`, Update: `false`},
		"filter":         RepresentationGroup{Required, channelDataSourceFilterRepresentation}}
	channelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_mysql_channel.test_channel.id}`}},
	}

	channelRepresentation = map[string]interface{}{
		"source":         RepresentationGroup{Required, channelSourceRepresentation},
		"target":         RepresentationGroup{Required, channelTargetRepresentation},
		"compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"defined_tags":   Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{RepType: Optional, Create: `description`, Update: `description2`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":     Representation{RepType: Optional, Create: `true`, Update: `false`},
	}

	sslCaCertificateRepresentation = map[string]interface{}{
		"certificate_type": Representation{RepType: Optional, Update: "PEM"},
		"contents":         Representation{RepType: Optional, Update: "${var.ca_certificate_value}"},
	}

	channelSourceRepresentation = map[string]interface{}{
		"hostname":    Representation{RepType: Required, Create: `hostname.my.company.com`, Update: `hostname2.my.company.com`},
		"password":    Representation{RepType: Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"source_type": Representation{RepType: Required, Create: `MYSQL`},
		"username":    Representation{RepType: Required, Create: `username`, Update: `username2`},
		"ssl_mode":    Representation{RepType: Required, Create: `REQUIRED`, Update: `VERIFY_CA`},
		"port":        Representation{RepType: Optional, Create: `3300`, Update: `3306`},
	}
	channelSourceWithCertificateRepresentation = map[string]interface{}{
		"hostname":           Representation{RepType: Optional, Update: `hostname2.my.company.com`},
		"password":           Representation{RepType: Optional, Update: `BEstrO0ng_#12`},
		"source_type":        Representation{RepType: Optional, Update: `MYSQL`},
		"username":           Representation{RepType: Optional, Update: `username2`},
		"ssl_mode":           Representation{RepType: Optional, Update: `VERIFY_CA`},
		"ssl_ca_certificate": RepresentationGroup{Optional, sslCaCertificateRepresentation},
		"port":               Representation{RepType: Optional, Update: `3306`},
	}

	channelTargetRepresentation = map[string]interface{}{
		"db_system_id":     Representation{RepType: Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"target_type":      Representation{RepType: Required, Create: `DBSYSTEM`},
		"applier_username": Representation{RepType: Optional, Create: `adminUser`},
		"channel_name":     Representation{RepType: Optional, Create: `channelname`, Update: `channelname2`},
	}

	ChannelWithOptionalsResource = ChannelResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_mysql_channel", "test_channel", Optional, Create, channelRepresentation)

	ChannelUpdateResource = ChannelResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_mysql_channel", "test_channel", Optional, Update,
			GetUpdatedRepresentationCopy("source", RepresentationGroup{Optional, channelSourceWithCertificateRepresentation}, channelRepresentation))

	ChannelResourceConfig = ChannelUpdateResource

	ChannelResourceDependencies = MysqlDbSystemResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", Required, Create, mysqlDbSystemRepresentation) + caCertificateVariableStr
)

// issue-routing-tag: mysql/default
func TestMysqlChannelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlChannelResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_channel.test_channel"
	datasourceName := "data.oci_mysql_channels.test_channels"
	singularDatasourceName := "data.oci_mysql_channel.test_channel"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ChannelResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_mysql_channel", "test_channel", Optional, Create, channelRepresentation), "mysql", "channel", t)

	ResourceTest(t, testAccCheckMysqlChannelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ChannelRequiredOnlyResource,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.hostname", "hostname.my.company.com"),
				resource.TestCheckResourceAttr(resourceName, "source.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "MYSQL"),
				resource.TestCheckResourceAttr(resourceName, "source.0.ssl_mode", "REQUIRED"),
				resource.TestCheckResourceAttr(resourceName, "source.0.username", "username"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "target.0.target_type", "DBSYSTEM"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ChannelResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ChannelWithOptionalsResource,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.hostname", "hostname.my.company.com"),
				resource.TestCheckResourceAttr(resourceName, "source.0.password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "source.0.port", "3300"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "MYSQL"),
				resource.TestCheckResourceAttr(resourceName, "source.0.ssl_mode", "REQUIRED"),
				resource.TestCheckResourceAttr(resourceName, "source.0.ssl_ca_certificate.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "source.0.username", "username"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.applier_username", "adminUser"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.channel_name"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "target.0.target_type", "DBSYSTEM"),

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
			Config: config + compartmentIdVariableStr + ChannelUpdateResource,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "source.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.hostname", "hostname2.my.company.com"),
				resource.TestCheckResourceAttr(resourceName, "source.0.password", "BEstrO0ng_#12"),
				resource.TestCheckResourceAttr(resourceName, "source.0.port", "3306"),
				resource.TestCheckResourceAttr(resourceName, "source.0.source_type", "MYSQL"),
				resource.TestCheckResourceAttr(resourceName, "source.0.ssl_mode", "VERIFY_CA"),
				resource.TestCheckResourceAttr(resourceName, "source.0.ssl_ca_certificate.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "source.0.username", "username2"),
				resource.TestCheckResourceAttr(resourceName, "target.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "target.0.applier_username", "adminUser"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.channel_name"),
				resource.TestCheckResourceAttrSet(resourceName, "target.0.db_system_id"),
				resource.TestCheckResourceAttr(resourceName, "target.0.target_type", "DBSYSTEM"),

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
				GenerateDataSourceFromRepresentationMap("oci_mysql_channels", "test_channels", Optional, Update, channelDataSourceRepresentation) +
				compartmentIdVariableStr + ChannelUpdateResource,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "channel_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "db_system_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "is_enabled", "false"),

				resource.TestCheckResourceAttr(datasourceName, "channels.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_mysql_channel", "test_channel", Required, Create, channelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ChannelUpdateResource,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "channel_id"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ChannelUpdateResource,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"lifecycle_details",
				"source.0.password",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckMysqlChannelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).channelsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_mysql_channel" {
			noResourceFound = false
			request := oci_mysql.GetChannelRequest{}

			tmp := rs.Primary.ID
			request.ChannelId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "mysql")

			response, err := client.GetChannel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_mysql.ChannelLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("MysqlChannel") {
		resource.AddTestSweepers("MysqlChannel", &resource.Sweeper{
			Name:         "MysqlChannel",
			Dependencies: DependencyGraph["channel"],
			F:            sweepMysqlChannelResource,
		})
	}
}

func sweepMysqlChannelResource(compartment string) error {
	channelsClient := GetTestClients(&schema.ResourceData{}).channelsClient()
	channelIds, err := getChannelIds(compartment)
	if err != nil {
		return err
	}
	for _, channelId := range channelIds {
		if ok := SweeperDefaultResourceId[channelId]; !ok {
			deleteChannelRequest := oci_mysql.DeleteChannelRequest{}

			deleteChannelRequest.ChannelId = &channelId

			deleteChannelRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "mysql")
			_, error := channelsClient.DeleteChannel(context.Background(), deleteChannelRequest)
			if error != nil {
				fmt.Printf("Error deleting Channel %s %s, It is possible that the resource is already deleted. Please verify manually \n", channelId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &channelId, channelSweepWaitCondition, time.Duration(3*time.Minute),
				channelSweepResponseFetchOperation, "mysql", true)
		}
	}
	return nil
}

func getChannelIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "ChannelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	channelsClient := GetTestClients(&schema.ResourceData{}).channelsClient()

	listChannelsRequest := oci_mysql.ListChannelsRequest{}
	listChannelsRequest.CompartmentId = &compartmentId
	listChannelsRequest.LifecycleState = oci_mysql.ChannelLifecycleStateActive
	listChannelsResponse, err := channelsClient.ListChannels(context.Background(), listChannelsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Channel list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, channel := range listChannelsResponse.Items {
		id := *channel.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "ChannelId", id)
	}
	return resourceIds, nil
}

func channelSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if channelResponse, ok := response.Response.(oci_mysql.GetChannelResponse); ok {
		return channelResponse.LifecycleState != oci_mysql.ChannelLifecycleStateDeleted
	}
	return false
}

func channelSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.channelsClient().GetChannel(context.Background(), oci_mysql.GetChannelRequest{
		ChannelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
