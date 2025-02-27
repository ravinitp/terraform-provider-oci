// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apm "github.com/oracle/oci-go-sdk/v55/apmcontrolplane"
)

func init() {
	RegisterDatasource("oci_apm_apm_domains", ApmApmDomainsDataSource())
}

func ApmApmDomainsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readApmApmDomains,
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
			"apm_domains": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(ApmApmDomainResource()),
			},
		},
	}
}

func readApmApmDomains(d *schema.ResourceData, m interface{}) error {
	sync := &ApmApmDomainsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).apmDomainClient()

	return ReadResource(sync)
}

type ApmApmDomainsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm.ApmDomainClient
	Res    *oci_apm.ListApmDomainsResponse
}

func (s *ApmApmDomainsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmApmDomainsDataSourceCrud) Get() error {
	request := oci_apm.ListApmDomainsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_apm.ListApmDomainsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "apm")

	response, err := s.Client.ListApmDomains(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListApmDomains(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ApmApmDomainsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("ApmApmDomainsDataSource-", ApmApmDomainsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		apmDomain := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			apmDomain["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			apmDomain["description"] = *r.Description
		}

		if r.DisplayName != nil {
			apmDomain["display_name"] = *r.DisplayName
		}

		apmDomain["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			apmDomain["id"] = *r.Id
		}

		if r.IsFreeTier != nil {
			apmDomain["is_free_tier"] = *r.IsFreeTier
		}

		apmDomain["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			apmDomain["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			apmDomain["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, apmDomain)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ApmApmDomainsDataSource().Schema["apm_domains"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("apm_domains", resources); err != nil {
		return err
	}

	return nil
}
