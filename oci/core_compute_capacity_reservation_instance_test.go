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
	computeCapacityReservationInstanceDataSourceRepresentation = map[string]interface{}{
		"capacity_reservation_id": Representation{RepType: Required, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"availability_domain":     Representation{RepType: Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          Representation{RepType: Optional, Create: `${var.compartment_id}`},
	}

	capacityReservationInstanceRepresentation = map[string]interface{}{
		"availability_domain":     Representation{RepType: Required, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"compartment_id":          Representation{RepType: Required, Create: `${var.compartment_id}`},
		"shape":                   Representation{RepType: Required, Create: `VM.Standard2.1`},
		"capacity_reservation_id": Representation{RepType: Required, Create: `${oci_core_compute_capacity_reservation.test_compute_capacity_reservation.id}`},
		"subnet_id":               Representation{RepType: Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"image":                   Representation{RepType: Required, Create: `${var.InstanceImageOCID[var.region]}`},
	}
)

// issue-routing-tag: core/computeSharedOwnershipVmAndBm
func TestCoreComputeCapacityReservationInstanceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreComputeCapacityReservationInstanceResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_core_compute_capacity_reservation_instances.test_compute_capacity_reservation_instances"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				OciImageIdsVariable +
				ComputeCapacityReservationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Required, Create, computeCapacityReservationRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, capacityReservationInstanceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation_instances", "test_compute_capacity_reservation_instances", Required, Create, computeCapacityReservationInstanceDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.fault_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.shape"),
				resource.TestCheckResourceAttr(datasourceName, "capacity_reservation_instances.0.shape_config.#", "1"),
			),
		},
		{
			Config: config +
				OciImageIdsVariable +
				ComputeCapacityReservationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_compute_capacity_reservation", "test_compute_capacity_reservation", Required, Create, computeCapacityReservationRepresentation) +
				GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, capacityReservationInstanceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_core_compute_capacity_reservation_instances", "test_compute_capacity_reservation_instances", Optional, Create, computeCapacityReservationInstanceDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.availability_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.fault_domain"),
				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "capacity_reservation_instances.0.shape"),
				resource.TestCheckResourceAttr(datasourceName, "capacity_reservation_instances.0.shape_config.#", "1"),
			),
		},
	})
}
