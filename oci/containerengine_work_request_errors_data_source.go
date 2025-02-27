// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/v55/containerengine"
)

func init() {
	RegisterDatasource("oci_containerengine_work_request_errors", ContainerengineWorkRequestErrorsDataSource())
}

func ContainerengineWorkRequestErrorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readContainerengineWorkRequestErrors,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"work_request_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"work_request_errors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"timestamp": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readContainerengineWorkRequestErrors(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineWorkRequestErrorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient()

	return ReadResource(sync)
}

type ContainerengineWorkRequestErrorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.ListWorkRequestErrorsResponse
}

func (s *ContainerengineWorkRequestErrorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineWorkRequestErrorsDataSourceCrud) Get() error {
	request := oci_containerengine.ListWorkRequestErrorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if workRequestId, ok := s.D.GetOkExists("work_request_id"); ok {
		tmp := workRequestId.(string)
		request.WorkRequestId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "containerengine")

	response, err := s.Client.ListWorkRequestErrors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineWorkRequestErrorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("ContainerengineWorkRequestErrorsDataSource-", ContainerengineWorkRequestErrorsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		workRequestError := map[string]interface{}{}

		if r.Code != nil {
			workRequestError["code"] = *r.Code
		}

		if r.Message != nil {
			workRequestError["message"] = *r.Message
		}

		if r.Timestamp != nil {
			workRequestError["timestamp"] = r.Timestamp.String()
		}

		resources = append(resources, workRequestError)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ContainerengineWorkRequestErrorsDataSource().Schema["work_request_errors"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("work_request_errors", resources); err != nil {
		return err
	}

	return nil
}
