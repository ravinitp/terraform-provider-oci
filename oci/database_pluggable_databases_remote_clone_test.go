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
	pluggableDatabasesRemoteCloneRepresentation = map[string]interface{}{
		"cloned_pdb_name":                    Representation{RepType: Required, Create: `NewSalesPdb`},
		"pdb_admin_password":                 Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"pluggable_database_id":              Representation{RepType: Required, Create: `${oci_database_pluggable_database.test_pluggable_database.id}`},
		"source_container_db_admin_password": Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"target_container_database_id":       Representation{RepType: Required, Create: `${data.oci_database_database.tClone.id}`},
		"target_tde_wallet_password":         Representation{RepType: Required, Create: `BEstrO0ng_#11`},
		"should_pdb_admin_account_be_locked": Representation{RepType: Optional, Create: `false`},
	}
	AvailabilityDomainConfigClone = GenerateDataSourceFromRepresentationMap("oci_identity_availability_domains", "test_availability_domains_clone", Required, Create, availabilityDomainDataSourceRepresentation)

	ResourcePluggableDatabaseBaseCloneConfig = `

	data "oci_identity_availability_domains" "ADsClone" {
		compartment_id = "${var.compartment_id}"
	}

	resource "oci_core_virtual_network" "tClone" {
		compartment_id = "${var.compartment_id}"
		cidr_block = "10.1.0.0/16"
		display_name = "-tf-vcn-clone"
		dns_label = "tfvcnclone"
	}

	resource "oci_core_route_table" "tClone" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.tClone.id}"
		route_rules {
			cidr_block = "0.0.0.0/0"
			network_entity_id = "${oci_core_internet_gateway.tClone.id}"
		}
	}
	resource "oci_core_internet_gateway" "tClone" {
		compartment_id = "${var.compartment_id}"
		vcn_id = "${oci_core_virtual_network.tClone.id}"
		display_name = "-tf-internet-gateway-clone"
	}

	resource "oci_core_subnet" "tClone" {
		availability_domain = "${data.oci_identity_availability_domains.ADsClone.availability_domains.0.name}"
		cidr_block          = "10.1.20.0/24"
		display_name        = "TFSubnetClone1"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.tClone.id}"
		route_table_id      = "${oci_core_route_table.tClone.id}"
		dhcp_options_id     = "${oci_core_virtual_network.tClone.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.tClone.default_security_list_id}"]
		dns_label           = "tfsubnetclone"
	}
	resource "oci_core_subnet" "t2Clone" {
		availability_domain = "${data.oci_identity_availability_domains.ADsClone.availability_domains.0.name}"
		cidr_block          = "10.1.21.0/24"
		display_name        = "TFSubnetClone2"
		compartment_id      = "${var.compartment_id}"
		vcn_id              = "${oci_core_virtual_network.tClone.id}"
		route_table_id      = "${oci_core_route_table.tClone.id}"
		dhcp_options_id     = "${oci_core_virtual_network.tClone.default_dhcp_options_id}"
		security_list_ids   = ["${oci_core_virtual_network.tClone.default_security_list_id}"]
		dns_label           = "tfsubnetclone2"
	}
	resource "oci_core_network_security_group" "test_network_security_group_clone" {
         compartment_id  = "${var.compartment_id}"
		 vcn_id            = "${oci_core_virtual_network.tClone.id}"
         display_name      =  "displayName"
    }

	resource "oci_core_network_security_group" "test_network_security_group_clone2" {
		compartment_id = "${var.compartment_id}"
		vcn_id            = "${oci_core_virtual_network.tClone.id}"
	}`

	dbSystemForPluggableDbCloneRepresentation = `
		resource "oci_database_db_system" "tClone" {
			compartment_id = "${var.compartment_id}"
			subnet_id = "${oci_core_subnet.tClone.id}"
			database_edition = "ENTERPRISE_EDITION"
			availability_domain = "${data.oci_identity_availability_domains.ADsClone.availability_domains.0.name}"
			disk_redundancy = "NORMAL"
			shape = "VM.Standard2.4"
			ssh_public_keys = ["ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"]
			display_name = "-tf-dbSystem-clone-001"
			domain = "${oci_core_subnet.tClone.dns_label}.${oci_core_virtual_network.tClone.dns_label}.oraclevcn.com"
			hostname = "myOracleDB" // this will be lowercased server side
			data_storage_size_in_gb = "256"
			license_model = "LICENSE_INCLUDED"
			node_count = "1"
			fault_domains = ["FAULT-DOMAIN-1"]
			db_home {
				db_version = "19.11.0.0"
				display_name = "-tf-db-home-clone"
				database {
					admin_password = "BEstrO0ng_#11"
					db_name = "aTFdbC"
					character_set = "AL32UTF8"
					defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
					freeform_tags = {"Department" = "Finance"}
					ncharacter_set = "AL16UTF16"
					db_workload = "OLTP"
					pdb_name = "pdbName"
				}
			}
			db_system_options {
				storage_management = "LVM"
			}
			defined_tags = "${map("example-tag-namespace-all.example-tag", "originalValue")}"
			freeform_tags = {"Department" = "Finance"}
			nsg_ids = ["${oci_core_network_security_group.test_network_security_group_clone.id}"]
			lifecycle {
				ignore_changes = [
					db_home.0.db_version,
					defined_tags,
					db_home.0.database.0.defined_tags,
				]
			}
		}
		data "oci_database_db_systems" "tClone" {
			compartment_id = "${var.compartment_id}"
			filter {
				name   = "id"
				values = ["${oci_database_db_system.tClone.id}"]
			}
		}
		data "oci_database_db_homes" "tClone" {
			compartment_id = "${var.compartment_id}"
			db_system_id = "${oci_database_db_system.tClone.id}"
			filter {
				name   = "db_system_id"
				values = ["${oci_database_db_system.tClone.id}"]
			}
		}
		data "oci_database_db_home" "tClone" {
			db_home_id = "${data.oci_database_db_homes.tClone.db_homes.0.db_home_id}"
		}
		data "oci_database_databases" "tClone" {
			compartment_id = "${var.compartment_id}"
			db_home_id = "${data.oci_database_db_homes.tClone.db_homes.0.id}"
			filter {
				name   = "db_name"
				values = ["${oci_database_db_system.tClone.db_home.0.database.0.db_name}"]
			}
		}
		data "oci_database_database" "tClone" {
			  database_id = "${data.oci_database_databases.tClone.databases.0.id}"
		}`

	PluggableDatabaseResourceCloneDependencies = ResourcePluggableDatabaseBaseCloneConfig + dbSystemForPluggableDbCloneRepresentation
)

// issue-routing-tag: database/default
func TestDatabasePluggableDatabasesRemoteCloneResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabasePluggableDatabasesRemoteCloneResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_pluggable_databases_remote_clone.test_pluggable_databases_remote_clone"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+PluggableDatabaseResourceDependencies+PluggableDatabaseResourceCloneDependencies+
		GenerateResourceFromRepresentationMap("oci_database_pluggable_databases_remote_clone", "test_pluggable_databases_remote_clone", Optional, Create, pluggableDatabasesRemoteCloneRepresentation), "database", "pluggableDatabasesRemoteClone", t)

	ResourceTest(t, nil, []resource.TestStep{

		//Remote Clone
		{
			Config: config + compartmentIdVariableStr + PluggableDatabaseResourceDependencies + AvailabilityDomainConfigClone + PluggableDatabaseResourceCloneDependencies +
				GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", Optional, Update, pluggableDatabaseRepresentation) +
				GenerateResourceFromRepresentationMap("oci_database_pluggable_databases_remote_clone", "test_pluggable_databases_remote_clone", Optional, Create, pluggableDatabasesRemoteCloneRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cloned_pdb_name", "NewSalesPdb"),
				resource.TestCheckResourceAttrSet(resourceName, "pluggable_database_id"),
				resource.TestCheckResourceAttr(resourceName, "source_container_db_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttrSet(resourceName, "target_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "target_tde_wallet_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "should_pdb_admin_account_be_locked", "false"),
			),
		},
	})
}
