// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DeployComputeInstanceGroupStageRequiredOnlyResource = DeployComputeInstanceGroupStageResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployComputeInstanceGroupStageRepresentation)

	DeployComputeInstanceGroupStageResourceConfig = DeployComputeInstanceGroupStageResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployComputeInstanceGroupStageRepresentation)

	deployComputeInstanceGroupStageSingularDataSourceRepresentation = map[string]interface{}{
		"deploy_stage_id": Representation{RepType: Required, Create: `${oci_devops_deploy_stage.test_deploy_stage.id}`},
	}

	deployComputeInstanceGroupStageRepresentation = GetUpdatedRepresentationCopy("deploy_stage_type", Representation{RepType: Required, Create: `COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT`},
		RepresentationCopyWithNewProperties(RepresentationCopyWithRemovedProperties(deployStageRepresentation, []string{"wait_criteria"}), map[string]interface{}{
			"compute_instance_group_deploy_environment_id": Representation{RepType: Required, Create: `${oci_devops_deploy_environment.test_deploy_instance_group_environment.id}`},
			"deployment_spec_deploy_artifact_id":           Representation{RepType: Required, Create: `${oci_devops_deploy_artifact.test_deploy_generic_artifact.id}`},
			"rollout_policy":                               RepresentationGroup{Required, deployComputeInstanceStageRolloutPolicyRepresentation},
			"rollback_policy":                              RepresentationGroup{Optional, deployStageRollbackPolicyRepresentation},
			"load_balancer_config":                         RepresentationGroup{Optional, deployStageLoadBalancerInstanceGroupConfigRepresentation},
			"failure_policy":                               RepresentationGroup{Optional, deployComputeInstanceStageFailurePolicyRepresentation},
			"lifecycle":                                    RepresentationGroup{Required, ignoreDefinedTagsDifferencesRepresentation},
		}))

	deployComputeInstanceStageRolloutPolicyRepresentation = map[string]interface{}{
		"policy_type":            Representation{RepType: Required, Create: `COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT`},
		"batch_delay_in_seconds": Representation{RepType: Optional, Create: `5`},
		"batch_count":            Representation{RepType: Required, Create: `5`},
	}

	deployComputeInstanceStageFailurePolicyRepresentation = map[string]interface{}{
		"policy_type":   Representation{RepType: Required, Create: `COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT`},
		"failure_count": Representation{RepType: Required, Create: `1`},
	}

	deployStageLoadBalancerInstanceGroupConfigRepresentation = map[string]interface{}{
		"backend_port":     Representation{RepType: Optional, Create: `8080`},
		"listener_name":    Representation{RepType: Required, Create: `LoadBalancerListener`, Update: `LoadBalancerListener2`},
		"load_balancer_id": Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer_1.id}`, Update: `${oci_load_balancer_load_balancer.test_load_balancer_2.id}`},
	}

	DeployComputeInstanceGroupStageResourceDependencies = GenerateResourceFromRepresentationMap("oci_devops_deploy_artifact", "test_deploy_generic_artifact", Required, Create, deployGenericArtifactRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_environment", "test_deploy_instance_group_environment", Required, Create, deployInstanceGroupEnvironmentRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_deploy_pipeline", "test_deploy_pipeline", Required, Create, deployPipelineRepresentation) +
		GenerateResourceFromRepresentationMap("oci_devops_project", "test_project", Required, Create, devopsProjectRepresentation) +
		GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer_1", Required, Create, loadBalancerRepresentation) +
		GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer_2", Optional, Create, loadBalancerRepresentation) +
		LoadBalancerResourceDependencies
)

// issue-routing-tag: devops/default
func TestDevopsDeployStageResource_computeInstanceGroup(t *testing.T) {
	httpreplay.SetScenario("TestDevopsDeployStageResource_computeInstanceGroup")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_devops_deploy_stage.test_deploy_stage"
	datasourceName := "data.oci_devops_deploy_stages.test_deploy_stages"
	singularDatasourceName := "data.oci_devops_deploy_stage.test_deploy_stage"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+DeployComputeInstanceGroupStageResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Create, deployComputeInstanceGroupStageRepresentation), "devops", "deployStage", t)

	ResourceTest(t, testAccCheckDevopsDeployStageDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DeployComputeInstanceGroupStageResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployComputeInstanceGroupStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.batch_count", "5"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_group_deploy_environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_spec_deploy_artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "0"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DeployComputeInstanceGroupStageResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DeployComputeInstanceGroupStageResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Create, deployComputeInstanceGroupStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_type", "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.batch_count", "5"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.batch_delay_in_seconds", "5"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_group_deploy_environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_spec_deploy_artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.0.policy_type", "AUTOMATED_STAGE_ROLLBACK_POLICY"),
				resource.TestCheckResourceAttr(resourceName, "load_balancer_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "load_balancer_config.0.listener_name", "LoadBalancerListener"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_config.0.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "failure_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "failure_policy.0.failure_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "failure_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT"),

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
			Config: config + compartmentIdVariableStr + DeployComputeInstanceGroupStageResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployComputeInstanceGroupStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "deploy_stage_predecessor_collection.0.items.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.batch_count", "5"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.batch_delay_in_seconds", "5"),
				resource.TestCheckResourceAttr(resourceName, "rollout_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT"),
				resource.TestCheckResourceAttrSet(resourceName, "compute_instance_group_deploy_environment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "deployment_spec_deploy_artifact_id"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rollback_policy.0.policy_type", "AUTOMATED_STAGE_ROLLBACK_POLICY"),
				resource.TestCheckResourceAttr(resourceName, "load_balancer_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "load_balancer_config.0.listener_name", "LoadBalancerListener2"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_config.0.load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "failure_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "failure_policy.0.failure_count", "1"),
				resource.TestCheckResourceAttr(resourceName, "failure_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT"),

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
				GenerateDataSourceFromRepresentationMap("oci_devops_deploy_stages", "test_deploy_stages", Optional, Update, deployStageDataSourceRepresentation) +
				compartmentIdVariableStr + DeployComputeInstanceGroupStageResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Optional, Update, deployComputeInstanceGroupStageRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "deploy_pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),

				resource.TestCheckResourceAttr(datasourceName, "deploy_stage_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_devops_deploy_stage", "test_deploy_stage", Required, Create, deployComputeInstanceGroupStageSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DeployComputeInstanceGroupStageResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_predecessor_collection.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deploy_stage_predecessor_collection.0.items.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "deploy_stage_type", "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "project_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollout_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollout_policy.0.batch_count", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollout_policy.0.batch_delay_in_seconds", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollout_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_instance_group_deploy_environment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "deployment_spec_deploy_artifact_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollback_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rollback_policy.0.policy_type", "AUTOMATED_STAGE_ROLLBACK_POLICY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancer_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "load_balancer_config.0.listener_name", "LoadBalancerListener2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_config.0.load_balancer_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "failure_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "failure_policy.0.failure_count", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "failure_policy.0.policy_type", "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DeployComputeInstanceGroupStageResourceConfig,
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
