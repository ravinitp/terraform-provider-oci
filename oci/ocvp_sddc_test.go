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
	oci_ocvp "github.com/oracle/oci-go-sdk/v55/ocvp"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	SddcRequiredOnlyResource = SddcResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", Required, Create, sddcRepresentation)

	SddcResourceConfig = SddcResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", Optional, Update, sddcRepresentation)

	SddcV7ResourceConfig = SddcResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", Optional, Update, sddcV7Representation)

	sddcSingularDataSourceRepresentation = map[string]interface{}{
		"sddc_id": Representation{RepType: Required, Create: `${oci_ocvp_sddc.test_sddc.id}`},
	}

	sddcDataSourceRepresentation = map[string]interface{}{
		"compartment_id":              Representation{RepType: Required, Create: `${var.compartment_id}`},
		"compute_availability_domain": Representation{RepType: Optional, Create: `${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}`},
		"display_name":                Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"state":                       Representation{RepType: Optional, Create: `ACTIVE`},
		"filter":                      RepresentationGroup{Required, sddcDataSourceFilterRepresentation}}
	sddcDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_ocvp_sddc.test_sddc.id}`}},
	}

	sddcRepresentation = map[string]interface{}{
		"compartment_id":               Representation{RepType: Required, Create: `${var.compartment_id}`},
		"initial_sku":                  Representation{RepType: Optional, Create: `HOUR`},
		"compute_availability_domain":  Representation{RepType: Required, Create: `${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}`},
		"esxi_hosts_count":             Representation{RepType: Required, Create: `3`},
		"nsx_edge_uplink1vlan_id":      Representation{RepType: Required, Create: `${oci_core_vlan.test_nsx_edge_uplink1_vlan.id}`},
		"nsx_edge_uplink2vlan_id":      Representation{RepType: Required, Create: `${oci_core_vlan.test_nsx_edge_uplink2_vlan.id}`},
		"nsx_edge_vtep_vlan_id":        Representation{RepType: Required, Create: `${oci_core_vlan.test_nsx_edge_vtep_vlan.id}`},
		"nsx_vtep_vlan_id":             Representation{RepType: Required, Create: `${oci_core_vlan.test_nsx_vtep_vlan.id}`},
		"provisioning_subnet_id":       Representation{RepType: Required, Create: `${oci_core_subnet.test_provisioning_subnet.id}`},
		"ssh_authorized_keys":          Representation{RepType: Required, Create: `ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`},
		"vmotion_vlan_id":              Representation{RepType: Required, Create: `${oci_core_vlan.test_vmotion_net_vlan.id}`},
		"vmware_software_version":      Representation{RepType: Required, Create: `6.7 update 3`, Update: `6.5 update 3`},
		"vsan_vlan_id":                 Representation{RepType: Required, Create: `${oci_core_vlan.test_vsan_net_vlan.id}`},
		"vsphere_vlan_id":              Representation{RepType: Required, Create: `${oci_core_vlan.test_vsphere_net_vlan.id}`},
		"defined_tags":                 Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                 Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"hcx_action":                   Representation{RepType: Optional, Create: UpgradeHcxAction},
		"hcx_vlan_id":                  Representation{RepType: Optional, Create: `${oci_core_vlan.test_hcx_vlan.id}`},
		"instance_display_name_prefix": Representation{RepType: Optional, Create: `njki`},
		"is_hcx_enabled":               Representation{RepType: Optional, Create: `true`},
		"workload_network_cidr":        Representation{RepType: Optional, Create: `172.20.0.0/24`},
		"provisioning_vlan_id":         Representation{RepType: Optional, Create: `${oci_core_vlan.test_provisioning_vlan.id}`},
		"replication_vlan_id":          Representation{RepType: Optional, Create: `${oci_core_vlan.test_replication_vlan.id}`},
		"refresh_hcx_license_status":   Representation{RepType: Optional, Create: `false`, Update: `true`},
		"lifecycle":                    RepresentationGroup{Required, ignoreDefinedTagsChangesRepresentation},
	}

	ignoreDefinedTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": Representation{RepType: Required, Create: []string{`defined_tags`}},
	}

	sddcV7Representation = RepresentationCopyWithNewProperties(sddcRepresentation, map[string]interface{}{
		"vmware_software_version": Representation{RepType: Required, Create: `7.0 update 2`},
	})

	SddcResourceDependencies = DefinedTagsDependencies +
		GenerateDataSourceFromRepresentationMap("oci_ocvp_supported_vmware_software_versions", "test_supported_vmware_software_versions", Required, Create, supportedVmwareSoftwareVersionDataSourceRepresentation) + `

data "oci_core_services" "test_services" {}

data "oci_identity_availability_domains" "ADs" {
    compartment_id = "${var.compartment_id}"
}

resource "oci_core_vcn" "test_vcn_ocvp" {
    cidr_block = "10.0.0.0/16"
    compartment_id = "${var.compartment_id}"
    display_name = "VmWareOCVP"
    dns_label = "vmwareocvp"
}


resource oci_core_nat_gateway test_nat_gateway_ocvp {
  block_traffic  = "false"
  compartment_id = var.compartment_id

  display_name = "NAT Gateway OCVP"
  freeform_tags = {
    "VCN" = "VCN-2020-09-11T00:43:42"
  }
  vcn_id = oci_core_vcn.test_vcn_ocvp.id
}

resource oci_core_route_table test_route_table_for_vsphere_vlan {
  compartment_id = var.compartment_id

  display_name = "Route Table for VLAN-grk-vSphere"
  freeform_tags = {
    "VMware" = "VMware-2020-09-11T00:47:02"
  }
  route_rules {
    #description = <<Optional value not found in discovery>>
    destination       = "0.0.0.0/0"
    destination_type  = "CIDR_BLOCK"
    network_entity_id = oci_core_nat_gateway.test_nat_gateway_ocvp.id
  }
  vcn_id = oci_core_vcn.test_vcn_ocvp.id
}

resource "oci_core_network_security_group" "test_nsg_allow_all" {
    compartment_id = "${var.compartment_id}"
    display_name = "nsg-allow-all"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
}

resource oci_core_network_security_group_security_rule test_nsg_security_rule_1 {
  destination_type          = ""
  direction                 = "INGRESS"
  network_security_group_id = "${oci_core_network_security_group.test_nsg_allow_all.id}"
  protocol                  = "all"
  source                    = "0.0.0.0/0"
  source_type               = "CIDR_BLOCK"
}

resource oci_core_network_security_group_security_rule test_nsg_security_rule_2 {
  destination               = "0.0.0.0/0"
  destination_type          = "CIDR_BLOCK"
  direction                 = "EGRESS"
  network_security_group_id = "${oci_core_network_security_group.test_nsg_allow_all.id}"
  protocol                  = "all"
  source_type = ""
}

resource "oci_core_service_gateway" "export_sgw" {
    compartment_id = "${var.compartment_id}"
    display_name = "sgw"
    services {
        service_id = "${lookup(data.oci_core_services.test_services.services[0], "id")}"
    }
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
}

resource "oci_core_default_dhcp_options" "default_dhcp_options_ocvp"{
    display_name = "Default DHCP Options for OCVP"
    manage_default_resource_id = "${oci_core_vcn.test_vcn_ocvp.default_dhcp_options_id}"
    options {
        custom_dns_servers = []
        server_type = "VcnLocalPlusInternet"
        type = "DomainNameServer"
    }
    options {
            search_domain_names = ["vmwareocvp.oraclevcn.com"]
            type = "SearchDomain"
    }
}

resource "oci_core_route_table" "private_rt" {
    compartment_id = "${var.compartment_id}"
    display_name = "private-rt"
    route_rules {
        destination = "${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}"
        destination_type = "SERVICE_CIDR_BLOCK"
        network_entity_id = "${oci_core_service_gateway.export_sgw.id}"
    }
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
}

resource "oci_core_security_list" "private_sl" {
    compartment_id = "${var.compartment_id}"
    display_name = "private-sl"
    egress_security_rules {
        destination = "0.0.0.0/0"
        destination_type = "CIDR_BLOCK"
        protocol = "all"
        stateless = "false"
    }
    ingress_security_rules {
        description = "TCP traffic for ports: 22 SSH Remote Login Protocol"
        protocol = "6"
        source = "10.0.0.0/16"
        source_type = "CIDR_BLOCK"
        stateless = "false"
        tcp_options {
            max = "22"
            min = "22"
        }
    }
    ingress_security_rules {
        description = "ICMP traffic for: 3 Destination Unreachable"
        icmp_options {
            code = "3"
            type = "3"
        }
        protocol = "1"
        source = "10.0.0.0/16"
        source_type = "CIDR_BLOCK"
        stateless = "false"
    }
    ingress_security_rules {
        protocol = "all"
        source = "0.0.0.0/0"
        source_type = "CIDR_BLOCK"
        stateless = "false"
    }
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
}

resource "oci_core_default_security_list" "default_security_list_ocvp" {
    display_name = "Default Security List for OCVP"
    egress_security_rules {
        destination = "0.0.0.0/0"
        destination_type = "CIDR_BLOCK"
        protocol = "all"
        stateless = "false"
    }
    ingress_security_rules {
        protocol = "6"
        source = "0.0.0.0/0"
        source_type = "CIDR_BLOCK"
        stateless = "false"
        tcp_options {
            max = "22"
            min = "22"
        }
    }
    ingress_security_rules {
        icmp_options {
            code = "4"
            type = "3"
        }
        protocol = "1"
        source = "0.0.0.0/0"
        source_type = "CIDR_BLOCK"
        stateless = "false"
    }
    ingress_security_rules {
        icmp_options {
            code = "-1"
            type = "3"
        }
        protocol = "1"
        source = "10.0.0.0/16"
        source_type = "CIDR_BLOCK"
        stateless = "false"
    }
    manage_default_resource_id = "${oci_core_vcn.test_vcn_ocvp.default_security_list_id}"
}

resource "oci_core_subnet" "test_provisioning_subnet" {
    cidr_block = "10.0.103.128/25"
    compartment_id = "${var.compartment_id}"
    dhcp_options_id = "${oci_core_vcn.test_vcn_ocvp.default_dhcp_options_id}"
    display_name = "provisioning-subnet"
    dns_label = "provisioningsub"
    prohibit_public_ip_on_vnic = "true"
    route_table_id = "${oci_core_route_table.private_rt.id}"
    security_list_ids = ["${oci_core_security_list.private_sl.id}"]
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
}

resource "oci_core_vlan" "test_nsx_edge_uplink2_vlan" {
    display_name = "NSX-Edge-UP2"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.0.103.0/25"
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
    nsg_ids = ["${oci_core_network_security_group.test_nsg_allow_all.id}"]
    route_table_id = "${oci_core_vcn.test_vcn_ocvp.default_route_table_id}"
}

resource "oci_core_vlan" "test_nsx_edge_uplink1_vlan" {
    display_name = "NSX-Edge-UP1"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.0.100.0/25"
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
    nsg_ids = ["${oci_core_network_security_group.test_nsg_allow_all.id}"]
    route_table_id = "${oci_core_vcn.test_vcn_ocvp.default_route_table_id}"
}

resource "oci_core_vlan" "test_nsx_vtep_vlan" {
    display_name = "NSX-vTep"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.0.101.0/25"
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
    nsg_ids = ["${oci_core_network_security_group.test_nsg_allow_all.id}"]
    route_table_id = "${oci_core_vcn.test_vcn_ocvp.default_route_table_id}"
}

resource "oci_core_vlan" "test_nsx_edge_vtep_vlan" {
    display_name = "NSX Edge-vTep"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.0.102.0/25"
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
    nsg_ids = ["${oci_core_network_security_group.test_nsg_allow_all.id}"]
    route_table_id = "${oci_core_vcn.test_vcn_ocvp.default_route_table_id}"
}

resource "oci_core_vlan" "test_vsan_net_vlan" {
    display_name = "vSAN-Net"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.0.101.128/25"
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
    nsg_ids = ["${oci_core_network_security_group.test_nsg_allow_all.id}"]
    route_table_id = "${oci_core_vcn.test_vcn_ocvp.default_route_table_id}"
}

resource "oci_core_vlan" "test_vmotion_net_vlan" {
    display_name = "vMotion-Net"
    availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
    cidr_block = "10.0.102.128/25"
    compartment_id = "${var.compartment_id}"
    vcn_id = "${oci_core_vcn.test_vcn_ocvp.id}"
    nsg_ids = ["${oci_core_network_security_group.test_nsg_allow_all.id}"]
    route_table_id = "${oci_core_vcn.test_vcn_ocvp.default_route_table_id}"
}

resource "oci_core_vlan" "test_vsphere_net_vlan" {
  display_name        = "vSphere-Net"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block          = "10.0.100.128/26"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id = oci_core_route_table.test_route_table_for_vsphere_vlan.id
}

resource "oci_core_vlan" "test_hcx_vlan" {
  display_name        = "hcx"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block          = "10.0.100.192/26"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_provisioning_vlan" {
  display_name        = "provisioning-vlan"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block          = "10.0.104.128/25"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

resource "oci_core_vlan" "test_replication_vlan" {
  display_name        = "replication-vlan"
  availability_domain = "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[0],"name")}"
  cidr_block          = "10.0.104.0/25"
  compartment_id      = var.compartment_id
  vcn_id              = oci_core_vcn.test_vcn_ocvp.id
  nsg_ids             = [oci_core_network_security_group.test_nsg_allow_all.id]
  route_table_id      = oci_core_vcn.test_vcn_ocvp.default_route_table_id
}

`
)

// issue-routing-tag: ocvp/default
func TestOcvpSddcResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpSddcResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ocvp_sddc.test_sddc"
	datasourceName := "data.oci_ocvp_sddcs.test_sddcs"
	singularDatasourceName := "data.oci_ocvp_sddc.test_sddc"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+SddcResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", Optional, Create, sddcRepresentation), "ocvp", "sddc", t)

	ResourceTest(t, testAccCheckOcvpSddcDestroy, []resource.TestStep{
		//verify Create
		{
			Config: config + compartmentIdVariableStr + SddcResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", Required, Create, sddcRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "actual_esxi_hosts_count", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_uplink1vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_uplink2vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ssh_authorized_keys"),
				resource.TestCheckResourceAttrSet(resourceName, "vmotion_vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", "6.7 update 3"),
				resource.TestCheckResourceAttrSet(resourceName, "vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vsphere_vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_enterprise_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_pending_downgrade", "false"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// verify Update VMware version
		{
			Config: config + compartmentIdVariableStr + SddcResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", Required, Update, sddcRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttr(resourceName, "esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "actual_esxi_hosts_count", "3"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_uplink1vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_uplink2vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ssh_authorized_keys"),
				resource.TestCheckResourceAttrSet(resourceName, "vmotion_vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", "6.5 update 3"),
				resource.TestCheckResourceAttrSet(resourceName, "vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vsphere_vlan_id"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SddcResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SddcResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", Optional, Create, sddcV7Representation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "actual_esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hcx_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "initial_sku", "HOUR"),
				resource.TestCheckResourceAttr(resourceName, "instance_display_name_prefix", "njki"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_uplink1vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_uplink2vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_manager_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_manager_private_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provisioning_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ssh_authorized_keys"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcenter_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "vcenter_private_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vmotion_vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", "7.0 update 2"),
				resource.TestCheckResourceAttrSet(resourceName, "vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vsphere_vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "workload_network_cidr", "172.20.0.0/24"),
				resource.TestCheckResourceAttrSet(resourceName, "hcx_on_prem_licenses.#"),
				resource.TestCheckResourceAttr(resourceName, "hcx_action", UpgradeHcxAction),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_enterprise_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_pending_downgrade", "false"),
				resource.TestCheckResourceAttr(resourceName, "refresh_hcx_license_status", "false"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + SddcResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", Optional, Create,
					RepresentationCopyWithNewProperties(sddcV7Representation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "actual_esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hcx_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "initial_sku", "HOUR"),
				resource.TestCheckResourceAttr(resourceName, "instance_display_name_prefix", "njki"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_uplink1vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_uplink2vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_manager_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_manager_private_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provisioning_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ssh_authorized_keys"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcenter_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "vcenter_private_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vmotion_vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", "7.0 update 2"),
				resource.TestCheckResourceAttrSet(resourceName, "vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vsphere_vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "workload_network_cidr", "172.20.0.0/24"),
				resource.TestCheckResourceAttrSet(resourceName, "hcx_on_prem_licenses.#"),
				resource.TestCheckResourceAttr(resourceName, "hcx_action", UpgradeHcxAction),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_enterprise_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_pending_downgrade", "false"),
				resource.TestCheckResourceAttr(resourceName, "refresh_hcx_license_status", "false"),

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
		// Cannot Update VMware version here because some of the optional arguments are not applicable to VMware version less than 7.0
		{
			Config: config + compartmentIdVariableStr + SddcResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", Optional, Update, sddcV7Representation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "actual_esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "hcx_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "initial_sku", "HOUR"),
				resource.TestCheckResourceAttr(resourceName, "instance_display_name_prefix", "njki"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_uplink1vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_uplink2vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_edge_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_manager_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_manager_private_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "nsx_vtep_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provisioning_subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "provisioning_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "replication_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "ssh_authorized_keys"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "vcenter_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "vcenter_private_ip_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vmotion_vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "vmware_software_version", "7.0 update 2"),
				resource.TestCheckResourceAttrSet(resourceName, "vsan_vlan_id"),
				resource.TestCheckResourceAttrSet(resourceName, "vsphere_vlan_id"),
				resource.TestCheckResourceAttr(resourceName, "workload_network_cidr", "172.20.0.0/24"),
				resource.TestCheckResourceAttrSet(resourceName, "hcx_on_prem_licenses.#"),
				resource.TestCheckResourceAttrSet(resourceName, "time_hcx_billing_cycle_end"),
				resource.TestCheckResourceAttrSet(resourceName, "time_hcx_license_status_updated"),
				resource.TestCheckResourceAttr(resourceName, "hcx_action", UpgradeHcxAction),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_enterprise_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_hcx_pending_downgrade", "false"),
				resource.TestCheckResourceAttr(resourceName, "refresh_hcx_license_status", "true"),

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
				GenerateDataSourceFromRepresentationMap("oci_ocvp_sddcs", "test_sddcs", Optional, Update, sddcDataSourceRepresentation) +
				compartmentIdVariableStr + SddcV7ResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.compute_availability_domain"),
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.0.vmware_software_version", "7.0 update 2"),
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.0.esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.0.actual_esxi_hosts_count", "3"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "sddc_collection.0.time_updated"),
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.0.state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "sddc_collection.0.freeform_tags.%", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_ocvp_sddc", "test_sddc", Required, Create, sddcSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SddcV7ResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sddc_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compute_availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "actual_esxi_hosts_count", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hcx_fqdn"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hcx_initial_password"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hcx_on_prem_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hcx_on_prem_licenses.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hcx_private_ip_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "hcx_vlan_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_sku", "HOUR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "instance_display_name_prefix", "njki"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_hcx_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_hcx_enterprise_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_hcx_pending_downgrade"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_edge_uplink_ip_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_manager_fqdn"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_manager_initial_password"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_manager_private_ip_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_manager_username"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nsx_overlay_segment_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ssh_authorized_keys"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_hcx_billing_cycle_end"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_hcx_license_status_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcenter_fqdn"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcenter_initial_password"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcenter_private_ip_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vcenter_username"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vmware_software_version", "7.0 update 2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "workload_network_cidr", "172.20.0.0/24"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + SddcV7ResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"hcx_action", "refresh_hcx_license_status"},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOcvpSddcDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).sddcClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ocvp_sddc" {
			noResourceFound = false
			request := oci_ocvp.GetSddcRequest{}

			tmp := rs.Primary.ID
			request.SddcId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "ocvp")

			response, err := client.GetSddc(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ocvp.LifecycleStatesDeleted): true,
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
	if !InSweeperExcludeList("OcvpSddc") {
		resource.AddTestSweepers("OcvpSddc", &resource.Sweeper{
			Name:         "OcvpSddc",
			Dependencies: DependencyGraph["sddc"],
			F:            sweepOcvpSddcResource,
		})
	}
}

func sweepOcvpSddcResource(compartment string) error {
	sddcClient := GetTestClients(&schema.ResourceData{}).sddcClient()
	sddcIds, err := getSddcIds(compartment)
	if err != nil {
		return err
	}
	for _, sddcId := range sddcIds {
		if ok := SweeperDefaultResourceId[sddcId]; !ok {
			deleteSddcRequest := oci_ocvp.DeleteSddcRequest{}

			deleteSddcRequest.SddcId = &sddcId

			deleteSddcRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "ocvp")
			_, error := sddcClient.DeleteSddc(context.Background(), deleteSddcRequest)
			if error != nil {
				fmt.Printf("Error deleting Sddc %s %s, It is possible that the resource is already deleted. Please verify manually \n", sddcId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &sddcId, sddcSweepWaitCondition, time.Duration(3*time.Minute),
				sddcSweepResponseFetchOperation, "ocvp", true)
		}
	}
	return nil
}

func getSddcIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "SddcId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	sddcClient := GetTestClients(&schema.ResourceData{}).sddcClient()

	listSddcsRequest := oci_ocvp.ListSddcsRequest{}
	listSddcsRequest.CompartmentId = &compartmentId
	listSddcsRequest.LifecycleState = oci_ocvp.ListSddcsLifecycleStateActive
	listSddcsResponse, err := sddcClient.ListSddcs(context.Background(), listSddcsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Sddc list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, sddc := range listSddcsResponse.Items {
		id := *sddc.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "SddcId", id)
	}
	return resourceIds, nil
}

func sddcSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sddcResponse, ok := response.Response.(oci_ocvp.GetSddcResponse); ok {
		return sddcResponse.LifecycleState != oci_ocvp.LifecycleStatesDeleted
	}
	return false
}

func sddcSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.sddcClient().GetSddc(context.Background(), oci_ocvp.GetSddcRequest{
		SddcId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
