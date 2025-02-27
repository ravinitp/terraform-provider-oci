// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_opsi "github.com/oracle/oci-go-sdk/v55/opsi"
)

func init() {
	RegisterDatasource("oci_opsi_awr_hub", OpsiAwrHubDataSource())
}

func OpsiAwrHubDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["awr_hub_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(OpsiAwrHubResource(), fieldMap, readSingularOpsiAwrHub)
}

func readSingularOpsiAwrHub(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiAwrHubDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).operationsInsightsClient()

	return ReadResource(sync)
}

type OpsiAwrHubDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opsi.OperationsInsightsClient
	Res    *oci_opsi.GetAwrHubResponse
}

func (s *OpsiAwrHubDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpsiAwrHubDataSourceCrud) Get() error {
	request := oci_opsi.GetAwrHubRequest{}

	if awrHubId, ok := s.D.GetOkExists("awr_hub_id"); ok {
		tmp := awrHubId.(string)
		request.AwrHubId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "opsi")

	response, err := s.Client.GetAwrHub(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpsiAwrHubDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AwrMailboxUrl != nil {
		s.D.Set("awr_mailbox_url", *s.Res.AwrMailboxUrl)
	}

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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ObjectStorageBucketName != nil {
		s.D.Set("object_storage_bucket_name", *s.Res.ObjectStorageBucketName)
	}

	if s.Res.OperationsInsightsWarehouseId != nil {
		s.D.Set("operations_insights_warehouse_id", *s.Res.OperationsInsightsWarehouseId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", systemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
