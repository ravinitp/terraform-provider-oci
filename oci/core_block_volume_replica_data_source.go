// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v55/core"
)

func init() {
	RegisterDatasource("oci_core_block_volume_replica", CoreBlockVolumeReplicaDataSource())
}

func CoreBlockVolumeReplicaDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreBlockVolumeReplica,
		Schema: map[string]*schema.Schema{
			"block_volume_replica_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"availability_domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"block_volume_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"size_in_gbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_synced": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularCoreBlockVolumeReplica(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBlockVolumeReplicaDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient()

	return ReadResource(sync)
}

type CoreBlockVolumeReplicaDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.GetBlockVolumeReplicaResponse
}

func (s *CoreBlockVolumeReplicaDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreBlockVolumeReplicaDataSourceCrud) Get() error {
	request := oci_core.GetBlockVolumeReplicaRequest{}

	if blockVolumeReplicaId, ok := s.D.GetOkExists("block_volume_replica_id"); ok {
		tmp := blockVolumeReplicaId.(string)
		request.BlockVolumeReplicaId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "core")

	response, err := s.Client.GetBlockVolumeReplica(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreBlockVolumeReplicaDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.BlockVolumeId != nil {
		s.D.Set("block_volume_id", *s.Res.BlockVolumeId)
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

	if s.Res.SizeInGBs != nil {
		s.D.Set("size_in_gbs", strconv.FormatInt(*s.Res.SizeInGBs, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastSynced != nil {
		s.D.Set("time_last_synced", s.Res.TimeLastSynced.String())
	}

	return nil
}
