// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v55/core"
)

func init() {
	RegisterDatasource("oci_core_public_ip_pool", CorePublicIpPoolDataSource())
}

func CorePublicIpPoolDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["public_ip_pool_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(CorePublicIpPoolResource(), fieldMap, readSingularCorePublicIpPool)
}

func readSingularCorePublicIpPool(d *schema.ResourceData, m interface{}) error {
	sync := &CorePublicIpPoolDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CorePublicIpPoolDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetPublicIpPoolResponse
}

func (s *CorePublicIpPoolDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CorePublicIpPoolDataSourceCrud) Get() error {
	request := oci_core.GetPublicIpPoolRequest{}

	if publicIpPoolId, ok := s.D.GetOkExists("public_ip_pool_id"); ok {
		tmp := publicIpPoolId.(string)
		request.PublicIpPoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "core")

	response, err := s.Client.GetPublicIpPool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CorePublicIpPoolDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("cidr_blocks", s.Res.CidrBlocks)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
