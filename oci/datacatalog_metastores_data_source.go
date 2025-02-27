// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_datacatalog "github.com/oracle/oci-go-sdk/v55/datacatalog"
)

func init() {
	RegisterDatasource("oci_datacatalog_metastores", DatacatalogMetastoresDataSource())
}

func DatacatalogMetastoresDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatacatalogMetastores,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metastores": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(DatacatalogMetastoreResource()),
			},
		},
	}
}

func readDatacatalogMetastores(d *schema.ResourceData, m interface{}) error {
	sync := &DatacatalogMetastoresDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataCatalogClient()

	return ReadResource(sync)
}

type DatacatalogMetastoresDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacatalog.DataCatalogClient
	Res    *oci_datacatalog.ListMetastoresResponse
}

func (s *DatacatalogMetastoresDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatacatalogMetastoresDataSourceCrud) Get() error {
	request := oci_datacatalog.ListMetastoresRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datacatalog.ListMetastoresLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "datacatalog")

	response, err := s.Client.ListMetastores(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMetastores(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatacatalogMetastoresDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatacatalogMetastoresDataSource-", DatacatalogMetastoresDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		metastore := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			metastore["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			metastore["display_name"] = *r.DisplayName
		}

		metastore["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			metastore["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			metastore["lifecycle_details"] = *r.LifecycleDetails
		}

		metastore["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			metastore["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			metastore["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, metastore)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatacatalogMetastoresDataSource().Schema["metastores"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("metastores", resources); err != nil {
		return err
	}

	return nil
}
