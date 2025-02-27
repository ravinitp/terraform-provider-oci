// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v55/common"
	oci_dns "github.com/oracle/oci-go-sdk/v55/dns"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	recordDataSourceRepresentation = map[string]interface{}{
		"zone_name_or_id":   Representation{RepType: Required, Create: `${oci_dns_zone.test_global_zone.name}`},
		"if_modified_since": Representation{RepType: Optional, Create: `ifModifiedSince`},
		"if_none_match":     Representation{RepType: Optional, Create: `ifNoneMatch`},
		"compartment_id":    Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"domain":            Representation{RepType: Optional, Create: `domain`},
		"domain_contains":   Representation{RepType: Optional, Create: `domainContains`},
		"rtype":             Representation{RepType: Optional, Create: `rtype`},
		"zone_version":      Representation{RepType: Optional, Create: `zoneVersion`},
		"filter":            RepresentationGroup{Required, recordDataSourceFilterRepresentation}}
	recordDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_dns_record.test_record.id}`}},
	}

	recordRepresentation = map[string]interface{}{
		"domain":          Representation{RepType: Required, Create: `${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test`},
		"rdata":           Representation{RepType: Required, Create: `192.168.0.1`, Update: `77.77.77.77`},
		"rtype":           Representation{RepType: Required, Create: `A`},
		"ttl":             Representation{RepType: Required, Create: `3600`, Update: `1000`},
		"zone_name_or_id": Representation{RepType: Required, Create: `${oci_dns_zone.test_global_zone.name}`},
		"compartment_id":  Representation{RepType: Optional, Create: `${var.compartment_id}`},
	}

	RecordResourceDependencies = `
data "oci_identity_tenancy" "test_tenancy" {
	tenancy_id = "${var.tenancy_ocid}"
}

resource "oci_dns_zone" "test_global_zone" {
	#Required
	compartment_id = "${var.compartment_id}"
	name = "${data.oci_identity_tenancy.test_tenancy.name}.{{.token}}.oci-record-test"
	zone_type = "PRIMARY"
}
`
)

// issue-routing-tag: dns/default
func TestDnsRecordsResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDnsRecordsResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_dns_record.test_record"

	_, tokenFn := TokenizeWithHttpReplay("dns_resource")
	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_dns_record", "test_record", Required, Create, recordRepresentation), nil), "dns", "record", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+
				GenerateResourceFromRepresentationMap("oci_dns_record", "test_record", Required, Create, recordRepresentation), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "zone_name_or_id"),

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

		// delete before next Create
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies, nil),
		},
		// verify Create with optionals
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+
				GenerateResourceFromRepresentationMap("oci_dns_record", "test_record", Optional, Create, recordRepresentation), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestMatchResourceAttr(resourceName, "domain", regexp.MustCompile("\\.oci-record-test")),
				resource.TestCheckResourceAttr(resourceName, "is_protected", "false"),
				resource.TestCheckResourceAttr(resourceName, "rdata", "192.168.0.1"),
				resource.TestCheckResourceAttrSet(resourceName, "record_hash"),
				resource.TestCheckResourceAttrSet(resourceName, "rrset_version"),
				resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
				resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
				TestCheckResourceAttributesEqual(resourceName, "zone_name_or_id", "oci_dns_zone.test_global_zone", "name"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("resource was not recreated after delete")
					}
					resId = resId2
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+
				GenerateResourceFromRepresentationMap("oci_dns_record", "test_record", Optional, Update, recordRepresentation), nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestMatchResourceAttr(resourceName, "domain", regexp.MustCompile("\\.oci-record-test")),
				resource.TestCheckResourceAttr(resourceName, "is_protected", "false"),
				resource.TestCheckResourceAttr(resourceName, "rdata", "77.77.77.77"),
				resource.TestCheckResourceAttrSet(resourceName, "record_hash"),
				resource.TestCheckResourceAttrSet(resourceName, "rrset_version"),
				resource.TestCheckResourceAttr(resourceName, "rtype", "A"),
				resource.TestCheckResourceAttr(resourceName, "ttl", "1000"),
				TestCheckResourceAttributesEqual(resourceName, "zone_name_or_id", "oci_dns_zone.test_global_zone", "name"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId == resId2 {
						return fmt.Errorf("record hash was the same after an Update, it should be different")
					}
					return err
				},
			),
		},
	})
}

// The datasource tests are kept separate from the previous test steps.
// This was because the datasource steps do not Create a record resource (and won't need one because, because a zone has default records).
// If this was kept in the previous test case, the CheckDestroy step would run after the datasource steps ran and would fail
// because it wouldn't have a record resource to delete and to verify destruction for.
// issue-routing-tag: dns/default
func TestDnsRecordsResource_datasources(t *testing.T) {
	httpreplay.SetScenario("TestDnsRecordsResource_datasources")
	defer httpreplay.SaveScenario()
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	_, tokenFn := TokenizeWithHttpReplay("dns_data_source")
	datasourceName := "data.oci_dns_records.test_records"

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
data "oci_dns_records" "test_records" {
  zone_name_or_id = "${oci_dns_zone.test_global_zone.name}"

  # optional
  domain = "${oci_dns_zone.test_global_zone.name}"
  rtype = "NS"
  sort_by = "ttl"
  sort_order = "DESC"
}`, nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "records.0.rtype", "NS"),
				resource.TestCheckResourceAttr(datasourceName, "records.0.ttl", "86400"),
			),
		},
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
data "oci_dns_records" "test_records" {
  zone_name_or_id = "${oci_dns_zone.test_global_zone.name}"
  domain = "${oci_dns_zone.test_global_zone.name}"
	filter {
	  name = "rtype"
	  values = ["SOA"]
	}
}`, nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "records.#", "1"),
			),
		},
	})
}

// issue-routing-tag: dns/default
func TestDnsRecordsResource_diffSuppression(t *testing.T) {
	httpreplay.SetScenario("TestDnsRecordsResource_diffSuppression")
	defer httpreplay.SaveScenario()
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	_, tokenFn := TokenizeWithHttpReplay("dns_diff")
	resourceName := "oci_dns_record.test_record"
	var resId, resId2 string

	ResourceTest(t, nil, []resource.TestStep{
		// verify AAAA ipv6 shortening does not cause diffs
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${oci_dns_zone.test_global_zone.name}"
	domain = "${oci_dns_zone.test_global_zone.name}"
	rtype = "AAAA"
	rdata = "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	ttl = "3600"
}`, nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "rtype", "AAAA"),
				resource.TestCheckResourceAttr(resourceName, "rdata", "2001:db8:85a3::8a2e:370:7334"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${oci_dns_zone.test_global_zone.name}"
	domain = "${oci_dns_zone.test_global_zone.name}"
	rtype = "AAAA"
	rdata = "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	ttl = "3600"
}`, nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "rtype", "AAAA"),
				resource.TestCheckResourceAttr(resourceName, "rdata", "2001:db8:85a3::8a2e:370:7334"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${oci_dns_zone.test_global_zone.name}"
	domain = "${oci_dns_zone.test_global_zone.name}"
	rtype = "AAAA"
	rdata = "0000:0000:0000:0000:0000:8a2e:0370:0001"
	ttl = "3600"
}`, nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "rdata", "::8a2e:370:1"),
			),
		},
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${oci_dns_zone.test_global_zone.name}"
	domain = "${oci_dns_zone.test_global_zone.name}"
	rtype = "AAAA"
	rdata = "8a2e:0000:0000:0000:0000:0370:0000:0000"
	ttl = "3600"
}`, nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "rdata", "8a2e::370:0:0"),
			),
		},
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${oci_dns_zone.test_global_zone.name}"
	domain = "${oci_dns_zone.test_global_zone.name}"
	rtype = "TXT"
	rdata = "arbitrary text"
	ttl = "3600"
}`, nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "rdata", "\"arbitrary\" \"text\""),
			),
		},
		// this tests represents several record types where the service appends a `.` to the rdata
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${oci_dns_zone.test_global_zone.name}"
	domain = "${oci_dns_zone.test_global_zone.name}"
	rtype = "ALIAS"
	rdata = "other.tf-provider.oci-record-test"
	ttl = "3600"
}`, nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "rdata", "other.tf-provider.oci-record-test."),
			),
		},
	})
}

// issue-routing-tag: dns/default
func TestDnsRecordsResource_badUpdate(t *testing.T) {
	httpreplay.SetScenario("TestDnsRecordsResource_badUpdate")
	defer httpreplay.SaveScenario()
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	_, tokenFn := TokenizeWithHttpReplay("dns_bad_update")
	resourceName := "oci_dns_record.test_record"

	ResourceTest(t, nil, []resource.TestStep{
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${oci_dns_zone.test_global_zone.name}"
	domain = "${oci_dns_zone.test_global_zone.name}"
	rtype = "A"
	rdata = "192.168.0.1"
	ttl = "3600"
}`, nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "rdata", "192.168.0.1"),
			),
		},
		{
			Config: tokenFn(config+compartmentIdVariableStr+RecordResourceDependencies+`
resource "oci_dns_record" "test_record" {
	zone_name_or_id = "${oci_dns_zone.test_global_zone.name}"
	domain = "${oci_dns_zone.test_global_zone.name}"
	rtype = "A"
	rdata = "192.168.0.1"
	ttl = "-1"
}`, nil),
			Check: ComposeAggregateTestCheckFuncWrapper(
			//resource.TestCheckResourceAttr(resourceName, "rdata", "192.168.0.1"),
			// todo: this test was attempting to verify the resource is not changed if the Update operation fails
			// but this terraform testing library does not run "Checks" if you add an error expectation ;_;
			),
			ExpectError: regexp.MustCompile("-1 is not a valid value for TTL"),
		},
	})
}

func testAccCheckDnsRecordDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dnsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dns_record" {
			noResourceFound = false
			request := oci_dns.GetZoneRecordsRequest{}

			if value, ok := rs.Primary.Attributes["zone_name_or_id"]; ok {
				request.ZoneNameOrId = &value
			}

			if value, ok := rs.Primary.Attributes["compartment_id"]; ok {
				request.CompartmentId = &value
			}

			response, err := client.GetZoneRecords(context.Background(), request)
			if err == nil {
				// Convert the InstanceState attributes to a ResourceData expected by the lookup function
				attributes := convertToObjectMap(rs.Primary.Attributes)
				resourceData := schema.TestResourceDataRaw(&testing.T{}, DnsRecordResource().Schema, attributes)

				//page through records
				recordCollection := response.RecordCollection
				request.Page = response.OpcNextPage

				for request.Page != nil {
					listResponse, err := client.GetZoneRecords(context.Background(), request)
					if err != nil {
						return err
					}

					recordCollection.Items = append(recordCollection.Items, listResponse.Items...)
					request.Page = listResponse.OpcNextPage
				}
				_, err = findItem(&response.Items, resourceData)
				if err == nil {
					return fmt.Errorf("resource still exists")
				}

				// no error and item not found, item is deleted
				return nil
			}

			// TODO: If we get here, then technically this isn't verifying that the record resource was destroyed.
			// But it is verifying that at least the zone was destroyed (which guarantees that the records were destroyed)
			// This is a test gap because of Terraform test framework destroying all resources.
			// Ideally, the test framework should do a targeted destroy of the record prior to calling CheckDestroy.
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
