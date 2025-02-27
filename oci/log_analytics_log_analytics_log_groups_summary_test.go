// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	logAnalyticsLogGroupsSummarySingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"namespace":      Representation{RepType: Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	LogAnalyticsLogGroupsSummaryResourceDependencies = GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)

	LogAnalyticsLogGroupsSummaryResourceConfig = ""
)

// issue-routing-tag: log_analytics/default
func TestLogAnalyticsLogAnalyticsLogGroupsSummaryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLogAnalyticsLogAnalyticsLogGroupsSummaryResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_log_analytics_log_analytics_log_groups_summary.test_log_analytics_log_groups_summary"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				compartmentIdVariableStr +
				LogAnalyticsLogGroupsSummaryResourceDependencies +
				LogAnalyticsLogGroupsSummaryResourceConfig +
				GenerateDataSourceFromRepresentationMap("oci_log_analytics_log_analytics_log_groups_summary", "test_log_analytics_log_groups_summary", Required, Create, logAnalyticsLogGroupsSummarySingularDataSourceRepresentation),

			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "log_group_count"),
			),
		},
	})
}
