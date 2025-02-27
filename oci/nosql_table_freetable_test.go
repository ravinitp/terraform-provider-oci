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
	FreeTableResourceConfig = GenerateResourceFromRepresentationMap("oci_nosql_table", "test_freetable", Required, Create, freeTableRepresentation)

	freeTableDdlStatement = "CREATE TABLE IF NOT EXISTS test_freetable(id INTEGER, name STRING, age STRING, PRIMARY KEY(SHARD(id)))"

	freeTableRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"ddl_statement":       Representation{RepType: Required, Create: freeTableDdlStatement},
		"name":                Representation{RepType: Required, Create: "test_freetable"},
		"table_limits":        RepresentationGroup{Required, freeTableTableLimitsRepresentation},
		"is_auto_reclaimable": Representation{RepType: Required, Create: `true`},
	}
	freeTableTableLimitsRepresentation = map[string]interface{}{
		"max_read_units":     Representation{RepType: Required, Create: `50`},
		"max_write_units":    Representation{RepType: Required, Create: `50`},
		"max_storage_in_gbs": Representation{RepType: Required, Create: `1`},
	}

	freeTableDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"name":           Representation{RepType: Optional, Create: `test_freetable`},
		"state":          Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, freeTableDataSourceFilterRepresentation}}
	freeTableDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_nosql_table.test_freetable.id}`}},
	}

	freeTableSingularDataSourceRepresentation = map[string]interface{}{
		"table_name_or_id": Representation{RepType: Required, Create: `${oci_nosql_table.test_freetable.id}`},
		"compartment_id":   Representation{RepType: Required, Create: `${var.compartment_id}`},
	}
)

// issue-routing-tag: nosql/default
func TestNosqlTableResource_freeTable(t *testing.T) {
	httpreplay.SetScenario("TestNosqlTableResource_freeTable")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_nosql_table.test_freetable"
	dataResourceName := "data.oci_nosql_tables.test_freetables"
	singularDatasourceName := "data.oci_nosql_table.test_freetable"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckNosqlTableDestroy,
		Steps: []resource.TestStep{

			// verify Create table: free table with name of "test_freetable"
			{
				Config: config + compartmentIdVariableStr +
					GenerateResourceFromRepresentationMap("oci_nosql_table", "test_freetable", Required, Create, freeTableRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "ddl_statement", freeTableDdlStatement),
					resource.TestCheckResourceAttr(resourceName, "name", "test_freetable"),
					resource.TestCheckResourceAttr(resourceName, "is_auto_reclaimable", "true"),
					resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "table_limits.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_read_units", "50"),
					resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_write_units", "50"),
					resource.TestCheckResourceAttr(resourceName, "table_limits.0.max_storage_in_gbs", "1"),
				),
			},
			// verify datasource: test_freetable
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_nosql_tables", "test_freetables", Optional, Create, freeTableDataSourceRepresentation) +
					compartmentIdVariableStr + FreeTableResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(dataResourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(dataResourceName, "name", "test_freetable"),
					resource.TestCheckResourceAttr(dataResourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(dataResourceName, "table_collection.#", "1"),
					resource.TestCheckResourceAttrSet(dataResourceName, "table_collection.0.id"),
					resource.TestCheckResourceAttr(dataResourceName, "table_collection.0.name", "test_freetable"),
					resource.TestCheckResourceAttr(dataResourceName, "table_collection.0.state", "ACTIVE"),
					resource.TestCheckResourceAttr(dataResourceName, "table_collection.0.is_auto_reclaimable", "true"),
					resource.TestCheckResourceAttr(dataResourceName, "table_collection.0.system_tags.%", "1"),
				),
			},
			// verify singular datasource: test_freetable
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_nosql_table", "test_freetable", Required, Create, freeTableSingularDataSourceRepresentation) +
					compartmentIdVariableStr + FreeTableResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "table_name_or_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "test_freetable"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ddl_statement", freeTableDdlStatement),
					resource.TestCheckResourceAttr(singularDatasourceName, "schema.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_auto_reclaimable", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.0.max_read_units", "50"),
					resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.0.max_write_units", "50"),
					resource.TestCheckResourceAttr(singularDatasourceName, "table_limits.0.max_storage_in_gbs", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				),
			},
			// remove the free table resource
			{
				Config: config + compartmentIdVariableStr + FreeTableResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
